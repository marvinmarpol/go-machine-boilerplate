package app

import (
	"testing"
	"time"
)

func TestScheduleAndRunDueTasks(t *testing.T) {
	scheduler := NewTaskScheduler()

	scheduler.ScheduleTask("low-priority", 1, 1)
	scheduler.ScheduleTask("high-priority", 1, 10)

	time.Sleep(2 * time.Second)

	// Capture stdout with a hook if needed; here we just trust output manually
	scheduler.RunDueTasks()

	if scheduler.tasks.Len() != 0 {
		t.Errorf("Expected all tasks to be executed")
	}
}

func TestRunDueTasksEmpty(t *testing.T) {
	scheduler := NewTaskScheduler()
	scheduler.RunDueTasks() // should not panic
}
