package temp

import "testing"

func TestTemp(t *testing.T) {
	temp := &Temp{
		C: 30.5,
		F: 86.9,
		K: 303.65,
	}
	t.Run("Equals", func(t *testing.T) {
		other := &Temp{
			C: 30.5,
			F: 86.9,
			K: 303.65,
		}
		if got := temp.Equals(other); !got {
			t.Errorf("Equals(%v) = %v, want %v", other, got, true)
		}
		other = &Temp{
			C: 30.8,
			F: 86.9,
			K: 303.65,
		}
		if got := temp.Equals(other); got {
			t.Errorf("Equals(%v) = %v, want %v", other, got, false)
		}
	})
}

func TestRoundFloat(t *testing.T) {
	tests := []struct {
		input     float64
		precision uint8
		want      float64
	}{
		{-29.367, 2, -29.37},
		{21.456, 2, 21.46},
		{21.452, 2, 21.45},
		{5.333948, 3, 5.334},
		{5.333948, 4, 5.3339},
	}
	for _, test := range tests {
		if got := roundFloat(test.input, test.precision); got != test.want {
			t.Errorf("roundFloat(%v, %v) = %v, want %v", test.input, test.precision, got, test.want)
		}
	}
}

func TestCToF(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{0, 32},
		{-20, -4},
		{-5, 23},
		{10, 50},
		{30, 86},
	}
	for _, test := range tests {
		if got := CToF(test.input); got != test.want {
			t.Errorf("CToF(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestCToK(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{0, 273.15},
		{-10, 263.15},
		{-5, 268.15},
		{10, 283.15},
		{30, 303.15},
	}
	for _, test := range tests {
		if got := CToK(test.input); got != test.want {
			t.Errorf("CToK(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestFToC(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{0, -17.78},
		{-10, -23.33},
		{-5, -20.56},
		{10, -12.22},
		{30, -1.11},
	}
	for _, test := range tests {
		if got := FToC(test.input); got != test.want {
			t.Errorf("FToC(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestFToK(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{0, 255.37},
		{-10, 249.82},
		{-5, 252.59},
		{10, 260.93},
		{30, 272.04},
	}
	for _, test := range tests {
		if got := FToK(test.input); got != test.want {
			t.Errorf("FToK(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestKToC(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{0, -273.15},
		{-10, -283.15},
		{-5, -278.15},
		{10, -263.15},
		{30, -243.15},
	}
	for _, test := range tests {
		if got := KToC(test.input); got != test.want {
			t.Errorf("KToC(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestKToF(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{0, -459.67},
		{-10, -477.67},
		{-5, -468.67},
		{10, -441.67},
		{30, -405.67},
	}
	for _, test := range tests {
		if got := KToF(test.input); got != test.want {
			t.Errorf("KToF(%v) = %v, want %v", test.input, got, test.want)
		}
	}
}
