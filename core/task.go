package core

import (
	"fmt"
	"log"
	"notion-igdb-autocomplete/choose"
	"notion-igdb-autocomplete/howlongtobeat"
	"notion-igdb-autocomplete/igdb"
	"notion-igdb-autocomplete/notion"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id        uuid.UUID
	Type      TaskType
	Status    TaskStatus
	Query     string
	NotionId  string
	CreatedAt time.Time
	QueuedAt  time.Time
	StartedAt time.Time
	EndedAt   time.Time
	Error     error
}

func NewTask(tType TaskType, query string, notionId string) *Task {
	queryCleaner := strings.NewReplacer("{{", "", "}}", "")

	return &Task{
		Id:        uuid.New(),
		Type:      tType,
		Status:    WaitingTask,
		Query:     queryCleaner.Replace(query),
		NotionId:  notionId,
		CreatedAt: time.Now(),
		Error:     nil,
	}
}

func (t Task) String() string {
	return fmt.Sprintf("Id=%v Type=%s Query='%s' NotionId=%s", t.Id, TaskTypeString(t.Type), t.Query, t.NotionId)
}

func (t Task) Elapsed() time.Duration {
	switch t.Status {
	case WaitingTask:
		return time.Now().Sub(t.CreatedAt)
	case RunningTask:
		return time.Now().Sub(t.StartedAt)
	case FinishedTask:
		return t.EndedAt.Sub(t.StartedAt)
	default:
		return time.Duration(0)
	}
}

func (t *Task) Run(notifier chan<- string, notionClient *notion.Client, igdbClient *igdb.Client, hltbClient *howlongtobeat.Client) {
	t.Log("Starting")
	t.StartedAt = time.Now()
	t.Status = RunningTask
	notifier <- fmt.Sprintf("Started task %s\n", t.Id)

	switch t.Type {
	case GameInfosTask:
		if err := t.runGameInfos(notionClient, igdbClient); err != nil {
			t.Error = err
		}
	case TimeToBeatTask:
		if err := t.runTimeToBeat(notionClient, hltbClient); err != nil {
			t.Error = err
		}
	}

	t.EndedAt = time.Now()
	t.Status = FinishedTask
	notifier <- fmt.Sprintf("Finished task %s\n", t.Id)
	t.LogResume()
	t.Log("Finished")
}

func (t Task) Log(msg string) {
	header := fmt.Sprintf("[%s]", t.Id)

	log.Printf("%s %s\n", header, msg)
}

func (t Task) LogResume() {
	resume := fmt.Sprintf("task resume:\n....Id=%s\n....Type=%s\n....Status=%s\n....Query='%s'\n....NotionId=%s\n....Error=%v\n....CreatedAt=%s\n....QueuedAt=%s\n....StartedAt=%s\n....EndedAt=%s\n",
		t.Id, t.TypeString(), t.StatusString(), t.Query, t.NotionId,
		t.Error,
		t.CreatedAt.Format("2006-01-02 15:04:05"),
		t.QueuedAt.Format("2006-01-02 15:04:05"),
		t.StartedAt.Format("2006-01-02 15:04:05"),
		t.EndedAt.Format("2006-01-02 15:04:05"),
	)

	t.Log(resume)
}

func (t *Task) StatusString() string {
	return TaskStatusString(t.Status)
}

func (t *Task) TypeString() string {
	return TaskTypeString(t.Type)
}

func (t *Task) Succeed() bool {
	return t.Status == FinishedTask && t.Error == nil
}

func (t *Task) runGameInfos(notionClient *notion.Client, igdbClient *igdb.Client) error {
	// Fetch
	t.Log("Searching game...")
	query := igdb.NewSearchQuery(t.Query, "name", "platforms.name", "first_release_date", "franchises.name", "genres.name", "cover.image_id")
	results, err := igdbClient.SearchGame(query)
	if err != nil {
		return err
	}
	t.Log(fmt.Sprintf("%d games found: %s", len(results), results))

	foundGame := &igdb.Game{Name: fmt.Sprintf("Not found (%s)", t.Query)}
	if len(results) > 0 {
		foundGame = choose.Game(t.Query, results)
		t.Log(fmt.Sprintf("Chosen game: '%s'", foundGame.Name))
	}

	// Update
	t.Log("Updating Notion page...")
	_, err = notionClient.Page(t.NotionId).Update(foundGame)
	if err != nil {
		return err
	}

	return nil
}

func (t *Task) runTimeToBeat(notionClient *notion.Client, hltbClient *howlongtobeat.Client) error {
	// Fetch
	games, err := hltbClient.SearchGame(t.Query)
	if err != nil {
		return err
	}

	foundGame := &howlongtobeat.Game{Name: fmt.Sprintf("Not found (%s)", t.Query), CompletionMain: 0, CompletionPlus: 0, CompletionFull: 0}
	if len(games) > 0 {
		foundGame = &games[0]
	}

	// Update
	_, err = notionClient.Page(t.NotionId).Update(foundGame)
	if err != nil {
		return err
	}

	return nil
}
