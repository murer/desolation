package guest

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/murer/desolation/message"
	"github.com/murer/desolation/util"
)

func Start() {

	// http.HandleFunc("/", Handle)
	log.Printf("Starting server")
	err := http.ListenAndServe("0.0.0.0:5010", Handler())
	util.Check(err)
}

var static = ""

func Handler() http.Handler {
	static = "guest/public"
	if !util.FileExists(static) {
		static = "public"
		if !util.FileExists(static) {
			log.Panicf("static dir not found: %s", static)
		}
	}
	mux := http.NewServeMux()
	mux.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir(static))))
	mux.Handle("/", http.HandlerFunc(Handle))
	return mux
}

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Access: %s %s %s", r.RemoteAddr, r.Method, r.URL)
	if r.Method == "GET" && r.URL.Path == "/api/version.txt" {
		util.RespText(w, util.Version)
	} else if r.Method == "POST" && r.URL.Path == "/api/command" {
		HandleCommand(w, r)
	} else if r.Method == "GET" && r.URL.Path == "/" {
		HandleIndex(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func messageExtract(r *http.Request) *message.Message {
	reqBody := util.ReadAllString(r.Body)
	return message.Decode(reqBody)
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadFile(static + "/index.html")
	util.Check(err)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write(body)
}

func HandleCommand(w http.ResponseWriter, r *http.Request) {
	msg := messageExtract(r)
	var ret *message.Message
	if msg.Name == "echo" {
		ret = msg
	} else if msg.Name == "write" {
		ret = HandleCommandWrite(msg, w, r)
	} else if msg.Name == "read" {
		ret = HandleCommandRead(msg, w, r)
	} else if msg.Name == "cw" {
		ret = HandleCommandCW(msg, w, r)
	} else if msg.Name == "init" {
		ret = HandleCommandInit(msg, w, r)
	} else {
		ret = &message.Message{Name: "unknown", Headers: map[string]string{}, Payload: ""}
	}
	rid, exists := msg.Headers["rid"]
	if exists {
		ret.Headers["rid"] = rid
	}
	respBody := message.Encode(ret)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(respBody))
}
