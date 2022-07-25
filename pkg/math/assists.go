package math

func findDivisors(divisible int) []int {
	divisors := []int{}

	if divisible < 0 {
		for i := divisible; i <= 0; i++ {
			if divisible%i == 0 {
				divisors = append(divisors, i, -i)
			}
		}
	} else {
		for i := 1; i <= divisible; i++ {
			if divisible%i == 0 {
				divisors = append(divisors, i)
			}
		}
	}

	return divisors
}
