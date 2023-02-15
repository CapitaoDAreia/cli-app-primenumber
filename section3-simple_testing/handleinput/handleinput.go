package handleinput

import (
	"bufio"
	"cli-app/introduction"
	"fmt"
	"os"
)

func ReadUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		res, done := CheckNumbers(scanner)
		if done {
			doneChan <- true
		}

		fmt.Println(res)
		introduction.Prompt()
	}
}
