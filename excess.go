package excgor

import (
	"sync"
	"log"
)

type Excess struct {
	m          sync.Mutex

	// кол-во процессов, выполняющихся в данный момент
	inProcess  uint32

	// максимальное кол-во процессов, которые могут выполнятся одновременно
	maxProcess uint32
}


func (e *Excess) Do(f func()) bool {
	//if !e.addProcess() {
	//	return false
	//}
	e.inProcess++

	f()

	e.inProcess--
	//log.Println("before odd")
	//e.oddProcess()
	//log.Println("process odded")

	return true
}

func (e *Excess) SetMax(value uint32) {
	e.m.Lock()
	defer e.m.Unlock()

	e.maxProcess = value
}

func (e *Excess) addProcess() bool {
	//e.m.Lock()
	//defer e.m.Unlock()

	if e.inProcess >= e.getRealMax() {
		return false
	}

	e.inProcess += 1
	return true
}

func (e *Excess) oddProcess() {
	log.Println("start odding")

	//e.m.Lock()
	//defer e.m.Unlock()

	log.Println("insaid odd")

	if e.inProcess < 1 {
		return
	}
	e.inProcess -= 1
}

func (e *Excess) getRealMax() uint32 {
	//если maxProcess не инициализирован, то считаем его равным 1
	if e.maxProcess <= 1 {
		return 1
	}
	return e.maxProcess
}

// Prints debug info if DebugMode is set.
//func debug(m string, v interface{}) {
//	if debugMode {
//		log.Printf(m+":%+v", v)
//	}
//}
