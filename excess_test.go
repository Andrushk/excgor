package excgor

import (
	"testing"
	"sync/atomic"
)

type counter int32

func (c *counter) Increment() {
	atomic.AddInt32(*c, 1)
	//*c++
}

func run(t *testing.T, e *Excess, c *counter, f chan bool){
	c <- true
}

func TestOne(t *testing.T){
	c := new(counter)
	e := new(Excess)
	f := make(chan bool)

	const n = 10
	for i:=0; i<n; i++ {
		go run(t, e, c, f)
	}
	for i:=0;i<n; i++{
		<-c
	}
}
