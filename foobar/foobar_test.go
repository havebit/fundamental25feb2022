package foobar_test

import (
	"testing"

	"gitlab.com/cjexpress/tildi/signac/learngo/foobar"
)

func TestFooBarGiven1WantsString1(t *testing.T) {
	given := 1
	wants := "1"

	result := foobar.Say(given)

	if wants != result {
		t.Errorf("given %d wants %q but got %q\n", given, wants, result)
	}
}

func TestFooBarGiven2WantsString2(t *testing.T) {
	given := 2
	wants := "2"

	result := foobar.Say(given)

	if wants != result {
		t.Errorf("given %d wants %q but got %q\n", given, wants, result)
	}
}
