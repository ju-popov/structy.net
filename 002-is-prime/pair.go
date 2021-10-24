package isprime

func Pair(n int64) bool {
	//nolint:gomnd
	if n < 2 {
		return false
	}

	//nolint:gomnd
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
