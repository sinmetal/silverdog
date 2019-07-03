package dogtime

import "time"

type Ticker interface {
	Chan() <-chan time.Time
	Stop()
}

type MockTicker interface {
	Chan() <-chan time.Time
	Stop()
	Fire()
}

type MockTickerCreator interface {
	NewTicker(d time.Duration) Ticker
	GetMockTicker(i int) (MockTicker, error)
}

var mtc MockTickerCreator

func NewTicker(d time.Duration) Ticker {
	if mtc != nil {
		return mtc.NewTicker(d)
	}
	t := time.NewTicker(d)
	return &defaultTicker{*t}
}

func SetMockTickerCreator(creator MockTickerCreator) {
	mtc = creator
}

type defaultTicker struct {
	ticker time.Ticker
}

func (t *defaultTicker) Chan() <-chan time.Time {
	return t.ticker.C
}

func (t *defaultTicker) Stop() {
	t.ticker.Stop()
}
