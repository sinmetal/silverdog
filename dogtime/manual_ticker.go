package dogtime

import (
	"errors"
	"fmt"
	"time"
)

type ManualTicker struct {
	ch chan time.Time
}

func NewManualTicker(d time.Duration) MockTicker {
	return &ManualTicker{
		make(chan time.Time, 1),
	}
}

func (t *ManualTicker) Chan() <-chan time.Time {
	return t.ch
}

func (t *ManualTicker) Stop() {
	// noop
}

func (t *ManualTicker) Fire() {
	t.ch <- time.Now()
}

type ManualTickerCreator struct {
	ms []MockTicker
}

var _ MockTickerCreator = &ManualTickerCreator{}

func NewManualTickerCreator() *ManualTickerCreator {
	return &ManualTickerCreator{}
}

func (c *ManualTickerCreator) NewTicker(d time.Duration) Ticker {
	mt := NewManualTicker(d)
	c.ms = append(c.ms, mt)
	return mt
}

func (c *ManualTickerCreator) GetMockTicker(i int) (MockTicker, error) {
	if len(c.ms) <= i {
		return nil, errors.New(fmt.Sprintf("out of range. []MockTicker.length is %d target index is %d", len(c.ms), i))
	}
	return c.ms[i], nil
}
