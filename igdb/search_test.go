package igdb

import "testing"

func TestNewQuery(t *testing.T) {
	search := "i'm searching for this please"

	result := NewSearchQuery(search, "field1", "myField2", "field3")
	expected := `search "i'm searching for this please";fields field1,myField2,field3;`

	if result != expected {
		t.Errorf("Unexpected query: expected=%s, got=%s", expected, result)
	}
}
