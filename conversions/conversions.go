// Package conversions provides conversion functions to impsi cli
package conversions

func FtToM(x float64) float64 {
	return x * 0.3048
}

func MToFt(x float64) float64 {
	return x / 0.3048
}
