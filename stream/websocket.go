// stream/stream.go

package stream

import (
	"sync"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/golang/glog"
)

var mu sync.Mutex
var lastUpdateTime time.Time

func UpdateLastUpdateTime() {
	mu.Lock()
	lastUpdateTime = time.Now()
	mu.Unlock()
}

func ConnectWsDepthServe(symbol string, updateLastUpdateTimeFunc func(), wsDepthHandler func(event *binance.WsDepthEvent), errHandler func(error)) (doneC, stopC chan struct{}, err error) {
	doneC, stopC, err = binance.WsDepthServe(symbol, wsDepthHandler, errHandler)
	if err != nil {
		return nil, nil, err
	}
	updateLastUpdateTimeFunc() // Initialize last update time
	return
}

func Reconnect(symbol string, updateLastUpdateTimeFunc func()) {
	for {
		time.Sleep(5 * time.Second) // Check every 5 seconds
		mu.Lock()
		if time.Since(lastUpdateTime) > 5*time.Second {
			glog.Infoln("More than 5 seconds since last update, reconnecting...")
			mu.Unlock() // Unlock before reconnecting because it may call updateLastUpdateTimeFunc

			doneC, stopC, err := ConnectWsDepthServe(symbol, updateLastUpdateTimeFunc, func(event *binance.WsDepthEvent) {
				glog.Infoln(event)
				updateLastUpdateTimeFunc()
			}, func(err error) {
				glog.Errorln(err)
			})
			if err != nil {
				glog.Errorln(err)
				return
			}
			// Use stopC to exit or reconnect
			go func() {
				time.Sleep(5 * time.Second)
				stopC <- struct{}{}
			}()
			<-doneC                    // Wait until done
			updateLastUpdateTimeFunc() // Reset last update time after reconnection
		} else {
			mu.Unlock()
		}
	}
}
