package excgor

import (
	"sync"
)

type Excess struct {
	m          sync.Mutex

	// кол-во процессов, выполняющихся в данный момент
	inProcess  uint32

	// максимальное кол-во процессов, которые могут выполнятся одновременно
	maxProcess uint32
}

func (e Excess) Do(f func()) bool {
	if !e.addProcess() {
		return false
	}

	f()
	e.oddProcess()

	return true
}

func (e Excess) SetMax(value uint32) {
	e.m.Lock()
	defer e.m.Unlock()

	e.maxProcess = value
}

func (e Excess) addProcess() bool {
	e.m.Lock()
	defer e.m.Unlock()

	if e.inProcess >= e.getRealMax() {
		return false
	}

	e.inProcess += 1
	return true
}

func (e Excess) oddProcess() {
	e.m.Lock()
	defer e.m.Unlock()

	if e.inProcess < 1 {
		return
	}
	e.inProcess -= 1
}

func (e Excess) getRealMax() uint32 {
	//если maxProcess не инициализирован, то считаем его равным 1
	if e.maxProcess <= 1 {
		return 1
	}
	return e.maxProcess
}
