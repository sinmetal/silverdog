package dogtime

import "time"

type Nower interface {
	Now() time.Time
}

var mock Nower

type StockNower struct {
	currentIndex int
	stockTime    []time.Time
}

func (n *StockNower) Now() time.Time {
	now := n.stockTime[n.currentIndex]
	n.currentIndex++
	return now
}

func (n *StockNower) AddStockTime(now time.Time) {
	n.stockTime = append(n.stockTime, now)
}

func SetNower(nower Nower) {
	mock = nower
}

func Now() time.Time {
	if mock != nil {
		return mock.Now()
	}
	return time.Now()
}
