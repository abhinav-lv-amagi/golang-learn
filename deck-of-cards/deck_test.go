package deck

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	expected_length := 16

	if len(d) != expected_length {
		t.Errorf("Expected deck length of %d, but got: %d", expected_length, len(d))
	}
}