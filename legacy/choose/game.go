package choose

import (
	"notion-igdb-autocomplete/igdb"
	"sort"

	"github.com/agnivade/levenshtein"
)

type gameComparisons []gameComparison
type gameComparison struct {
	index int
	game  *igdb.Game
}

// Implements interface sort.Interface
func (gc gameComparisons) Len() int {
	return len(gc)
}

// Implements interface sort.Interface
func (gc gameComparisons) Swap(i, j int) {
	gc[i], gc[j] = gc[j], gc[i]
}

// Implements interface sort.Interface
func (gc gameComparisons) Less(i, j int) bool {
	return gc[i].index < gc[j].index
}

// Game returns the best choice in the list for the given search
func Game(search string, list igdb.Games) *igdb.Game {
	var comparisons gameComparisons

	for _, game := range list {
		comparisons = append(comparisons, gameComparison{
			game:  game,
			index: levenshtein.ComputeDistance(search, game.Name),
		})
	}

	sort.Sort(comparisons)

	return comparisons[0].game
}
