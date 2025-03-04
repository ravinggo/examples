package main

import (
	"github.com/ravinggo/game/common/closer"

	_ "github.com/ravinggo/examples/login-server/env"
	"github.com/ravinggo/examples/login-server/service"
)

func main() {
	s := service.NewLoginServer([]string{"nats://192.168.0.166:4222"})
	s.Start()
	closer.WaitClose(
		func() {
			s.Stop()
		},
	)
}
