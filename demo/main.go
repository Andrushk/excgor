package main

import (
	"fmt"
	"time"
	"github.com/andrushk/excgor"
)

const (
	hardDuration  = 10000000
	waitGoRoutine = 1000000
)

func main()  {
	ex := new(excgor.Excess)

	//can do job
	run(ex, "Andrey", job)
	time.Sleep(waitGoRoutine)

	//have no chances, will be skipped
	run(ex, "Ivan", job)
	run(ex, "Roman", job)

	//wait when "Andrey" will finish his work
	time.Sleep(hardDuration*2)

	//can do job
	run(ex, "Anton" ,job)
	time.Sleep(waitGoRoutine)

	//will be skipped
	run(ex, "Anna",job)
	run(ex, "Elena", job)

	time.Sleep(waitGoRoutine)
}

func job()  {
	fmt.Println("....do some work...")

	//do hard work
	time.Sleep(hardDuration)
}

func run(ex *excgor.Excess, workerName string, f func())  {
	go func() {
		ex.Do(func() {
			fmt.Printf("%v can do his job:", workerName)
			f()
		})
	}()
}
