package logger

import (
	"os"
	"os/signal"
)

type exitTask struct {
	hooks []func()
}

func (h *exitTask) Add(task func()) {
	h.hooks = append(h.hooks, task)
}

func (h *exitTask) Start() {
	for _, task := range h.hooks {
		task()
	}
}

var ExitHook = &exitTask{}

func init() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)

	go func() {
		select {
		case sig := <-ch:
			ExitHook.Start()
			if Logger != nil {
				Logger.Info("exit..", sig)
			}
			os.Exit(0)
		}
	}()
}
