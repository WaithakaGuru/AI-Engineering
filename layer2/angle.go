package matrix

// here we are going to implement the functions for Cos Sin and Tan

//  some math constants
const (
	Pi 		= 	3.14159265358979323846264338327950288
	ConvRad = 	57.2957795130823208767981548141051703  // 180 / Pi
)

// Get the Sine of an angle 
func Sin(deg Pt) Pt {
	x := ConvertDegToRadian(deg)
	// x3 := x * x *x 
	// x5 := x3 * x * x
	// x7 := x5 * x * x
	// x9 := x7 * x * x
	// return x - x3/6 + x5/120 - x7/5040 - x9/362880

	// better simpler method 
	term, sum := x, x
	for n:= 1; n<10; n++ {
		term *= -x * x/ Pt((2* n) * (2*n+1))
		sum  += term
	}
	return sum
}
// Get the Cosine of an angle
func Cos(deg Pt) Pt { 
	x := ConvertDegToRadian(deg)

	var term, sum Pt = 1.0 , 1.0 
	for n := 1; n< 10; n++ {
		term *= -x * x/ (2* Pt(n)-1) * (2*Pt(n))
		sum += term
	} 
	return sum
}
// Get the Tan value of an angle
func Tan(deg Pt) Pt{
	deg = ConvertDegToRadian(deg)
	return Sin(deg) / Cos(deg )
}

func ConvertDegToRadian(deg Pt) Pt {
	return deg / ConvRad
}
func ConvertRadToDegree(rad Pt) Pt {
	return rad * ConvRad
}