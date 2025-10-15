// Package conversions provides conversion functions to impsi cli
package conversions

// Distance and length conversions

func FtToM(x float64) float64 {
	return x * 0.3048
}

func MToFt(x float64) float64 {
	return x / 0.3048
}

func YdToM(x float64) float64 {
	return x * 0.9144
}

func MToYd(x float64) float64 {
	return x / 0.9144
}

func MilesToKm(x float64) float64 {
	return x * 1.60934
}

func KmToMiles(x float64) float64 {
	return x
}

func InToCm(x float64) float64 {
	return x * 2.54
}

func CmToIn(x float64) float64 {
	return x / 2.54
}

// Weight and mass conversions

func LbsToKg(x float64) float64 {
	return x * 0.453592
}

func KgToLbs(x float64) float64 {
	return x / 0.453592
}

func OzToG(x float64) float64 {
	return x * 28.3495
}

func GToOz(x float64) float64 {
	return x / 28.3495
}

// Volume conversions

func GalToL(x float64) float64 {
	return x * 3.78541
}

func LToGal(x float64) float64 {
	return x / 3.78541
}

func FlOzToMl(x float64) float64 {
	return x * 29.5735
}

func MlToFlOz(x float64) float64 {
	return x / 29.5735
}
