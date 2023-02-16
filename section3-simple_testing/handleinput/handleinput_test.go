package handleinput

import (
	"bytes"
	"testing"
)

func TestReadUserInput(t *testing.T) {
	doneChan := make(chan bool)

	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go ReadUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
