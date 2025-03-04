package main

import (
	"github.com/ravinggo/game/common/closer"

	_ "github.com/ravinggo/examples/game-server/env"
	"github.com/ravinggo/examples/game-server/service"
)

func main() {
	s := service.NewGameServer([]string{"nats://192.168.0.166:4222"})
	s.Start()
	closer.WaitClose(
		func() {
			s.Stop()
		},
	)
}
