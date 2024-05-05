package main

import (
	"github.com/aldalee/high-performance-online-game-server/comm/log"
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
}
