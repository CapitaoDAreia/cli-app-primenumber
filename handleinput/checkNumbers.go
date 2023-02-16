package handleinput

import (
	"bufio"
	"cli-app/isprime"
	"strconv"
	"strings"
)

func CheckNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number!", false
	}

	_, msg := isprime.IsPrime(numToCheck)

	return msg, false
}
