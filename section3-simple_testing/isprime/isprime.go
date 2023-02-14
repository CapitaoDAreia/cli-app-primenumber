package isprime

import "fmt"

func IsPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not a prime number by definition!", n)
	}

	if n < 0 {
		return false, "Negative numbers is not prime by definition!"
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a priime number because it is divisible by %d.", n, i)
		}
	}

	return true, fmt.Sprintf("%d is probably a prime number.", n)
}
