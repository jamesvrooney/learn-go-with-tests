package main

import "testing"

func TestPersonalHello(t *testing.T) {
    actual := Hello("James")
    expected := "Hello, James!"

    if actual != expected {
        t.Errorf("actual %q - expected %q", actual, expected)
    }
}
