package collatzconjecture

type NonPositiveError struct{}

func (err *NonPositiveError) Error() string {
	return "n must be a positive number."
}

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return -1, &NonPositiveError{}
	}

	count := 0
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		count++
	}
	return count, nil
}
