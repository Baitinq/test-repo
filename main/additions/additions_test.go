package main

import (
	"os"
	"testing"

	"ci-visibility-test-github/main/civisibility/integrations/gotesting"
)

func TestMain(m *testing.M) {
	// NOTE: This is the only needed thing to enable test visibility instrumentation.
	os.Exit(gotesting.RunM(m))
}

func Test_AddNumbers(t *testing.T) {
	ans := addNumbers(2, 5)

	if ans != 7 {
		t.Fatal("Expected 7 but instead got ", ans)
	}
}

func Test_AddNumbers2(t *testing.T) {
	ans := addNumbers(7, 5)

	if ans != 12 {
		t.Fatal("Expected 12 but instead got ", ans)
	}
}
