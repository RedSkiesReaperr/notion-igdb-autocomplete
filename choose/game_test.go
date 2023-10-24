package choose

import (
	"notion-igdb-autocomplete/igdb"
	"sort"
	"testing"
)

func TestGameComparisonsSorting(t *testing.T) {
	comparisons := gameComparisons{
		gameComparison{index: 89},
		gameComparison{index: 4},
		gameComparison{index: 64},
		gameComparison{index: 5},
	}

	expected := []gameComparison{
		{index: 4},
		{index: 5},
		{index: 64},
		{index: 89},
	}

	sort.Sort(comparisons)

	for i, v := range comparisons {
		if v != expected[i] {
			t.Errorf("Unexpected comparison[%d]: expected=%v, got=%v", i, expected[i], v)
		}
	}
}

func TestGameChoice(t *testing.T) {
	search := "world oF warcraft"
	choices := igdb.Games{
		&igdb.Game{Name: "World of Warcraft: Battle for Azeroth"},
		&igdb.Game{Name: "World of Warcraft: Wrath of the Lich King"},
		&igdb.Game{Name: "World of Warcraft: Warlords of Draenor"},
		&igdb.Game{Name: "World of Warcraft: Legion"},
		&igdb.Game{Name: "World of Warcraft"},
		&igdb.Game{Name: "World of Warcraft: Dragonflight"},
		&igdb.Game{Name: "World of Warcraft: Cataclysm"},
		&igdb.Game{Name: "World of Warcraft: The Burning Crusade"},
		&igdb.Game{Name: "World of Warcraft: Mists of Pandaria"},
		&igdb.Game{Name: "World of Warcraft: Shadowlands"},
	}
	expectedName := "World of Warcraft"
	result := Game(search, choices)

	if result.Name != expectedName {
		t.Errorf("Unexpected choice: expected:%s, got:%s", expectedName, result.Name)
	}
}
