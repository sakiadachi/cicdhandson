package main

import (
	"testing"
)

func TestGreetMeeting(t *testing.T) {
	want := "Hello, Taro"
	got := makeGreeting("Taro")
	if got != want {
		t.Errorf("got = %s; want %s", got, want)
	}
}
