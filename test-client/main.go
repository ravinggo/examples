package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ravinggo/game/common/logger"
	"github.com/ravinggo/game/common/natsclient"

	"github.com/ravinggo/examples/test-client/service"
	"github.com/ravinggo/examples/test-client/stat"
)

func main() {
	logger.InitDefaultLogger()
	stat_ := stat.NewStat()
	cnc := natsclient.NewClusterClient(
		[]string{
			"nats://192.168.0.166:4222",
		}, time.Second*10,
	)
	totalCount := int64(1e6)
	runCount := totalCount
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		client := service.NewClient(cnc, stat_)
		go func() {
			defer wg.Done()
			for {
				if atomic.AddInt64(&runCount, -1) > 0 {
					err := client.LoginReq()
					if err != nil {
						panic(err)
					}
				} else {
					return
				}
				if atomic.AddInt64(&runCount, -1) > 0 {
					err := client.Echo()
					if err != nil {
						panic(err)
					}
				} else {
					return
				}
			}
		}()
	}
	oldCount := runCount
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(runCount, oldCount-runCount)
			oldCount = runCount
		}
	}()
	wg.Wait()
	fmt.Println("totalCount: ", totalCount, "avg delay: ", stat_.Avg())
}
