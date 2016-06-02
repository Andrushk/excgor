package excgor

import (
	"testing"
	"sync/atomic"
)

func realRun(t *testing.T, ex *Excess, counter *int32, quit chan bool) {
	isExecuted := ex.Do(func() {
		atomic.AddInt32(counter, 1)
		<-quit
	})

	if !isExecuted {
		t.Errorf("expected instance did not execute")
	}
}

func excessRun(t *testing.T, ex *Excess, counter *int32, wait chan bool) {
	isExecuted := ex.Do(func() {
		atomic.AddInt32(counter, 1)
	})

	if isExecuted {
		t.Errorf("excess instance should not work")
	}

	wait <- true
}

func TestOne(t *testing.T) {
	ex := new(Excess)
	counter := new(int32)
	quit := make(chan bool)
	wait := make(chan bool)

	go realRun(t, ex, counter, quit)

	const n = 10
	for i := 0; i < n; i++ {
		go excessRun(t, ex, counter, wait)
	}

	for i := 0; i < n; i++ {
		<-wait
	}

	quit <- true

	if *counter != 1 {
		t.Errorf("more than one instance was executed: %d is not 1", *counter)
	}
}

