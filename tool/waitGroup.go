package tool

import (
	"sync"
)

type WaitGroup struct {
	limt       int
	wg         sync.WaitGroup
	mutex      *sync.Mutex
	taskList   []func()
	runNum     int
	TaskRunLen int
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
	newTaskLen := len(cb)
	w.taskList = append(w.taskList, cb...)

	for w.runNum < w.limt && newTaskLen > 0 {
		w.runNum += 1
		newTaskLen -= 1
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
		w.TaskRunLen++
	}
	return task
}

func (w *WaitGroup) startTask(id int) {
	for task := w.getTask(); task != nil; task = w.getTask() {
		task()
	}
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.runNum -= 1
	w.wg.Done()
	return
}

func (w *WaitGroup) GetLock(cb func()) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	cb()
}

func (w *WaitGroup) Wait() {
	w.wg.Wait()
}
