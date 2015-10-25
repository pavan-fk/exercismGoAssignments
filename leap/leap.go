package leap

//TestVersion is the exercism test version.
const TestVersion = 1

//IsLeapYear checks whether a year is a leap year or not.
func IsLeapYear(year int) bool {
	return ((year%400 == 0) || (year%100 != 0 && year%4 == 0))
}
