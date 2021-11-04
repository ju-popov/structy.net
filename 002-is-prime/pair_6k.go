package isprime

//nolint:varnamelen
func Pair6K(n int) bool {
	//nolint:gomnd
	if n < 2 {
		return false
	}

	//nolint:gomnd
	if (n == 2) || (n == 3) {
		return true
	}

	if (n%2 == 0) || (n%3 == 0) {
		return false
	}

	//nolint:gomnd
	for i := int(5); i*i <= n; i += 6 {
		if (n%i == 0) || (n%(i+2) == 0) {
			return false
		}
	}

	return true
}
