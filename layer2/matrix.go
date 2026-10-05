package matrix

import (
	"math"
)

// Errors, for some common cases that may occur during matrix transforamtions and calcs

type MatError struct{
	errString string
}

func (m MatError) Error () string {
	return m.errString
}

func MError(err string) error {
	return MatError{errString: err}
}

// error statements 
const(
	Null Pt = -1000000.10001000
	ErrDETMatrixNotSquare = "Cannot find determinant of a non-square Matrix [Rows and Cols must be equal]"
	ErrINVMatrixNot2by2 = "Cannot find the 2X2 inverse [The matrix must have a shape of (2ROWS by 2COLS)]"
	ErrMulMatrixNotCompatible = "Failed to Multiply the matrices as they are incompatible [M1.Cols must be equal to M2.Rows]"
	ErrINVMatrixNotSquare = "Cannot find the Inverse of a Non-square Matrix [M.Cols Must be Equal to M.Rows]"
)


// Define a vector shape, a matrix shape, 
// define matrix transformations - shear, scale, rotation, inverse2X2, determinant, 
// define constants for all the above transformation

var(
	UnitMatrix *Matrix = NewMatrix([]Vector{{1,0},{0,1}})
	// Rotate45 *Matrix = NewMatrix([]Vector{{1,0},{0,1}})
	// Rotate90 *Matrix = NewMatrix([]Vector{{1,0},{0,1}})
	// Rotate135 *Matrix = NewMatrix([]Vector{{1,0},{0,1}})
	// Rotate180 *Matrix = NewMatrix([]Vector{{1,0},{0,1}})
	// Rotate225 *Matrix = NewMatrix([]Vector{{1,0},{0,1}})
	// Rotate270 *Matrix = NewMatrix([]Vector{{1,0},{0,1}})
	// Rotate315 *Matrix = NewMatrix([]Vector{{1,0},{0,1}})
	// ScaleY *Matrix = NewMatrix([]Vector{{1,0},{0,1}})
	// ScaleX *Matrix = NewMatrix([]Vector{{1,0},{0,1}})
	// Shearx *Matrix = NewMatrix([]Vector{{1,0},{0,1}})
	// ShearY *Matrix = NewMatrix([]Vector{{1,0},{0,1}})
	// ReflectXequalsY *Matrix = NewMatrix([]Vector{{1,0},{0,1}})
	// ReflectY *Matrix = NewMatrix([]Vector{{1,0},{0,1}})
	// ReflectX *Matrix = NewMatrix([]Vector{{1,0},{0,1}})

)

type Pt float64

type Vector []Pt 

// Add some methods to the vector Type 

// Sum of PointSquared then find the Squareroot of the whole 
func(v Vector) Magnitude()Pt {
	var mag Pt = 0 
	for _, point :=  range v {
		mag += point*point
	}
	return Pt(math.Sqrt(float64(mag)))
}

func (v Vector) Normalize() *Vector {
	var vec Vector = []Pt{ };
	var mag Pt = v.Magnitude(); 
	for _, point := range v {vec = append(vec, point/mag)}
	return &vec
}

// multiply vector by a scalar 
func (v Vector) Scale(scalar Pt) *Vector {
	var vec Vector = []Pt{}
	for _, point := range v { point *= scalar; vec = append(vec, point)}
	return &vec 
}

// Dimension of the Vector 
func (v Vector) Dimension() int{return len(v)}

// Sum of 2 vectors 
func(v Vector) Sum(other Vector) *Vector {
	var vec Vector = []Pt{}
	// ensure the vectors are of equal dimension 
	v, other = equalizeDim(v, other, 0)
	for i, point :=range v {vec = append(vec, point + other[i])}
	return &vec
}

// Difference of two vectors
func(v Vector) Sub(other Vector) * Vector {
	var vec Vector = []Pt{}
	v, other = equalizeDim(v,other, 0)
	for i, point := range v {vec = append(vec, point - other[i])}
	return &vec 
}
 
// find the Dot Product of two vectors 
func(v Vector) Dot(other Vector) Pt {
	// ensure the Vectors are of same Dimension
	v, other = equalizeDim(v,other,1);
	var dotP Pt = 0 
	for i, point := range v {dotP += point * other[i]}
	return dotP
}

// find the cosine similarity of 2 vectors 
func(v Vector) CosineSim(other Vector) Pt {
	return v.Dot(other) / (v.Magnitude() * other.Magnitude())
}

// find ProjectOnto
func(v Vector) ProjectOnto(other Vector) *Vector {
	scalar := v.Dot(other) / other.Dot(other) 
	return other.Scale(scalar)
}

func equalizeDim(vec1, vec2 Vector, filler Pt) (Vector, Vector) {
	pad := func(vec Vector, n int) Vector{
		for range n {vec = append(vec, filler)}
		return vec
	}
	if len(vec1) != len(vec2) {
		if len(vec1) > len(vec2) {pad(vec2, len(vec1)-len(vec2))
		}else{pad(vec1, len(vec2)-len(vec1))}
	}
	return vec1, vec2
}


type Matrix struct {
	Vals []Vector 
	Rows int 
	Cols int 
	Shape []int
}
	
func NewMatrix (mat []Vector) *Matrix {
	var rows, cols int = len(mat), len(mat[0])
	return &Matrix{
		Vals: mat,
		Rows: rows,
		Cols: cols,
		Shape: []int{rows, cols},
	}
}

// Det, and inverse of 2 by 2 
// Det - Main Diagonal product - less diagonal product - the rows must be equals to cols 
func (mat *Matrix) Determinant2X2() (Pt, error) {
	if mat.Cols != mat.Rows {return Null, MError(ErrDETMatrixNotSquare)}
	// get the main and less diagonal products 
	var main, less Pt = 1, 1;
	var i, j int = 0, mat.Cols-1
	for _, row := range mat.Vals {
		main *= row[i] 
		less *= row[j]
		i++; j--
	}
	return main-less, nil
}

