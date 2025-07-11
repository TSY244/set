package set

type Set[T comparable] struct {
	m map[T]struct{} // unsafe.Sizeof(struct{}{}) == 0
}

func NewSet[T comparable]() *Set[T] {
	s := &Set[T]{}
	s.m = make(map[T]struct{})
	return s
}

func (s *Set[T]) Add(value T) {
	s.m[value] = struct{}{}
}

func (s *Set[T]) AddAll(value []T) {
	for _, v := range value {
		s.m[v] = struct{}{}
	}
}

func (s *Set[T]) Remove(value T) {
	delete(s.m, value)
}

func (s *Set[T]) Contains(value T) bool {
	_, ok := s.m[value]
	return ok
}

func (s *Set[T]) Size() int {
	return len(s.m)
}

func (s *Set[T]) Clear() {
	s.m = make(map[T]struct{})
}

func (s *Set[T]) Values() []T {
	values := make([]T, 0, len(s.m))
	for k, _ := range s.m {
		values = append(values, k)
	}
	return values
}
