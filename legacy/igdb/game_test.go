package igdb

import (
	"reflect"
	"testing"

	"github.com/jomei/notionapi"
)

func TestCoverURL(t *testing.T) {
	game := Game{Cover: Cover{Id: 1, ImageID: "lich_king"}}
	result := game.CoverURL()
	expected := "https://images.igdb.com/igdb/image/upload/t_cover_big/lich_king.png"

	if result != expected {
		t.Errorf("Unexpected CoverUrl(): expected:%s, got:%s", expected, result)
	}
}

func TestNotionPlatforms(t *testing.T) {
	var game Game = Game{}
	var tests = []struct {
		platforms Platforms
		expected  []notionapi.Option
	}{
		{platforms: Platforms{}, expected: []notionapi.Option{}},
		{platforms: Platforms{{Id: 1, Name: "Wii"}, {Id: 2, Name: "PC"}}, expected: []notionapi.Option{{Name: "Wii"}, {Name: "PC"}}},
	}

	for _, tt := range tests {
		game.Platforms = tt.platforms
		result := game.NotionPlatforms()

		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("Unexpected NotionPlatforms(): expected:%s, got:%s", tt.expected, result)
		}
	}
}

func TestNotionGenres(t *testing.T) {
	var game Game = Game{}
	var tests = []struct {
		genres   Genres
		expected []notionapi.Option
	}{
		{genres: Genres{}, expected: []notionapi.Option{}},
		{genres: Genres{{Id: 5, Name: "Roleplay"}, {Id: 6, Name: "Shooter"}}, expected: []notionapi.Option{{Name: "Roleplay"}, {Name: "Shooter"}}},
	}

	for _, tt := range tests {
		game.Genres = tt.genres
		result := game.NotionGenres()

		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("Unexpected NotionGenres(): expected:%s, got:%s", tt.expected, result)
		}
	}
}

func TestNotionFranchises(t *testing.T) {
	var game Game = Game{}
	var tests = []struct {
		franchises Franchises
		expected   []notionapi.Option
	}{
		{franchises: Franchises{}, expected: []notionapi.Option{}},
		{franchises: Franchises{{Id: 3, Name: "An awesome franchise"}, {Id: 4, Name: "Another awesome franchise"}}, expected: []notionapi.Option{{Name: "An awesome franchise"}, {Name: "Another awesome franchise"}}},
	}

	for _, tt := range tests {
		game.Franchises = tt.franchises
		result := game.NotionFranchises()

		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("Unexpected NotionFranchises(): expected:%s, got:%s", tt.expected, result)
		}
	}
}