// Transpose 
func (mat Matrix) Transpose() *Matrix {
	v := make([]Vector, mat.Cols)
	for i := range mat.Vals {
		for j :=0 ; j < mat.Cols; j++ {
			v[j] = append(v[j], mat.Vals[i][j]) 
		}
	}
	return NewMatrix(v)
}

// Inverse2X2 - swap main diagonal values, change sign of the other values, divide by Determinant
func (mat Matrix) Inverse2X2() (*Matrix, error) {
	if mat.Shape[0] != 2 || mat.Shape[1] != 2 {return nil, MError(ErrINVMatrixNot2by2) }
	det, _ := mat.Determinant2X2() //ignoring the error as it is already handled in Det() function
	mat.Vals[0][0],mat.Vals[1][1] = mat.Vals[1][1],mat.Vals[0][0]

	for i := range mat.Vals {
		for j:=0; j<mat.Cols; j++{
			if i != j{mat.Vals[i][j] *= -1}
			mat.Vals[i][j] /= det
		}
	}
	return NewMatrix(mat.Vals), nil
}

func IdentityMatrix(dim int) *Matrix {
	vec := make([]Vector, dim)
	for row :=range dim {
		for col := range dim {
			if row == col {vec[row] = append(vec[row], 1)
			}else {vec[row] = append(vec[row], 0)}
		}
	}
	return NewMatrix(vec)
}

// A trial to implement, the Inverse method of any square Matrix --[Gauss-Jordan Elimination]
func (m Matrix) Inverse() (*Matrix, error) {
	// Check if the Matrix is a Square 
	if m.Cols != m.Rows {return nil, MError(ErrINVMatrixNotSquare) }

	n := m.Cols // original column count 
	var vec []Vector = make([]Vector, m.Cols)
	mat := CopyMatrix(m)
	mat = *mat.JoinMatToIdentityMat()

	for col := range n {
		// find the pivot 
		pivot := col 

		for row := col + 1; row < n; row++ {
			if abs(mat.Vals[row][col]) > abs(mat.Vals[pivot][col]){
				pivot = row
			}
		}
		// put pivot Row in position 
		if pivot != col {mat = swapRows(mat, pivot, col)}

		// make pivot = 1 
		pivotValue := mat.Vals[col][col]
		mat = scaleRow(mat, col, 1/pivotValue)

		// Eliminate column - if the column is a pivot column, jump iteration 
		for row := range n {
			if row == col {continue}
			factor := -mat.Vals[row][col]
			mat = addRowMultiple(mat, row, col, factor)
		}
	}

	// extract the inverse part of the matrix 
	for i := range n {
		vec[i] = make([]Pt, n) 
		for j := range n {vec[i][j] = mat.Vals[i][n+j]}
	}
	
	return NewMatrix(vec), nil
}

// Helper functions for the Inverse Function 

// Augmenting the Matrix to Identity Matrix 
func (mat Matrix) JoinMatToIdentityMat() *Matrix {
	IdtMat := IdentityMatrix(mat.Cols)
	for i := range mat.Vals {
		mat.Vals[i] = append(mat.Vals[i], IdtMat.Vals[i]...)
	}
	return NewMatrix(mat.Vals)
}

func swapRows(m Matrix, row1,row2 int) Matrix {
	m.Vals[row1], m.Vals[row2] = m.Vals[row2], m.Vals[row1]
	return m
}

func scaleRow(m Matrix, row int, scalar Pt) Matrix{
	for i := range m.Vals[row] {
		m.Vals[row][i] *= scalar
	}
	return m
}

// for example R2 <- R2[targetRow] + Multiple.R1[targetRow] // e.g R3 <- R3 + 2.5R1
func addRowMultiple(m Matrix, targetRow, pivotRow int, multiple Pt) Matrix {
	for i, value := range m.Vals[pivotRow] {
		m.Vals[targetRow][i] +=(value* multiple)
	}
	return m 
}

func abs(val Pt) Pt {
	if val < 0 {return -1 * val}
	return val
}
//-- End of Helper functions --//

// Scale - scalar multiplication 
func (m Matrix) Scale(scalar Pt) *Matrix{
	mat := CopyMatrix(m)
	for i:= 0; i < len(mat.Vals); i++ {
		for j := range mat.Vals[i] {
			mat.Vals[i][j] *= scalar
		} 
	}
	return NewMatrix(mat.Vals)
}

// Matrix multiplication - MatMul
func (mat Matrix) MatMul(mtx Matrix) (*Matrix, error) {
	// check for compatibility of Matrix multiplication 
	if mat.Shape[1] != mtx.Shape[0] {return nil, MError(ErrMulMatrixNotCompatible)}
	// result Matrix 
	var v = make([]Vector, mat.Rows)
	// The matrix Multiplication uses 3 intertwinned for loops 
	for i := range mat.Rows {
		for j := range mtx.Cols {
			var point Pt = 0
			for k := range mtx.Rows {
				point += mat.Vals[i][k] * mtx.Vals[k][j]
			}
			v[i] = append(v[i], point)
		}
	}
	return NewMatrix(v),nil  
}


//  A func to deep copy a Matrix and return the copy 
func CopyMatrix(m Matrix) Matrix {
	var vec []Vector = make([]Vector, len(m.Vals))
	for i := range len(vec) {
		vec[i] = make([]Pt, len(m.Vals[i]))
		copy(vec[i], m.Vals[i])
	}
	return *NewMatrix(vec)
}
