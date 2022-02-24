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

func TestFooBarGiven3WantsStringFoo(t *testing.T) {
	given := 3
	wants := "Foo"

	result := foobar.Say(given)

	if wants != result {
		t.Errorf("given %d wants %q but got %q\n", given, wants, result)
	}
}

func TestFooBarGiven4WantsString4(t *testing.T) {
	given := 4
	wants := "4"

	result := foobar.Say(given)

	if wants != result {
		t.Errorf("given %d wants %q but got %q\n", given, wants, result)
	}
}
func TestFooBarGiven5WantsStringBar(t *testing.T) {
	given := 5
	wants := "Bar"

	result := foobar.Say(given)

	if wants != result {
		t.Errorf("given %d wants %q but got %q\n", given, wants, result)
	}
}
