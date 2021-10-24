package isprime

func Naive(n int64) bool {
	//nolint:gomnd
	if n < 2 {
		return false
	}

	//nolint:gomnd
	for i := int64(2); i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
