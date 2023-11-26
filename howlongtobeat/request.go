package howlongtobeat

import (
	"strings"
)

type Request struct {
	SearchTerms []string
	PageSize    int
}

type body struct {
	SearchType    string        `json:"searchType,required"`
	SearchTerms   []string      `json:"searchTerms,required"`
	SearchPage    int           `json:"searchPage,required"`
	Size          int           `json:"size,required"`
	SearchOptions searchOptions `json:"searchOptions,required"`
}

type searchOptions struct {
	Games      searchOptionsGames `json:"games,required"`
	Users      searchOptionsUsers `json:"users,required"`
	Filter     string             `json:"filter,required"`
	Sort       int                `json:"sort,required"`
	Randomizer int                `json:"randomizer,required"`
}

type searchOptionsGames struct {
	UserId        int                         `json:"userId,required"`
	Platform      string                      `json:"platform,required"`
	SortCategory  string                      `json:"sortCategory,required"`
	RangeCategory string                      `json:"rangeCategory,required"`
	RangeTime     searchOptionsGamesRangeTime `json:"rangeTime,required"`
	Gameplay      searchOptionsGamesGameplay  `json:"gameplay,required"`
	Modifier      string                      `json:"modifier,required"`
}

type searchOptionsUsers struct {
	SortCategory string `json:"sortCategory"`
}

type searchOptionsGamesRangeTime struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type searchOptionsGamesGameplay struct {
	Perspective string `json:"perspective"`
	Flow        string `json:"flow"`
	Genre       string `json:"genre"`
}

func NewRequest(gameName string) Request {
	var terms []string = strings.Split(gameName, " ")

	return Request{
		SearchTerms: terms,
		PageSize:    1,
	}
}

func (r Request) Body() body {
	return body{
		SearchType:  "games",
		SearchTerms: r.SearchTerms,
		SearchPage:  1,
		Size:        r.PageSize,
		SearchOptions: searchOptions{
			Games: searchOptionsGames{
				UserId:        0,
				Platform:      "",
				SortCategory:  "name",
				RangeCategory: "main",
				RangeTime:     searchOptionsGamesRangeTime{Min: 0, Max: 0},
				Gameplay:      searchOptionsGamesGameplay{Perspective: "", Flow: "", Genre: ""},
				Modifier:      "",
			},
			Users:      searchOptionsUsers{SortCategory: "postcount"},
			Filter:     "",
			Sort:       0,
			Randomizer: 0,
		},
	}
}
