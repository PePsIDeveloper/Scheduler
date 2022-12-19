package scheduler

import (
	"fmt"
	"time"
)
type Task func()
type Scheduler struct {
	tasks []Task
}
func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: []Task{},
	}
}

func (s *Scheduler) Schedule(task Task, interval time.Duration) {
	s.tasks = append(s.tasks, task)
	go func() {
		for {
			time.Sleep(interval)
			task()
		}
	}()
}
func (s *Scheduler) Run() {
	for _, task := range s.tasks {
		task()
	}
}
