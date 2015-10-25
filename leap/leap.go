package leap

//TestVersion is the exercism test version
const TestVersion = 1

//IsLeapYear checks whether a year is a leap year or not
//it returns a boolean
func IsLeapYear(year int) bool {
	var divisor = 4
	// check for century year
	if year%100 == 0 {
		divisor = 400
	}
	return (year%divisor == 0)
}
