package leap

const TestVersion = 1

func IsLeapYear(year int) bool {
	var divisor int = 4
	// check for century year
	if year%100 == 0 {
		divisor = 400
	}
	return (year%divisor == 0)
}
