package main

import "testing"

func TestAllocViewF(t *testing.T) {

	got := AllocViewF()
	expected := "50"

	t.Run("MA", func(t *testing.T) {
		if got > expected {
			t.Errorf("got %s but expected %s", got, expected)
		}
	})
	t.Run("LA", func(t *testing.T) {
		if got < expected {
			t.Errorf("got %s but expected %s", got, expected)
		}
	})

}

func TestAllocViewG(t *testing.T) {

	got := AllocViewG()
	expected := "4 * 1024"

	t.Run("MakeMA", func(t *testing.T) {
		if got > expected {
			t.Errorf("got %s but expected %s", got, expected)
		}
	})
	t.Run("MakeLA", func(t *testing.T) {
		if got < expected {
			t.Errorf("got %s but expected %s", got, expected)
		}
	})

}
