package main

import "testing"

func TestHello(t *testing.T) {
	got := hello()
	want := "hello world"

	t.Errorf("i have %q but i want %q", got, want)
}
