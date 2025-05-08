package stack

type Stack[Item any] struct {
	arr []Item
}

func NewStack[Item any](cap int) Stack[Item] {
	return Stack[Item]{
		arr: make([]Item, 0, cap),
	}
}

func (s *Stack[Item]) Push(item Item) {
	s.arr = append(s.arr, item)
}

func (s *Stack[Item]) Pop() Item {
	val := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return val
}

func (s *Stack[Item]) Top() Item {
	return s.arr[len(s.arr)-1]
}

func (s *Stack[Item]) Size() int {
	return len(s.arr)
}

func (s *Stack[Item]) IsEmpty() bool {
	return len(s.arr) == 0
}
