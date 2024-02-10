package main

import (
	"flag"
	"time"

	"github.com/BullionBear/binance-mongo/stream"
	"github.com/adshao/go-binance/v2"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()

	glog.Infoln("Starting without loading config.")
	defer glog.Flush()

	symbol := "SOLUSDT"

	// Start initial WebSocket connection
	doneC, stopC, err := stream.ConnectWsDepthServe(symbol, stream.UpdateLastUpdateTime, func(event *binance.WsDepthEvent) {
		glog.Infoln(event)
		stream.UpdateLastUpdateTime()
	}, func(err error) {
		glog.Errorln(err)
	})
	if err != nil {
		glog.Errorln(err)
		return
	}

	// Start the reconnection monitor in a goroutine
	go stream.Reconnect(symbol, stream.UpdateLastUpdateTime)

	// Use stopC to control exit
	go func() {
		time.Sleep(10 * time.Minute) // Example: stop after 10 minutes for demonstration
		stopC <- struct{}{}
	}()

	<-doneC // Wait here until done
}
