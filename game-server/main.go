package main

import (
	"time"

	"github.com/ravinggo/game/common/closer"

	_ "github.com/ravinggo/examples/game-server/env"
	"github.com/ravinggo/examples/login-server/service"
)

func main() {
	s := service.NewLoginServer([]string{"127.0.0.1:4222"}, false, 0, 0, time.Second*10)
	s.Start()
	closer.WaitClose(
		func() {
			s.Stop()
		},
	)
}
