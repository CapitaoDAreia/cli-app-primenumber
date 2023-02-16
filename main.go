package main

import (
	"cli-app/handleinput"
	"cli-app/introduction"
	"fmt"
	"os"
)

func main() {
	introduction.Intro()
	introduction.Prompt()

	doneChan := make(chan bool)

	go handleinput.ReadUserInput(os.Stdin, doneChan)
	<-doneChan
	close(doneChan)

	fmt.Println("Goodbye!")
}
