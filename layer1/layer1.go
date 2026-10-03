package stat

// a type to hold the statistical data values
type Stat[T int | float64] struct {
	data []T
	count int 

}

func NewStat[T int | float64] () *Stat[T] {
	return &Stat[T]{data: []T{}} 
}

// Count - return the total number of elements in s.data
func (s Stat[T]) Count() int {return s.count}

// Push - Populates the s.data with the values v 
func (s *Stat[T]) Push(v ...T) {
	s.data = append(s.data, v...)
	s.count += len(v)
} 

// Reverse - a function to reverse the order of and array via addition/subtraction 
func (s Stat[T]) Reverse()[]T{
	for i, j := 0, s.count - 1; i <= j; i++{
		// the difference between two values that are to be exahanged
		var factor = s.data[i] - s.data[j]
		s.data[i], s.data[j] = s.data[i] - T(factor), s.data[j] + T(factor)
		j--
	}
	return s.data
}

// GetData - return the array of Data in it's latest state
func (s Stat[T]) GetData() []T {
	return s.data
}

// Sum - calculates the sum of all the elements in s.data and return it
func (s Stat[T]) Sum() T {
	sum := T(0)
	for i, j := 0, s.count - 1; i<= j; i++ {
		if i==j{
			sum += s.data[i]
		}else {sum += (s.data[i] + s.data[j])}
		j--
	}
	return sum 
}

// Mean - return the mean/average of the data i.e. SUM(s.data)/coun
func (s *Stat[T]) Mean() T {
	return s.Sum() / T(s.count)
}

// Sort - sort the elements of the s.data in ascending order 
// Done with a minHeap Sort algorithm 
func (s *Stat[T]) HeapSort () []T { 
	ordered, count := s.data, len(s.data)
	for {
		if count == 0{ break
		}else{
			newArr := Pluck(ordered, count)
			ordered = newArr
			count -- 
		}
	}
	return ordered
}

// Sorting Helper functions 
// siftDown() - maintains the minHeap properties via sifting down from index idx -- Top to Bottom 
func siftDown[T int | float64] (arr []T, idx int) {
	count, current := len(arr), idx
	if count < 2  {return}
	for { 
		lt, rt := left(current), right(current)
		// choose between right and left node for exchange 
		if lt >= count {return}  
		if rt >= count || arr[lt] <= arr[rt] {
			// exchange for left 
			if arr[current] > arr[lt] { arr[current], arr[lt] = arr[lt], arr[current]}
			current = lt
		}else {
			// exchange for right
			if arr[current] > arr[rt] {arr[current], arr[rt] = arr[rt], arr[current]}
			current = rt
		}
	}
}

// SiftUp() - help maintain the minHeap properties via sifting up from index idx in the slice arr -- Bottom to Top
func SiftUp[T int | float64] (arr[]T, idx int ) {
	if count := len(arr); count < 2 {
		return
	}
	for{ 
		par := parent(idx)
		if arr[idx] < arr[par] {
			arr[idx], arr[par] = arr[par], arr[idx]
			idx = par
		}else{break}
	}	
}

// Pluck() - removes the top element in the minHeap while maintaining the Heap Property
func Pluck[T int | float64] (arr []T, max int) ([]T) {
	for i := int(max/2); i >= 0; i-- {
		siftDown(arr[:max-1], i)
	}
	minEl, trimmed := arr[0], arr[1:]
	arr = append(trimmed, minEl)
	return arr
}

func left (idx int) int{
	if idx < 0 {return -1}
	return (idx << 1) + 1
}
func right (idx int) int{
	if idx < 0 {return -1}
	return (idx << 1) + 2
}
func parent (idx int) int{
	if idx >= 0 && idx < 3 {return 0}
	if idx < 0 {return -1}
	return int((idx - 1) / 2) 
}

// Abs - returns the absolute value of num
func Abs(num int) int {
	if num < 0 {return num * -1}
	return num 
}