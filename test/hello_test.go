package main

import "testing"

func TestHello(t *testing.T) {
	want := "hello, world"
	got := chello()
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}