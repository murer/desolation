package guest_test

import (
	"bytes"
	"io"
	"log"
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

	msg := message.CreateString(200, 4, "murer")
	resp, err := http.Post(server.URL+"/api/command", "text/plain", bytes.NewReader([]byte(msg.Encode())))
	util.Check(err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "text/plain; charset=utf-8", resp.Header.Get("Content-Type"))
	assert.Equal(t, message.CreateUnknown(200, 4), message.Decode(util.ReadAllString(resp.Body)))
}

func TestEchoJson(t *testing.T) {
	server := httptest.NewServer(http.Handler(guest.Handler()))
	defer server.Close()
	t.Logf("URL: %s", server.URL)

	msg := message.CreateString(message.OpEcho, 5, "murer")
	resp, err := http.Post(server.URL+"/api/command", "text/plain", bytes.NewReader([]byte(msg.Encode())))
	util.Check(err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "text/plain; charset=utf-8", resp.Header.Get("Content-Type"))
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
	guest.GuestInOutInit()
	server := httptest.NewServer(http.Handler(guest.Handler()))
	defer server.Close()
	t.Logf("URL: %s", server.URL)

	msg := message.CreateString(message.OpWrite, 5, "murer")
	resp, err := http.Post(server.URL+"/api/command", "text/plain", bytes.NewReader([]byte(msg.Encode())))
	util.Check(err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "text/plain; charset=utf-8", resp.Header.Get("Content-Type"))
	rmsg := message.Decode(util.ReadAllString(resp.Body))
	assert.Equal(t, message.OpOk, rmsg.Op)
	assert.Equal(t, uint32(5), rmsg.Rid)
	assert.Equal(t, []byte{}, rmsg.Payload)
}

func TestCommandRead(t *testing.T) {
	original := guest.In
	guest.In = util.ChannelReader(bytes.NewReader([]byte("test")), 256)
	defer func() { guest.In = original }()

	server := httptest.NewServer(http.Handler(guest.Handler()))
	defer server.Close()
	t.Logf("URL: %s", server.URL)

	msg := message.CreateString(message.OpRead, 6, "test")
	resp, err := http.Post(server.URL+"/api/command", "text/plain", bytes.NewReader([]byte(msg.Encode())))
	util.Check(err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "text/plain; charset=utf-8", resp.Header.Get("Content-Type"))
	rmsg := message.Decode(util.ReadAllString(resp.Body))
	assert.Equal(t, message.OpOk, rmsg.Op)
	assert.Equal(t, uint32(6), rmsg.Rid)
	assert.Equal(t, "test", string(rmsg.Payload))
}

func TestCommandCW(t *testing.T) {
	original := guest.Out
	pin, pout := io.Pipe()
	guest.Out = pout
	defer func() { guest.Out = original }()

	server := httptest.NewServer(http.Handler(guest.Handler()))
	defer server.Close()
	t.Logf("URL: %s", server.URL)

	msg := message.Create(message.OpCloseWrite, 7, []byte{})
	resp, err := http.Post(server.URL+"/api/command", "text/plain", bytes.NewReader([]byte(msg.Encode())))
	util.Check(err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "text/plain; charset=utf-8", resp.Header.Get("Content-Type"))
	rmsg := message.Decode(util.ReadAllString(resp.Body))
	assert.Equal(t, message.OpOk, rmsg.Op)
	assert.Equal(t, uint32(7), rmsg.Rid)
	assert.Equal(t, "", string(rmsg.Payload))

	assert.Equal(t, "", util.ReadAllString(pin))
}

func TestCommandInit(t *testing.T) {
	server := httptest.NewServer(http.Handler(guest.Handler()))
	defer server.Close()
	t.Logf("URL: %s", server.URL)

	msg := message.Create(message.OpInit, 1, []byte{})
	code := msg.Encode()
	log.Printf("init code: %s", code)
	resp, err := http.Post(server.URL+"/api/command", "text/plain", bytes.NewReader([]byte(code)))
	util.Check(err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "text/plain; charset=utf-8", resp.Header.Get("Content-Type"))
	rmsg := message.Decode(util.ReadAllString(resp.Body))
	assert.Equal(t, message.OpOk, rmsg.Op)
	assert.Equal(t, uint32(1), rmsg.Rid)
	m := rmsg.PayloadMap()
	assert.Equal(t, "127.0.0.1", m["host"])
	assert.Equal(t, "22", m["port"])
}
