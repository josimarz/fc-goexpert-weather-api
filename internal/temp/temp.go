package temp

type Temp struct {
	C float64
	F float64
	K float64
}

func CToF(c float64) float64 {
	return c*9/5 + 32
}

func CToK(c float64) float64 {
	return c + 273.15
}

func FToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func FToK(f float64) float64 {
	return (f-32)*5/9 + 273.15
}

func KToC(k float64) float64 {
	return k - 273.15
}

func KToF(k float64) float64 {
	return (k-273.15)*9/5 + 32
}
