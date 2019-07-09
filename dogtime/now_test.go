package dogtime_test

import (
	"testing"
	"time"

	"github.com/sinmetal/silverdog/dogtime"
)

func TestNow(t *testing.T) {
	now := dogtime.Now()
	if now.IsZero() {
		t.Errorf("Now is zero")
	}
}

func TestStockNower_Now(t *testing.T) {
	now := time.Now()

	mock := dogtime.StockNower{}
	dogtime.SetNower(&mock)
	mock.AddStockTime(now)

	got := dogtime.Now()
	if now != got {
		t.Errorf("%+v not equal %+v", now, got)
	}
}
