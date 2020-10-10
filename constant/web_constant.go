package constant

import "golang.org/x/net/websocket"

type Ws struct {
	ClientId string
	DevName  string
	Conn     *websocket.Conn
}

var Socks = make(map[string]Ws)
