package test

import "testing"

func SimpleAssertion(t *testing.T, got, expected any) {

	if expected != got {
		t.Errorf("got %#v but expected %#v", got, expected)
	}

}
