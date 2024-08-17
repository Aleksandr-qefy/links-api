package set

type Set[T comparable] struct {
	set map[T]struct{}
}

func NewSet[T comparable](els ...T) Set[T] {
	set := make(map[T]struct{})
	for _, el := range els {
		set[el] = struct{}{}
	}
	return Set[T]{set: set}
}

func (s Set[T]) Add(el T) {
	s.set[el] = struct{}{}
}

func (s Set[T]) Remove(el T) {
	delete(s.set, el)
}

func (s Set[T]) Union(s2 Set[T]) Set[T] {
	unionSet := make(map[T]struct{})

	for el, _ := range s.set {
		unionSet[el] = struct{}{}
	}

	for el, _ := range s2.set {
		unionSet[el] = struct{}{}
	}

	return Set[T]{set: unionSet}
}

func (s Set[T]) Minus(s2 Set[T]) Set[T] {
	minusSet := make(map[T]struct{})

	for el, _ := range s.set {
		if _, ok := s2.set[el]; !ok {
			minusSet[el] = struct{}{}
		}
	}

	return Set[T]{set: minusSet}
}

func (s Set[T]) Intersection(s2 Set[T]) Set[T] {
	intersectionSet := make(map[T]struct{}, len(s.set))

	for el, _ := range s.set {
		if _, ok := s2.set[el]; ok {
			intersectionSet[el] = struct{}{}
		}
	}

	return Set[T]{set: intersectionSet}
}

func (s Set[T]) Slice() []T {
	slc := make([]T, len(s.set))

	k := 0
	for el, _ := range s.set {
		slc[k] = el
		k++
	}

	return slc
}
