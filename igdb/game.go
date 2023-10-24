package igdb

import (
	"fmt"

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

type Games []*Game
type Platforms []Platform
type Franchises []Franchise
type Genres []Genre

func (g *Game) CoverURL() string {
	return fmt.Sprintf("https://images.igdb.com/igdb/image/upload/t_cover_big/%s.png", g.Cover.ImageID)
}

func (g *Game) NotionPlatforms() (platforms []notionapi.Option) {
	for _, platform := range g.Platforms {
		platforms = append(platforms, notionapi.Option{Name: platform.Name})
	}

	return
}

func (g *Game) NotionGenres() (genres []notionapi.Option) {
	for _, genre := range g.Genres {
		genres = append(genres, notionapi.Option{Name: genre.Name})
	}

	return
}

func (g *Game) NotionFranchises() (franchises []notionapi.Option) {
	for _, franchise := range g.Franchises {
		franchises = append(franchises, notionapi.Option{Name: franchise.Name})
	}

	return
}
