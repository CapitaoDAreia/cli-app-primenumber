package introduction

import (
	"io"
	"os"
	"testing"
)

func TestPrompt(t *testing.T) {
	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	Prompt()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if string(out) != "> " {
		t.Errorf("incorrect prompt: expected '> ' but got %s", string(out))
	}
}
