// main_test.go
package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	sanityCheck(t)
}

func sanityCheck(t *testing.T) {
	t.Logf("Sanity Check")
	var testString = "expect 1 to be 1"
	if 1 == 1 {
		t.Logf(testString)
	} else {
		t.Errorf(testString)
	}
}
