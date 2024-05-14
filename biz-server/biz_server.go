package main

import (
	"github.com/aldalee/high-performance-online-game-server/comm/log"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
	"path"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func webSocketHandshake(w http.ResponseWriter, r *http.Request) {
	if w == nil || r == nil {
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("webSocket upgrade error, %v+", err)
		return
	}
	defer func() { _ = conn.Close() }()
	log.Info("new clients connected")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Error("%v+", err)
			break
		}
		log.Info("%v", msg)
	}
}

func main() {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	log.Config(path.Dir(exe) + "/logs/biz-server")
	log.Info("start business server")

	http.HandleFunc("/websocket", webSocketHandshake)
	_ = http.ListenAndServe("127.0.0.1:80", nil)
}
