package core

import (
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	result := NewTask(GameInfosTask, "my_query", "my_notionid")

	if result.Type != GameInfosTask {
		t.Errorf("Unexpected Type: expected:%v, got:%v", GameInfosTask, result.Type)
	}
	if result.Query != "my_query" {
		t.Errorf("Unexpected Query: expected:%v, got:%v", "my_query", result.Query)
	}
	if result.NotionId != "my_notionid" {
		t.Errorf("Unexpected NotionId: expected:%v, got:%v", "my_notionid", result.NotionId)
	}
	if result.CreatedAt.IsZero() != false {
		t.Errorf("Unexpected CreatedAt: expected:%v, got:%v", false, result.CreatedAt)
	}
	if result.QueuedAt.IsZero() != true {
		t.Errorf("Unexpected QueuedAt: expected:%v, got:%v", true, result.QueuedAt)
	}
	if result.StartedAt.IsZero() != true {
		t.Errorf("Unexpected StartedAt: expected:%v, got:%v", true, result.StartedAt)
	}
	if result.EndedAt.IsZero() != true {
		t.Errorf("Unexpected EndedAt: expected:%v, got:%v", true, result.EndedAt)
	}
}

func TestElapsedWaiting(t *testing.T) {
	task := Task{
		Status:    WaitingTask,
		CreatedAt: time.Now().Add(-2 * time.Hour),
	}
	expectedElapsed := 2 * time.Hour
	elapsed := task.Elapsed()
	if elapsed < expectedElapsed || elapsed > expectedElapsed+time.Second {
		t.Errorf("Expected elapsed time to be close to %v, got %v", expectedElapsed, elapsed)
	}
}

func TestElapsedRunning(t *testing.T) {
	task := Task{
		Status:    RunningTask,
		StartedAt: time.Now().Add(-1 * time.Hour),
	}
	expectedElapsed := 1 * time.Hour
	elapsed := task.Elapsed()
	if elapsed < expectedElapsed || elapsed > expectedElapsed+time.Second {
		t.Errorf("Expected elapsed time to be close to %v, got %v", expectedElapsed, elapsed)
	}
}

func TestElapsedFinished(t *testing.T) {
	task := Task{
		Status:    FinishedTask,
		StartedAt: time.Now().Add(-3 * time.Hour),
		EndedAt:   time.Now().Add(-1 * time.Hour),
	}
	expectedElapsed := 2 * time.Hour
	elapsed := task.Elapsed()
	if elapsed < expectedElapsed || elapsed > expectedElapsed+time.Second {
		t.Errorf("Expected elapsed time to be close to %v, got %v", expectedElapsed, elapsed)
	}
}

func TestElapsedFallback(t *testing.T) {
	task := Task{Status: -1}
	expectedElapsed := time.Duration(0)
	elapsed := task.Elapsed()
	if elapsed < expectedElapsed || elapsed > expectedElapsed+time.Second {
		t.Errorf("Expected elapsed time to be close to %v, got %v", expectedElapsed, elapsed)
	}
}
