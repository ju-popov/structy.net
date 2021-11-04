package isprime

//nolint:varnamelen
func Naive(n int) bool {
	//nolint:gomnd
	if n < 2 {
		return false
	}

	//nolint:gomnd
	for i := int(2); i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
