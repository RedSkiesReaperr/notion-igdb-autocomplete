package igdb

import (
	"fmt"
	"time"

	"github.com/jomei/notionapi"
)

type Game struct {
	Id          int        `json:"id,required"`
	Name        string     `json:"name,required"`
	Platforms   Platforms  `json:"platforms,required"`
	ReleaseDate int64      `json:"first_release_date,required"`
	Franchises  Franchises `json:"franchises,required"`
	Genres      Genres     `json:"genres,required"`
	Cover       Cover      `json:"cover,required"`
}

type Platform struct {
	Id   int    `json:"id,required"`
	Name string `json:"name,required"`
}

type Franchise struct {
	Id   int    `json:"id,required"`
	Name string `json:"name,required"`
}

type Genre struct {
	Id   int    `json:"id,required"`
	Name string `json:"name,required"`
}

type Cover struct {
	Id      int    `json:"id,required"`
	ImageID string `json:"image_id,required"`
}

type Platforms []Platform
type Franchises []Franchise
type Genres []Genre

type Games []*Game

func (g Games) String() string {
	str := ""
	for i, game := range g {
		gameName := fmt.Sprintf("'%s'", game.Name)

		if i == 0 || i == len(g)-1 {
			str += gameName
		} else {
			str += fmt.Sprintf(", %s", gameName)
		}
	}

	return str
}

func (g *Game) CoverURL() string {
	return fmt.Sprintf("https://images.igdb.com/igdb/image/upload/t_cover_big/%s.png", g.Cover.ImageID)
}

func (g *Game) NotionPlatforms() []notionapi.Option {
	platforms := make([]notionapi.Option, 0)

	for _, platform := range g.Platforms {
		platforms = append(platforms, notionapi.Option{Name: platform.Name})
	}

	return platforms
}

func (g *Game) NotionGenres() []notionapi.Option {
	genres := make([]notionapi.Option, 0)

	for _, genre := range g.Genres {
		genres = append(genres, notionapi.Option{Name: genre.Name})
	}

	return genres
}

func (g *Game) NotionFranchises() []notionapi.Option {
	franchises := make([]notionapi.Option, 0)

	for _, franchise := range g.Franchises {
		franchises = append(franchises, notionapi.Option{Name: franchise.Name})
	}

	return franchises
}

// Implements notion.PageUpdateRequester interface
func (g Game) UpdateRequest() notionapi.PageUpdateRequest {
	releaseDate := notionapi.Date(time.Unix(g.ReleaseDate, 0))

	return notionapi.PageUpdateRequest{
		Cover: &notionapi.Image{
			Type: "external",
			External: &notionapi.FileObject{
				URL: g.CoverURL(),
			},
		},
		Properties: notionapi.Properties{
			"Title": notionapi.TitleProperty{
				Type: notionapi.PropertyTypeTitle,
				Title: []notionapi.RichText{
					{Text: &notionapi.Text{Content: g.Name}},
				},
			},
			"Release date": notionapi.DateProperty{
				Type: notionapi.PropertyTypeDate,

				Date: &notionapi.DateObject{
					Start: &releaseDate,
				},
			},
			"Franchises": notionapi.MultiSelectProperty{
				Type:        notionapi.PropertyTypeMultiSelect,
				MultiSelect: g.NotionFranchises(),
			},
			"Genres": notionapi.MultiSelectProperty{
				Type:        notionapi.PropertyTypeMultiSelect,
				MultiSelect: g.NotionGenres(),
			},
			"Platforms": notionapi.MultiSelectProperty{
				Type:        notionapi.PropertyTypeMultiSelect,
				MultiSelect: g.NotionPlatforms(),
			},
		},
	}
}
