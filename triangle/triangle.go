package triangle

import "math"

type Kind int

const (
	Equ Kind = iota
	Iso
	Sca
	NaT
)

func KindFromSides(s1, s2, s3 float64) Kind {
	if !satisfiesTriangleInequality(s1, s2, s3) {
		return Kind(NaT)
	}
	if s1 == s2 && s2 == s3 {
		return Kind(Equ)
	}
	if s1 == s2 || s2 == s3 || s3 == s1 {
		return Kind(Iso)
	}
	return Kind(Sca)
}

func satisfiesTriangleInequality(s1, s2, s3 float64) bool {
	if s1 == 0 || s2 == 0 || s3 == 0 ||
		math.IsNaN(s1) || math.IsNaN(s2) || math.IsNaN(s3) ||
		math.IsInf(s1, 0) || math.IsInf(s2, 0) || math.IsInf(s3, 0) {
		return false
	}
	if s1+s2 <= s3 || s2+s3 <= s1 || s3+s1 <= s2 {
		return false
	}
	return true
}
