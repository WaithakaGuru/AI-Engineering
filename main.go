package main

import (
	"fmt"

	stat "github.com/WaithakaGuru/go-algorithms/Statistics/layer1"
	matrix "github.com/WaithakaGuru/go-algorithms/Statistics/layer2"
)

func main() {
	var s = stat.NewStat[int]()
	s.Push(23, 21, 20, 45, 3, 21, 34, 55, 43, 61, 11, 14, 18, 0, 9, 7, 37, 83, 77, 42, 90)
	// fmt.Println("Original Data state: ", s.GetData())
	// fmt.Println("The reverse Data Array/Slice: ", s.Reverse())
	// fmt.Println("A check on whether the Reverse() method changes the internal data order: ", s.GetData())
	// fmt.Println("The Sorted Data -- Ascending Order--: ", s.HeapSort())
	// fmt.Println("Mean of the Data: ", s.Mean())
	// fmt.Println("Sum of the Data: ", s.Sum())
	// fmt.Println(s.Count())
	// fmt.Println( 714 / s.Count())
	
	// for i := range stat.EvenNumsTo100{
	// 	println(i)
	// }

	// // Testing the collection Iterators e.g. Evens
	// items := []int{1,2,3,4,5,6,7,8,9,0,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24}
	// collection1 := stat.NewCollection[int, any](items)
	
	// // print all even numbers in the Collection
	// for even := range collection1.Evens(){
	// 	fmt.Println(even)
	// }
	// // Odds
	// for odd := range collection1.Odds(){
	// 	fmt.Println(odd)
	// }


	matrix1 := matrix.NewMatrix([]matrix.Vector{{1,3},{2,4}})
	// fmt.Println(matrix1.Transpose().Vals)
	// fmt.Println(matrix1.Inverse2X2())

	// fmt.Println(*matrix1.Scale(5))
	// fmt.Println(matrix1.Determinant())

	// Testing the Matrix Multiplication with unitMatrix - thus the answer should just be the same 
	fmt.Println("Matrix1 before any Multiplication: ")
	fmt.Println(matrix1)
	fmt.Println("Multiplying with UnitMatrix")
	fmt.Println(matrix1.MatMul(*matrix.UnitMatrix)) 
	// Answer Got for the test above is -> [[1 3] [2 4]]  - same matrix after MatMul with UnitMat -> Test above PASSED

	mat2 := matrix.NewMatrix([]matrix.Vector{{2,4}})
	// fmt.Println(matrix1.Shape,mat2.Shape) // Mat2 is incompatible for multplication as it has the shape [1 2]
	// transpose mat2 to make it compatible 
	mat2 = mat2.Transpose()
	fmt.Println("Multiplying with Mat2 of shape[2 1] {[2],[4]}")
	fmt.Println(matrix1.MatMul(*mat2))

	// Another test with a 2X3 matrix 
	fmt.Println("Multiplying with mat3 of shape[2 3] {[1,2,3], [4,5,6.455]}")	
	mat3 := matrix.NewMatrix([]matrix.Vector{{1,2,3},{4,5,6.455}})
	fmt.Println(matrix1.MatMul(*mat3))

	fmt.Println(matrix.IdentityMatrix(4))  

	// testing the trigonometry functions 
	fmt.Println(matrix.Sin(60))
	fmt.Println(matrix.ConvertDegToRadian(60))
	fmt.Println(matrix.ConvertRadToDegree(1.0471976))

	// Testing the JoinTOIdentityMatrix method 
	mat4 := matrix.NewMatrix([]matrix.Vector{{1,2,3,4}, {5,6,7,8}, {9,1,2,3},{0,4,5,6}})
	mat4 = mat4.JoinMatToIdentityMat()
	fmt.Println(mat4)
}