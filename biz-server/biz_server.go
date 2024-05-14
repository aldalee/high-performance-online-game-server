package main

import (
	"github.com/aldalee/high-performance-online-game-server/comm/log"
	"net/http"
	"os"
	"path"
)

func main() {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	log.Config(path.Dir(exe) + "/logs/biz-server")
	log.Info("start business server")

	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello World!"))
	})
	_ = http.ListenAndServe("127.0.0.1:80", nil)
}
