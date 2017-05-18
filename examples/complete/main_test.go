package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	// Run main method in a goroutine to make sure it runs.
	// Then let it die and just capture panics.
	go func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("PANIC: %s", r)
			}
		}()

		main()
	}()
}
