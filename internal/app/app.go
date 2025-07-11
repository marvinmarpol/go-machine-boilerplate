package app

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

// Task represents a scheduled task
type Task struct {
	ID       string
	RunAt    time.Time
	Priority int
	Index    int // Required by heap.Interface
}

// TaskQueue implements heap.Interface
type TaskQueue []*Task

func (pq TaskQueue) Len() int { return len(pq) }

func (pq TaskQueue) Less(i, j int) bool {
	if pq[i].RunAt.Equal(pq[j].RunAt) {
		return pq[i].Priority > pq[j].Priority // higher priority first
	}
	return pq[i].RunAt.Before(pq[j].RunAt) // earlier time first
}

func (pq TaskQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *TaskQueue) Push(x interface{}) {
	n := len(*pq)
	task := x.(*Task)
	task.Index = n
	*pq = append(*pq, task)
}
func (pq *TaskQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	task := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return task
}

// TaskScheduler manages scheduled tasks
type TaskScheduler struct {
	mu    sync.Mutex
	tasks TaskQueue
}

func NewTaskScheduler() *TaskScheduler {
	pq := make(TaskQueue, 0)
	heap.Init(&pq)
	return &TaskScheduler{
		tasks: pq,
	}
}

func (ts *TaskScheduler) ScheduleTask(id string, delaySeconds int, priority int) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	runAt := time.Now().Add(time.Duration(delaySeconds) * time.Second)
	task := &Task{
		ID:       id,
		RunAt:    runAt,
		Priority: priority,
	}
	heap.Push(&ts.tasks, task)
	fmt.Printf("Task '%s' scheduled at %v\n", id, runAt)
}

func (ts *TaskScheduler) RunDueTasks() {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	now := time.Now()
	for ts.tasks.Len() > 0 {
		task := ts.tasks[0]
		if task.RunAt.After(now) {
			break
		}
		heap.Pop(&ts.tasks)
		fmt.Printf("Executing task: %s\n", task.ID)
	}
}

func Run() error {

	scheduler := NewTaskScheduler()
	scheduler.ScheduleTask("email", 5, 10)
	scheduler.ScheduleTask("report", 3, 20)
	scheduler.ScheduleTask("cleanup", 3, 5)

	time.Sleep(5 * time.Second)
	scheduler.RunDueTasks()

	return nil
}
