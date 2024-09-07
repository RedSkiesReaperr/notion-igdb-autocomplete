package core

import (
	"fmt"
	"log"
	"notion-igdb-autocomplete/config"
	"notion-igdb-autocomplete/igdb"
	"notion-igdb-autocomplete/notion"
	"time"

	"github.com/RedSkiesReaperr/howlongtobeat"
	"github.com/jomei/notionapi"
)

//TODO: Should not start TimeToBeat task if name matches pattern "Not Found (GAME_NAME)"
//TODO: Kill all core goroutine at Stop()

type Core struct {
	Config                  *config.Config
	notion                  *notion.Client
	igdb                    *igdb.Client
	hltb                    *howlongtobeat.Client
	Tasks                   []*Task
	TasksQueue              chan *Task
	TasksNotificationsQueue chan string
	IsRunning               bool
}

func New(config *config.Config) (*Core, error) {
	newCore := Core{
		Config:                  config,
		Tasks:                   []*Task{},
		TasksQueue:              make(chan *Task, 20),
		TasksNotificationsQueue: make(chan string),
		IsRunning:               false,
	}

	return &newCore, nil
}

func (c *Core) Initialize() error {
	log.Println("Initializing core...")
	if err := c.initializeNotion(); err != nil {
		return fmt.Errorf("notion client initialize: %v", err)
	}

	if err := c.initializeIGDB(); err != nil {
		return fmt.Errorf("igdb client initialize: %v", err)
	}

	if err := c.initializeHLTB(); err != nil {
		return fmt.Errorf("hltb client initialize: %v", err)
	}

	log.Println("Core initialized")
	return nil
}

func (c *Core) Launch(selfReadNotifications bool) {
	if c.IsRunning {
		log.Println("Core already running!")
		return
	}

	log.Println("Launching core...")
	if selfReadNotifications {
		go c.readTasksNotifications()
	}

	c.IsRunning = true
	log.Println("Core launched")
	go c.watch()
	c.runTasks()
}

func (c *Core) readTasksNotifications() {
	for range c.TasksNotificationsQueue {
		<-c.TasksNotificationsQueue
	}
}

func (c *Core) watch() error {
	for range time.Tick(time.Duration(c.Config.RefreshDelay) * time.Second) {
		if !c.IsRunning {
			return nil
		}

		emptyGamesInfos, err := c.notion.Database(c.Config.NotionPageID).GetEmptyGamesEntries()
		if err != nil {
			return err
		}

		for _, entry := range emptyGamesInfos {
			title := entry.Properties["Title"].(*notionapi.TitleProperty).Title[0].Text.Content
			task := NewTask(GameInfosTask, title, entry.ID.String())

			if err := c.registerTask(task); err == nil {
				c.TasksQueue <- task
				task.QueuedAt = time.Now()
				c.TasksNotificationsQueue <- fmt.Sprintf("Registered task %s\n", task.Id)
			}
		}

		emptyTimeToBeat, err := c.notion.Database(c.Config.NotionPageID).GetEmptyTimeToBeatEntries()
		if err != nil {
			return err
		}

		for _, entry := range emptyTimeToBeat {
			title := entry.Properties["Title"].(*notionapi.TitleProperty).Title[0].Text.Content
			task := NewTask(TimeToBeatTask, title, entry.ID.String())

			if err := c.registerTask(task); err == nil {
				c.TasksQueue <- task
				task.QueuedAt = time.Now()
				c.TasksNotificationsQueue <- fmt.Sprintf("Registered task %s\n", task.Id)
			}
		}
	}

	return nil
}

func (c *Core) RetryTask(task *Task) {
	newTask := NewTask(task.Type, task.Query, task.NotionId)

	oldTaskIndex := -1
	for i, t := range c.Tasks {
		if t.Id == task.Id {
			oldTaskIndex = i
		}
	}

	if oldTaskIndex != -1 {
		c.Tasks = append(c.Tasks[:oldTaskIndex], c.Tasks[oldTaskIndex+1:]...) // Remove task from tasks list
	}

	task = nil // Remove task from memory

	if err := c.registerTask(newTask); err == nil {
		c.TasksQueue <- newTask
		newTask.QueuedAt = time.Now()
		c.TasksNotificationsQueue <- fmt.Sprintf("Registered task %s\n", newTask.Id)
	}
}

func (c *Core) runTasks() {
	for task := range c.TasksQueue {
		if task.Status == WaitingTask {
			go task.Run(c.TasksNotificationsQueue, c.notion, c.igdb, c.hltb)
		}
	}
}

func (c *Core) Stop() error {
	log.Println("Stopping core...")
	c.IsRunning = false
	log.Println("Core stopped")

	return nil
}
