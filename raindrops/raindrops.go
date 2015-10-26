package raindrops

import (
	"bytes"
	"strconv"
)

//Convert converts int to a string based on logic outlined in the problem.
func Convert(number int) string {
	hasAtleastOnePrimeFactor := false
	var buffer bytes.Buffer
	if number%3 == 0 {
		hasAtleastOnePrimeFactor = true
		buffer.WriteString("Pling")
	}
	if number%5 == 0 {
		hasAtleastOnePrimeFactor = true
		buffer.WriteString("Plang")
	}
	if number%7 == 0 {
		hasAtleastOnePrimeFactor = true
		buffer.WriteString("Plong")
	}
	if !hasAtleastOnePrimeFactor {
		buffer.WriteString(strconv.Itoa(number))
	}
	return buffer.String()
}
