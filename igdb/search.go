package igdb

import (
	"fmt"
	"strings"
)

func NewSearchQuery(search string, fields []string) string {
	return fmt.Sprintf(`search "%s";fields %s;`, search, strings.Join(fields, ","))
}
