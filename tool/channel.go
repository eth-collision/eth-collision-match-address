package tool

import "time"

func NewProxyTicker(d time.Duration) (chan time.Time, func()) {
	ticker := time.NewTicker(d)
	proxyChan := make(chan time.Time)
	callback := func() {
		proxyChan <- time.Now()
		for {
			select {
			case <-ticker.C:
				proxyChan <- time.Now()
			}
		}
	}
	return proxyChan, callback
}
