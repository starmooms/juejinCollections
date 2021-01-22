package tool

import (
	"sync"
)

type WaitGroup struct {
	limt     int
	wg       sync.WaitGroup
	mutex    *sync.Mutex
	taskList []func()
	runNum   int
}

func NewWaitGroup(limt int) *WaitGroup {
	return &WaitGroup{
		limt:  limt,
		mutex: &sync.Mutex{},
	}
}

func (w *WaitGroup) Add(cb ...func()) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.taskList = append(w.taskList, cb...)
	if w.runNum < w.limt {
		w.runNum += 1
		w.wg.Add(1)
		go w.startTask(w.runNum)
	}
}

func (w *WaitGroup) getTask() func() {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	var task func() = nil
	list := w.taskList
	if len(list) > 0 {
		task = list[0]
		w.taskList = list[1:]
	}
	return task
}

func (w *WaitGroup) startTask(id int) {
	task := w.getTask()
	if task == nil {
		w.mutex.Lock()
		defer w.mutex.Unlock()
		w.runNum -= 1
		w.wg.Done()
		return
	}
	task()
	w.startTask(id)
}

func (w *WaitGroup) Wait() {
	w.wg.Wait()
}
