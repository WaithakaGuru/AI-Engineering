package stat

import "iter"

// First am going to recap on Generics in Go and iterators
// iterators -> take the function shape of func(yield func(val R) bool)

// this iterator yields even number from 0 to 100
func EvenNumsTo100(yield func(int) bool){
	for i := range 101 {
		if i % 2 == 0 {
		  if !yield(i) {return}
		}
	}
}

// Generic functions - means functions with an open and loose shape so 
	// to allow reuse by parameters of different types 
type Collection[N ~int, R any] struct {
	array []N
}

func NewCollection[N ~int, R any](List []N)Collection[N, R] {
	return Collection[N, R]{array: List}
}

// a Collection method allow one to Use a func on each of the array items -> can be a reducer or map
// Map
func(c Collection[N, R]) Use(F func(N)R) []R {
	result := make([]R, len(c.array))
	for i, item := range c.array {
		result[i] = F(item)
	}
	return result 
}

// Apply multiple maps to a value - e.g Add2 then Multiply by 2.45 and check if it is Even
// func (c Collection[N, R]) UseAll(F []func(N)N)[]R {
// 	result := []R{}
// 	for _, item := range c.array {
// 		val := item
// 		for _, mapFunc := range F {val = mapFunc(val)}
// 		result = append(result, R(val))
// 	}
// }


// Reducer - draws a single value from the whole slice e.g. adding the values 
func(c Collection[N, R]) Deduce(F func([]N)R)R{
	return F(c.array)
}

// func to apply multiple reducers - return an array of results from each reducer 
func(c Collection[N, R]) DeduceAll(F []func([]N)R) []R {
	results := make([]R, len(F))
	for _, reducer := range F {
		results = append(results, reducer(c.array))
	}
	return results
}

// Some iterators from the Collection
// Get all Even numbers
func(c Collection[N, R]) Evens()func(yield func(N)bool){
	return func(yield func(N) bool) {
		for _, item := range c.array {
			if item % 2 == 0 {if !yield(item) { return}}
		}
	}
}

// Get all odd numbers - replace the iterator shape ith predefined builtin iter.seq
func (c Collection[N, R]) Odds() iter.Seq[N] {
	return func(yield func(N) bool) {
		for _, item := range c.array {
			if item % 2 == 1 {if !yield(item) {return}}
		}
	}
}