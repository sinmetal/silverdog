package dogtime_test

import (
	"testing"
	"time"

	"github.com/sinmetal/silverdog/dogtime"
)

func TestNewTicker(t *testing.T) {
	timeout := time.NewTicker(1 * time.Second)
	defer timeout.Stop()

	ticker := dogtime.NewTicker(1 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.Chan():
			return
		case <-timeout.C:
			t.Error("timeout!!")
		}
	}
}
