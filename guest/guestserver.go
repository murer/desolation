package guest

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/murer/desolation/guest/public"
	"github.com/murer/desolation/message"
	"github.com/murer/desolation/util"
)

func Start() {

	GuestInOutInit()
	// http.HandleFunc("/", Handle)
	log.Printf("Starting server: http://localhost:5010/")
	err := http.ListenAndServe("0.0.0.0:5010", Handler())
	util.Check(err)
}

var static = ""

func Handler() http.Handler {
	static = "guest/public"
	if !util.FileExists(static) {
		static = "public"
		if !util.FileExists(static) && len(public.StaticFiles) == 0 {
			log.Panicf("static dir not found: %s", static)
		}
	}
	mux := http.NewServeMux()
	//mux.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir(static))))
	// mux.Handle("/public", http.HandlerFunc(HandleStatic))
	mux.Handle("/", http.HandlerFunc(Handle))
	return mux
}

func HandleStatic(w http.ResponseWriter, r *http.Request, url string) {
	filename := filepath.Base(url)
	contentType := ""
	if strings.HasSuffix(filename, ".js") {
		contentType = "application/json"
	} else if strings.HasSuffix(filename, ".css") {
		contentType = "text/css; charset=utf-8"
	} else if strings.HasSuffix(filename, ".html") {
		contentType = "text/html; charset=utf-8"
	} else if strings.HasSuffix(filename, ".txt") {
		contentType = "text/plain; charset=utf-8"
	} else {
		log.Printf("Unknown ext: %s", filename)
		http.NotFound(w, r)
		return
	}
	ret, ok := public.StaticFiles[filename]
	if !ok {
		content, err := ioutil.ReadFile(static + "/" + filename)
		util.Check(err)
		ret = string(content)
	}
	w.Header().Set("Content-Type", contentType)
	w.Write([]byte(ret))

}

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Access: %s %s %s", r.RemoteAddr, r.Method, r.URL)
	if r.Method == "GET" && r.URL.Path == "/api/version.txt" {
		util.RespText(w, util.Version)
	} else if r.Method == "POST" && r.URL.Path == "/api/command" {
		HandleCommand(w, r)
	} else if r.Method == "GET" && strings.HasPrefix(r.URL.Path, "/public") {
		HandleStatic(w, r, r.URL.Path)
	} else if r.Method == "GET" && r.URL.Path == "/" {
		HandleStatic(w, r, "/public/index.html")
	} else {
		http.NotFound(w, r)
	}
}

func messageExtract(r *http.Request) *message.Message {
	reqBody := util.ReadAllString(r.Body)
	return message.Decode(reqBody)
}

func HandleCommand(w http.ResponseWriter, r *http.Request) {
	msg := messageExtract(r)
	log.Printf("Received: %s", msg.Basic())
	var ret *message.Message
	if msg.Op == message.OpEcho {
		ret = msg
	} else if msg.Op == message.OpWrite {
		ret = HandleCommandWrite(msg, w, r)
	} else if msg.Op == message.OpRead {
		ret = HandleCommandRead(msg, w, r)
	} else if msg.Op == message.OpCloseWrite {
		ret = HandleCommandCW(msg, w, r)
	} else if msg.Op == message.OpInit {
		ret = HandleCommandInit(msg, w, r)
	} else {
		ret = message.Create(message.OpUnknown, 0, []byte{})
	}
	ret.Rid = msg.Rid
	log.Printf("Sent: %s", ret.Basic())
	respBody := ret.Encode()
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(respBody))
}
