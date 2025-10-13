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
