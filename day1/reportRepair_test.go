package main

import (
	"testing"
)

func Test_subtract2020FromAll(t *testing.T) {
	result := findAnswer(map[int]struct{}{
		1721: {},
		979:  {},
		366:  {},
		299:  {},
		675:  {},
		1456: {},
	}, 2020)
	want := 514579

	if result != want {
		t.Errorf("findAnswer() = %v, want %v", result, want)
	}

	result = findAnswer(map[int]struct{}{
		1721: {},
		979:  {},
		366:  {},
		299:  {},
		675:  {},
		1456: {},
	}, 1041)
	want = 247050

	if result != want {
		t.Errorf("findAnswer() = %v, want %v", result, want)
	}
}
