package sets

// Set represents a set of elements of type T
type Set[T comparable] map[T]struct{}

// Make
// create a new Set of type T
func Make[T comparable](values ...T) Set[T] {
	s := make(Set[T])
	s.Add(values...)
	return s
}

// Add
// adds values to the Set
func (s Set[T]) Add(values ...T) {
	for _, v := range values {
		s[v] = struct{}{}
	}
}

// Delete
// deletes values from the Set
func (s Set[T]) Delete(values ...T) {
	for _, v := range values {
		delete(s, v)
	}
}

// Truncate
// deletes all values from the Set
func (s *Set[T]) Truncate() {
	*s = Make[T]()
}

// Has
// returns true if Set contains the value or false if not
func (s Set[T]) Has(value T) bool {
	_, e := s[value]
	return e
}

// Len
// returns the length of the Set
func (s Set[T]) Len() int {
	return len(s)
}

// Values
// returns the Set values
func (s Set[T]) Values() []T {
	values := make([]T, 0, len(s))
	for v := range s {
		values = append(values, v)
	}
	return values
}

// Merge
// adds values of the other Sets
func (s Set[T]) Merge(others ...Set[T]) {
	for i := range others {
		s.Add(others[i].Values()...)
	}
}

// Diff
// removes all values represented in any of the other Sets
func (s Set[T]) Diff(others ...Set[T]) {
	if len(others) == 0 {
		return
	}

	others[0].Merge(others[1:]...)

	for _, v := range s.Values() {
		if others[0].Has(v) {
			s.Delete(v)
		}
	}
}

// Intersect
// removes all values not represented in all others Sets
func (s *Set[T]) Intersect(others ...Set[T]) {
	if len(others) == 0 {
		s.Truncate()
		return
	}

	for _, v := range s.Values() {
		for i := range others {
			if !others[i].Has(v) {
				s.Delete(v)
			}
		}
	}
}

// Equals
// returns true if the Sets are equal to each other
func (s Set[T]) Equals(other Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}

	for _, v := range s.Values() {
		if !other.Has(v) {
			return false
		}
	}

	return true
}

// Filter
// removes elements for which the func returns false
func (s Set[T]) Filter(f func(T) bool) {
	for _, v := range s.Values() {
		if !f(v) {
			s.Delete(v)
		}
	}
}

// Map
// calls the func for each element
func (s *Set[T]) Map(f func(T) T) {
	values := make([]T, 0, s.Len())
	for _, v := range s.Values() {
		values = append(values, f(v))
	}
	s.Truncate()
	s.Add(values...)
}

// Copy
// returns a copy of the Set
func (s Set[T]) Copy() Set[T] {
	newSet := Make[T]()
	newSet.Add(s.Values()...)
	return newSet
}
