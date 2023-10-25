package igdb

import (
	"reflect"
	"testing"

	"github.com/jomei/notionapi"
)

func TestCoverURL(t *testing.T) {
	game := createFullfilledGame()
	result := game.CoverURL()
	expected := "https://images.igdb.com/igdb/image/upload/t_cover_big/thisIsACoverID.png"

	if result != expected {
		t.Errorf("Unexpected coverUrl: expected:%s, got:%s", expected, result)
	}
}

func TestNotionPlatformsFullfilled(t *testing.T) {
	game := createFullfilledGame()
	result := game.NotionPlatforms()
	expected := []notionapi.Option{
		{Name: "Wii"},
		{Name: "PC"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected notionPlatforms: expected:%s, got:%s", expected, result)
	}
}

func TestNotionPlatformsEmpty(t *testing.T) {
	game := createFullfilledGame()
	game.Platforms = nil

	result := game.NotionPlatforms()
	expected := []notionapi.Option{}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected notionPlatforms: expected:%s, got:%s", expected, result)
	}
}

func TestNotionGenresFullfilled(t *testing.T) {
	game := createFullfilledGame()
	result := game.NotionGenres()
	expected := []notionapi.Option{
		{Name: "Roleplay"},
		{Name: "Shooter"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected notionGenres: expected:%s, got:%s", expected, result)
	}
}

func TestNotionGenresEmpty(t *testing.T) {
	game := createFullfilledGame()
	game.Genres = nil

	result := game.NotionGenres()
	expected := []notionapi.Option{}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected notionGenres: expected:%s, got:%s", expected, result)
	}
}

func TestNotionFranchisesFullfilled(t *testing.T) {
	game := createFullfilledGame()
	result := game.NotionFranchises()
	expected := []notionapi.Option{
		{Name: "An awesome franchise"},
		{Name: "Another awesome franchise"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected notionFranchises: expected:%s, got:%s", expected, result)
	}
}

func TestNotionFranchisesEmpty(t *testing.T) {
	game := createFullfilledGame()
	game.Franchises = nil

	result := game.NotionFranchises()
	expected := []notionapi.Option{}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected notionFranchises: expected:%s, got:%s", expected, result)
	}
}
func createFullfilledGame() Game {
	return Game{
		Id:          1,
		Name:        "This is a fullfilled game",
		ReleaseDate: 861199200,
		Platforms: Platforms{
			{Id: 1, Name: "Wii"},
			{Id: 2, Name: "PC"},
		},
		Franchises: Franchises{
			{Id: 3, Name: "An awesome franchise"},
			{Id: 4, Name: "Another awesome franchise"},
		},
		Genres: Genres{
			{Id: 5, Name: "Roleplay"},
			{Id: 6, Name: "Shooter"},
		},
		Cover: Cover{
			Id:      7,
			ImageID: "thisIsACoverID",
		},
	}
}
