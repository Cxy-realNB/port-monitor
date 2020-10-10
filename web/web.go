package web

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"golang.org/x/net/websocket"
	"html/template"
	"main/argsparse"
	"main/clog"
	"main/constant"
	"main/devices"
	"net/http"
)

var logRequest bool

func control(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("index.html").Funcs(template.FuncMap{
		"JSONS":     json.Marshal,
		"ByteToStr": ByteToString,
	})
	t, err := tmpl.ParseFiles("./static/index.html")
	if err != nil {
		clog.FatalLogger.Println("Can not create template: " + tmpl.Name())
		return
	}
	data := make(map[string]interface{})
	data["devices"] = devices.GetDevices()
	data["ip"] = argsparse.Args.IP
	data["port"] = argsparse.Args.Port
	t.ExecuteTemplate(w, "index.html", data)
}

func devListen(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("listener.html").Funcs(template.FuncMap{
		"JSONS":     json.Marshal,
		"ByteToStr": ByteToString,
	})
	t, err := tmpl.ParseFiles("./static/listener.html")
	if err != nil {
		clog.FatalLogger.Println("Can not create template: " + tmpl.Name())
		return
	}
	data := make(map[string]interface{})
	vars := r.URL.Query()
	data["device"], _ = devices.GetDevicesByName(vars["name"][0])
	data["ip"] = argsparse.Args.IP
	data["port"] = argsparse.Args.Port
	t.ExecuteTemplate(w, "listener.html", data)
}

func getClientId(w http.ResponseWriter, r *http.Request) {
	u1 := uuid.NewV4()
	w.Write([]byte(u1.String()))
}

func closeWs(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	clientId := vars["clientId"][0]
	ws, ok := constant.Socks[clientId]
	if ok {
		ws.Conn.Close()
		clog.InfoLogger.Println(clientId + " closed.")
		delete(constant.Socks, clientId)
	}
}

func Interceptor(f func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	// decorate real handle, do something like logging
	return func(w http.ResponseWriter, r *http.Request) {
		if logRequest {
			clog.InfoLogger.Println(*r)
		}
		f(w, r)
	}
}

func InitServer(ip string, port int, lr bool) {
	logRequest = lr
	// all request path below
	http.HandleFunc("/", Interceptor(control))
	http.HandleFunc("/index", Interceptor(control))
	http.HandleFunc("/listen", Interceptor(devListen))
	http.HandleFunc("/getClientId", Interceptor(getClientId))
	http.HandleFunc("/closeWs", Interceptor(closeWs))

	// websocket connection
	http.Handle("/ws", websocket.Handler(entry))

	url := ip + ":" + fmt.Sprintf("%d", port)
	err := http.ListenAndServe(url, nil)
	if err != nil {
		clog.FatalLogger.Println("Can not initial server: " + url)
	}
}
