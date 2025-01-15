package main

import (
	"time"

	"github.com/ravinggo/game/common/closer"

	_ "github.com/ravinggo/examples/game-server/env"
	"github.com/ravinggo/examples/game-server/service"
)

func main() {
	s := service.NewGameServer([]string{"nats://192.168.0.166:4222"}, false, 0, 0, time.Second*10)
	s.Start()
	closer.WaitClose(
		func() {
			s.Stop()
		},
	)
}
