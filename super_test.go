package main

import (
	"reflect"
	"testing"
)

func TestSomething(t *testing.T) {
	got := 8
	want := 8
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Beeeer")
	}
}

func TestWront(t *testing.T) {
	t.Fatalf("It should fail")
}
