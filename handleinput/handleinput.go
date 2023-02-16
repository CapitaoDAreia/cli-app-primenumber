package handleinput

import (
	"bufio"
	"cli-app/introduction"
	"fmt"
	"io"
)

func ReadUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := CheckNumbers(scanner)
		if done {
			doneChan <- true
		}

		fmt.Println(res)
		introduction.Prompt()
	}
}
