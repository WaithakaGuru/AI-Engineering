package matrix

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

type Pt float32

type Vector []Pt 

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
func (m *Matrix) Determinant() (Pt, error) {
	if m.Cols != m.Rows {return Null, MError(ErrDETMatrixNotSquare)}
	// get the main and less diagonal products 
	var main, less Pt = 1, 1;
	var i, j int = 0, m.Cols-1
	for _, row := range m.Vals {
		main *= row[i] 
		less *= row[j]
		i++; j--
	}
	return main-less, nil
}

// Transpose 
func (m Matrix) Transpose() *Matrix {
	v := make([]Vector, m.Cols)
	for i := range m.Vals {
		for j :=0 ; j < m.Cols; j++ {
			v[j] = append(v[j], m.Vals[i][j]) 
		}
	}
	return NewMatrix(v)
}

// Inverse2X2 - swap main diagonal values, change sign of the other values, divide by Determinant
func (m Matrix) Inverse2X2() (*Matrix, error) {
	if m.Shape[0] != 2 || m.Shape[1] != 2 {return nil, MError(ErrINVMatrixNot2by2) }
	det, _ := m.Determinant() //ignoring the error as it is already handled in Det() function
	m.Vals[0][0],m.Vals[1][1] = m.Vals[1][1],m.Vals[0][0]

	for i := range m.Vals {
		for j:=0; j<m.Cols; j++{
			if i != j{m.Vals[i][j] *= -1}
			m.Vals[i][j] /= det
		}
	}
	return NewMatrix(m.Vals), nil
}

// Scale - scalar multiplication 
func (m Matrix) Scale(scalar Pt) *Matrix{
	for i := range m.Vals {
		for j := range m.Vals[i] {
			m.Vals[i][j] *= scalar
		} 
	}
	return NewMatrix(m.Vals)
}

// Matrix multiplication - MatMul
func (m Matrix) MatMul(mtx Matrix) (*Matrix, error) {
	// check for compatibility of Matrix multiplication 
	if m.Shape[1] != mtx.Shape[0] {return nil, MError(ErrMulMatrixNotCompatible)}
	// result Matrix 
	var v = make([]Vector, m.Rows)
	// The matrix Multiplication uses 3 intertwinned for loops 
	for i := range m.Rows {
		for j := range mtx.Cols {
			var point Pt = 0
			for k := range mtx.Rows {
				point += m.Vals[i][k] * mtx.Vals[k][j]
			}
			v[i] = append(v[i], point)
		}
	}
	return NewMatrix(v),nil  
}
