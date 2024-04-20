package main

import "testing"

func TestMain(t *testing.T) {
	if "hello world" != getHello() {
		t.Fatal("failed test")
	}
}
