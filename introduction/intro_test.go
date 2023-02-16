package introduction

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestIntro(t *testing.T) {
	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	Intro()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "Enter a whole number,") {
		t.Errorf("intro text seems not to be correct: '%s'", string(out))
	}
}
