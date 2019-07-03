package dogtime_test

import (
	"testing"
	"time"

	"github.com/sinmetal/silverdog/dogtime"
)

func TestManualTicker_NewTicker(t *testing.T) {
	mtc := dogtime.NewManualTickerCreator()
	dogtime.SetMockTickerCreator(mtc)

	ticker := dogtime.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.Chan():
			return
		default:
			mt, err := mtc.GetMockTicker(0)
			if err != nil {
				t.Fatal(err)
			}
			mt.Fire()
		}
	}
}
