package igdb

import (
	"fmt"
	"strings"
)

// NewSearchQuery return a valid IGDB query for a search and some fields
func NewSearchQuery(search string, fields ...string) string {
	return fmt.Sprintf(`search "%s";fields %s;`, search, strings.Join(fields, ","))
}
