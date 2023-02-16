package handleinput

import (
	"bufio"
	"strings"
	"testing"
)

func TestCheckNumbers(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedReturn string
	}{
		{
			name:           "Empty value",
			input:          "",
			expectedReturn: "Please enter a whole number!",
		},
		{
			name:           "Prime number",
			input:          "2",
			expectedReturn: "2 is probably a prime number.",
		},
		{
			name:           "Decimal",
			input:          "2.1",
			expectedReturn: "Please enter a whole number!",
		},
		{
			name:           "Negative",
			input:          "-1",
			expectedReturn: "Negative numbers is not prime by definition!",
		},
		{
			name:           "Not prime",
			input:          "8",
			expectedReturn: "8 is not a prime number because it is divisible by 2.",
		},
		{
			name:           "Exit",
			input:          "q",
			expectedReturn: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input := strings.NewReader(test.input)
			reader := bufio.NewScanner(input)

			res, _ := CheckNumbers(reader)

			if !strings.EqualFold(res, test.expectedReturn) {
				t.Error("error")
			}
		})
	}

}
