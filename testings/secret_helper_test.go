package main

import "testing"

func TestDoSecret(t *testing.T) {

	mhelper := &MockSecretHelper{}
	got := DoSecret(mhelper)

	expected := "random number: 33, gender is gay: true, dog says: woff"
	if got != expected {
		t.Errorf("got: %s but expected: %s", got, expected)
	}

}
