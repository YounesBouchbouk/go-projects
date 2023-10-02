package stack

type Stack []int

// isEmpty check id stack is Empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

//Pop() to remove the last element and return if

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.

	return element, true
}

// Push a new value onto the stack
func (s *Stack) Push(str int) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}
