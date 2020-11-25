package guest_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/murer/desolation/guest"
	"github.com/murer/desolation/message"
	"github.com/murer/desolation/util"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	server := httptest.NewServer(http.Handler(guest.Handler()))
	defer server.Close()
	t.Logf("URL: %s", server.URL)
	resp, err := http.Get(server.URL + "/api/version.txt")
	util.Check(err)
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "text/plain; charset=UTF-8", resp.Header.Get("Content-Type"))
	assert.Equal(t, util.Version, util.ReadAllString(resp.Body))
}

func TestUnknown(t *testing.T) {
	server := httptest.NewServer(http.Handler(guest.Handler()))
	defer server.Close()
	t.Logf("URL: %s", server.URL)

	msg := &message.Message{
		Name:    "hfewi",
		Headers: map[string]string{"foo": "1", "bar": "2"},
		Payload: "bXVyZXI=",
	}
	resp, err := http.Post(server.URL+"/api/command", "application/json", bytes.NewReader([]byte(message.Encode(msg))))
	util.Check(err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, &message.Message{Name: "unknown", Headers: map[string]string{}, Payload: ""}, message.Decode(util.ReadAllString(resp.Body)))
}

func TestEchoJson(t *testing.T) {
	server := httptest.NewServer(http.Handler(guest.Handler()))
	defer server.Close()
	t.Logf("URL: %s", server.URL)

	msg := &message.Message{
		Name:    "echo",
		Headers: map[string]string{"foo": "1", "bar": "2"},
		Payload: "bXVyZXI=",
	}
	resp, err := http.Post(server.URL+"/api/command", "application/json", bytes.NewReader([]byte(message.Encode(msg))))
	util.Check(err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, msg, message.Decode(util.ReadAllString(resp.Body)))
}

func TestStatic(t *testing.T) {
	server := httptest.NewServer(http.Handler(guest.Handler()))
	defer server.Close()
	t.Logf("URL: %s", server.URL)
	resp, err := http.Get(server.URL + "/public/ping.txt")
	util.Check(err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "text/plain; charset=utf-8", resp.Header.Get("Content-Type"))
	assert.Equal(t, "OK", util.ReadAllString(resp.Body))
}

func TestCommandWrite(t *testing.T) {
	server := httptest.NewServer(http.Handler(guest.Handler()))
	defer server.Close()
	t.Logf("URL: %s", server.URL)

	msg := &message.Message{
		Name:    "write",
		Headers: map[string]string{"rid": "a"},
		Payload: "bXVyZXI=",
	}
	resp, err := http.Post(server.URL+"/api/command", "application/json", bytes.NewReader([]byte(message.Encode(msg))))
	util.Check(err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	rmsg := message.Decode(util.ReadAllString(resp.Body))
	assert.Equal(t, "ok", rmsg.Name)
	assert.Equal(t, "a", rmsg.Get("rid"))
	assert.Equal(t, "", rmsg.Payload)
}

func TestCommandRead(t *testing.T) {
	original := guest.In
	guest.In = ioutil.NopCloser(bytes.NewReader([]byte("test")))
	defer func() { guest.In = original }()

	server := httptest.NewServer(http.Handler(guest.Handler()))
	defer server.Close()
	t.Logf("URL: %s", server.URL)

	msg := &message.Message{
		Name:    "read",
		Headers: map[string]string{"rid": "a"},
		Payload: "",
	}
	resp, err := http.Post(server.URL+"/api/command", "application/json", bytes.NewReader([]byte(message.Encode(msg))))
	util.Check(err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	rmsg := message.Decode(util.ReadAllString(resp.Body))
	assert.Equal(t, "ok", rmsg.Name)
	assert.Equal(t, "a", rmsg.Get("rid"))
	assert.Equal(t, "test", string(util.B64Dec(rmsg.Payload)))
}

func TestCommandCW(t *testing.T) {
	original := guest.Out
	pin, pout := io.Pipe()
	guest.Out = pout
	defer func() { guest.Out = original }()

	server := httptest.NewServer(http.Handler(guest.Handler()))
	defer server.Close()
	t.Logf("URL: %s", server.URL)

	msg := &message.Message{
		Name:    "cw",
		Headers: map[string]string{"rid": "a"},
		Payload: "",
	}
	resp, err := http.Post(server.URL+"/api/command", "application/json", bytes.NewReader([]byte(message.Encode(msg))))
	util.Check(err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	rmsg := message.Decode(util.ReadAllString(resp.Body))
	assert.Equal(t, "ok", rmsg.Name)
	assert.Equal(t, "a", rmsg.Get("rid"))
	assert.Equal(t, "", string(util.B64Dec(rmsg.Payload)))

	assert.Equal(t, "", util.ReadAllString(pin))
}

func TestCommandInit(t *testing.T) {
	server := httptest.NewServer(http.Handler(guest.Handler()))
	defer server.Close()
	t.Logf("URL: %s", server.URL)

	msg := &message.Message{
		Name:    "init",
		Headers: map[string]string{"rid": "a"},
		Payload: "",
	}
	resp, err := http.Post(server.URL+"/api/command", "application/json", bytes.NewReader([]byte(message.Encode(msg))))
	util.Check(err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	rmsg := message.Decode(util.ReadAllString(resp.Body))
	assert.Equal(t, "ok", rmsg.Name)
	assert.Equal(t, "a", rmsg.Get("rid"))
	assert.Equal(t, "127.0.0.1", rmsg.Get("host"))
	assert.Equal(t, "22", rmsg.Get("port"))
	assert.Equal(t, "", string(util.B64Dec(rmsg.Payload)))
}
