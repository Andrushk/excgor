package excgor

import (
	"testing"
	"sync/atomic"
)

func realRun(t *testing.T, ex *Excess, counter *int32, waitRoot, quit chan bool) {
	waitRoot <- true

	isExecuted := ex.Do(func() {
		atomic.AddInt32(counter, 1)
		<-quit
	})

	if !isExecuted {
		t.Errorf("expected instance did not execute")
	}

	waitRoot <- true
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

// Test only one instance without excess
func TestWithoutExcess(t *testing.T) {
	ex := new(Excess)

	if ex.inProcess != 0 {
		t.Fatalf("inProcess should be 0, actual %v", ex.inProcess)
	}

	if ex.getRealMax() != 1 {
		t.Fatalf("getRealMax should be 1, actual %v", ex.getRealMax)
	}

	isProcessed := false
	quit := make(chan bool)
	go func() {
		isExecuted := ex.Do(func() {
			isProcessed = true

			if ex.inProcess != 1 {
				t.Fatalf("inProcess should be 1, actual %v", ex.inProcess)
			}
			if ex.getRealMax() != 1 {
				t.Fatalf("getRealMax should be 1, actual %v", ex.getRealMax)
			}
		})

		<-quit

		if !isExecuted {
			t.Errorf("expected instance did not execute")
		}
	}()


	// wait when instance will finish
	quit <- true

	if !isProcessed {
		t.Fatalf("isProcessed should be true, actual %v", isProcessed)
	}

	if ex.inProcess != 0 {
		t.Fatalf("inProcess should be 0, actual %v", ex.inProcess)
	}

	if ex.getRealMax() != 1 {
		t.Fatalf("getRealMax should be 1, actual %v", ex.getRealMax)
	}
}

// Test one work instance and many excess
func TestOneAndManyExcess(t *testing.T) {
	ex := new(Excess)
	counter := new(int32)
	waitRoot := make(chan bool)
	quit := make(chan bool)
	wait := make(chan bool)

	go realRun(t, ex, counter, waitRoot, quit)

	// wait until right instance will start
	<-waitRoot

	//provoke to execute more than one instance at the same time
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

	// wait until right instance will finish
	<-waitRoot
}

//func TestFiveTogether(t *testing.T) {
//	ex := new(Excess)
//	ex.SetMax(5)
//	counter := new(int32)
//	quit := make(chan bool)
//	wait := make(chan bool)
//
//	const realN = 5
//	for i := 0; i < realN; i++ {
//		go realRun(t, ex, counter, quit)
//	}
//
//	//provoke to execute more than five instance at the same time
//	const excessN = 10
//	for i := 0; i < excessN; i++ {
//		go excessRun(t, ex, counter, wait)
//	}
//
//	for i := 0; i < excessN; i++ {
//		<-wait
//	}
//
//	for i:=0; i< realN; i++{
//		quit <- true
//	}
//
//	if *counter != 5 {
//		t.Errorf("check sum is wrong: %d is not 5", *counter)
//	}
//}
