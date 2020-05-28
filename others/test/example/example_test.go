package main

import "testing"

func TestExample(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
			println("hi")
			result := Example("")
			expected := "blank name"
			if expected != result {
					t.Fatalf("failed Test")
			}
	})
	t.Run("Normal", func(t *testing.T) {
			result := Example("John")
			expected := "Hello John"
			if expected != result {
					t.Fatalf("failed Test")
			}
	})
}