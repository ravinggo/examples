package main

import (
	"github.com/ravinggo/game/common/closer"
	"github.com/ravinggo/game/common/logger"

	_ "github.com/ravinggo/examples/game-server/env"
	"github.com/ravinggo/examples/game-server/service"
)

func main() {
	logger.InitDefaultLogger()
	s := service.NewGameServer([]string{"nats://192.168.0.166:4222"})
	s.Start()
	closer.WaitClose(
		func() {
			s.Stop()
		},
	)
}
