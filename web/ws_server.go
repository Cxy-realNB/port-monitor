package web

import (
	"golang.org/x/net/websocket"
	"main/clog"
	"main/constant"
	_ "net/http/pprof"
)

func entry(conn *websocket.Conn) {
	defer conn.Close()
	clientId := conn.Request().URL.Query()["clientId"][0]
	devName := conn.Request().URL.Query()["devName"][0]
	if err := websocket.Message.Send(conn, "ping"); err == nil {
		ws := constant.Ws{
			ClientId: clientId,
			DevName:  devName,
			Conn:     conn,
		}
		constant.Socks[clientId] = ws
		clog.InfoLogger.Println(clientId + " connected.")
	}
	for {
		var reply string
		if err := websocket.Message.Receive(conn, &reply); err != nil {
		}
		if reply != "" {
			clog.InfoLogger.Println(clientId + ": " + reply)
		}
	}
}
