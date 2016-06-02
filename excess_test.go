package excgor

import (
	"testing"
	"log"
)

func run(t *testing.T, ex *Excess, result chan int, status chan bool) {
	isExecuted := ex.Do(func() {
		//log.Println("ENTERED!")
		//select {
		//case <-result:
		//	return
		//}
		result <- 1
	})

	status <- isExecuted
}

func TestOne(t *testing.T) {
	ex := new(Excess)
	result := make(chan int)
	status := make(chan bool)

	const n = 10
	for i := 0; i < n; i++ {
		go run(t, ex, result, status)
	}

	totalExcess := 0
	for i := 0; i < n; i++ {
		isExecuted, ok := <-status

		log.Printf("isExecuted=%v, ok=%v", isExecuted, ok)

		if !ok {
			break
		}

		if isExecuted {
			continue
		}

		totalExcess++
	}

	totalExecuted := n - totalExcess
	if totalExecuted != 1 {
		t.Errorf("too many executed instances: %d is not 1", totalExecuted)
	}

	//result<-1
	log.Println("Aaaaaaaaaa")
	<-result
	<-status

	//if *counter != 1 {
	//	t.Errorf("once failed outside run: %d is not 1", *counter)
	//}
}
