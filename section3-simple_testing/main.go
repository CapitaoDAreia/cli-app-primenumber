package main

import (
	"cli-app/handleinput"
	"cli-app/introduction"
	"fmt"
)

func main() {
	introduction.Intro()
	introduction.Prompt()

	doneChan := make(chan bool)

	go handleinput.ReadUserInput(doneChan)
	<-doneChan
	close(doneChan)

	fmt.Println("Goodbye!")
}
