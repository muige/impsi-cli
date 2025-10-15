package conversions

import (
	"math"
	"testing"
)

func TestFtToM(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"one foot", 1.0, 0.3048},
		{"ten feet", 10.0, 3.048},
		{"zero", 0.0, 0.0},
		{"negative", -5.0, -1.524},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FtToM(tt.input)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("FtToM(%f) = %f; want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMToFt(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"one meter", 1.0, 3.280839895},
		{"ten meters", 10.0, 32.80839895},
		{"zero", 0.0, 0.0},
		{"negative", -5.0, -16.40419948},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MToFt(tt.input)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("MToFt(%f) = %f; want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRoundTrip(t *testing.T) {
	tests := []float64{1.0, 10.0, 100.0, -5.5, 0.0}

	for _, val := range tests {
		t.Run("feet round trip", func(t *testing.T) {
			result := MToFt(FtToM(val))
			if math.Abs(result-val) > 0.0001 {
				t.Errorf("Round trip FtToM->MToFt(%f) = %f; want %f", val, result, val)
			}
		})

		t.Run("meter round trip", func(t *testing.T) {
			result := FtToM(MToFt(val))
			if math.Abs(result-val) > 0.0001 {
				t.Errorf("Round trip MToFt->FtToM(%f) = %f; want %f", val, result, val)
			}
		})
	}
}

func TestInToCm(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"one inch", 1.0, 2.54},
		{"ten inches", 10.0, 25.4},
		{"zero", 0.0, 0.0},
		{"negative", -5.0, -12.7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InToCm(tt.input)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("InToCm(%f) = %f; want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCmToIn(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"one cm", 1.0, 0.393701},
		{"ten cm", 10.0, 3.93701},
		{"zero", 0.0, 0.0},
		{"negative", -5.0, -1.968505},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CmToIn(tt.input)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("CmToIn(%f) = %f; want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestYdToM(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"one yard", 1.0, 0.9144},
		{"ten yards", 10.0, 9.144},
		{"zero", 0.0, 0.0},
		{"negative", -5.0, -4.572},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := YdToM(tt.input)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("YdToM(%f) = %f; want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMToYd(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"one meter", 1.0, 1.09361},
		{"ten meters", 10.0, 10.9361},
		{"zero", 0.0, 0.0},
		{"negative", -5.0, -5.46807},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MToYd(tt.input)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("MToYd(%f) = %f; want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMilesToKm(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"one mile", 1.0, 1.60934},
		{"ten miles", 10.0, 16.0934},
		{"zero", 0.0, 0.0},
		{"negative", -5.0, -8.0467},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MilesToKm(tt.input)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("MilesToKm(%f) = %f; want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestKmToMiles(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"one km", 1.0, 0.621371},
		{"ten km", 10.0, 6.21371},
		{"zero", 0.0, 0.0},
		{"negative", -5.0, -3.106855},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KmToMiles(tt.input)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("KmToMiles(%f) = %f; want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestLbsToKg(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"one pound", 1.0, 0.453592},
		{"ten pounds", 10.0, 4.53592},
		{"zero", 0.0, 0.0},
		{"negative", -5.0, -2.26796},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LbsToKg(tt.input)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("LbsToKg(%f) = %f; want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestKgToLbs(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"one kg", 1.0, 2.20462},
		{"ten kg", 10.0, 22.0462},
		{"zero", 0.0, 0.0},
		{"negative", -5.0, -11.0231},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KgToLbs(tt.input)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("KgToLbs(%f) = %f; want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestOzToG(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"one ounce", 1.0, 28.3495},
		{"ten ounces", 10.0, 283.495},
		{"zero", 0.0, 0.0},
		{"negative", -5.0, -141.7475},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := OzToG(tt.input)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("OzToG(%f) = %f; want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGToOz(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"one gram", 1.0, 0.035274},
		{"100 grams", 100.0, 3.5274},
		{"zero", 0.0, 0.0},
		{"negative", -50.0, -1.7637},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GToOz(tt.input)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("GToOz(%f) = %f; want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGalToL(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"one gallon", 1.0, 3.78541},
		{"ten gallons", 10.0, 37.8541},
		{"zero", 0.0, 0.0},
		{"negative", -5.0, -18.92705},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GalToL(tt.input)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("GalToL(%f) = %f; want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestLToGal(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"one liter", 1.0, 0.264172},
		{"ten liters", 10.0, 2.64172},
		{"zero", 0.0, 0.0},
		{"negative", -5.0, -1.32086},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LToGal(tt.input)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("LToGal(%f) = %f; want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFlOzToMl(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"one fl oz", 1.0, 29.5735},
		{"ten fl oz", 10.0, 295.735},
		{"zero", 0.0, 0.0},
		{"negative", -5.0, -147.8675},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FlOzToMl(tt.input)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("FlOzToMl(%f) = %f; want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMlToFlOz(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"one ml", 1.0, 0.033814},
		{"100 ml", 100.0, 3.3814},
		{"zero", 0.0, 0.0},
		{"negative", -50.0, -1.6907},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MlToFlOz(tt.input)
			if math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("MlToFlOz(%f) = %f; want %f", tt.input, result, tt.expected)
			}
		})
	}
}
