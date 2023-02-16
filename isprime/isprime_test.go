package isprime

import (
	"testing"
)

func TestIsprime(t *testing.T) {
	//test call of the function for invalids scenarios and validate if the returns
	//are the expected

	tests := []struct {
		testCase          string
		number            int
		expectedMessage   string
		expectedBoolValue bool
	}{
		{
			testCase:          "when a number is probably prime",
			number:            2,
			expectedMessage:   "2 is probably a prime number.",
			expectedBoolValue: true,
		},
		{
			testCase:          "when a number is negative",
			number:            -1,
			expectedMessage:   "Negative numbers is not prime by definition!",
			expectedBoolValue: false,
		},
		{
			testCase:          "when zero",
			number:            0,
			expectedMessage:   "0 is not a prime number by definition!",
			expectedBoolValue: false,
		},
		{
			testCase:          "when one",
			number:            1,
			expectedMessage:   "1 is not a prime number by definition!",
			expectedBoolValue: false,
		},
		{
			testCase:          "when is not prime",
			number:            8,
			expectedMessage:   "8 is not a prime number because it is divisible by 2.",
			expectedBoolValue: false,
		},
	}

	for _, test := range tests {
		t.Run(test.testCase, func(t *testing.T) {
			expectedBool, message := IsPrime(test.number)

			//expected bool not equals returned bool
			if test.expectedBoolValue && !expectedBool {
				t.Errorf("%s: expected true but the return was false", test.testCase)
			}

			//expected message not equals returned message
			if test.expectedMessage != message {
				t.Errorf("%s: expected: '%s' but got '%s'", test.testCase, test.expectedMessage, message)
			}

		})
	}
}
