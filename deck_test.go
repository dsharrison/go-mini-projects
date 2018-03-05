package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	exp := 52

	if len(d) != exp {
		t.Errorf("Expected deck length of %v, got %v", exp, len(d))
	}
}
