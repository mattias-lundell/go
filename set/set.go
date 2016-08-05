package set

type Set struct {
	elements map[interface{}]struct{}
}

func NewSet(xs ...interface{}) Set {
	s := Set{map[interface{}]struct{}{}}

	return s.Add(xs...)
}

func (s Set) Add(xs ...interface{}) Set {
	for _, x := range xs {
		s.elements[x] = struct{}{}
	}
	return s
}

func (s Set) Delete(xs ...interface{}) Set {
	for _, x := range xs {
		if _, found := s.elements[x]; found {
			delete(s.elements, x)
		}
	}
	return s
}

func (s Set) Cardinality() int {
	return len(s.elements)
}

func orderSetsByLength(s1 Set, s2 Set) (Set, Set) {
	if s1.Cardinality() < s2.Cardinality() {
		return s1, s2
	}
	return s2, s1
}

func (s1 Set) Union(s2 Set) Set {
	union := s1.Copy()

	for x, _ := range s2.elements {
		union.Add(x)
	}

	return union
}

func (s Set) Member(xs ...interface{}) bool {
	if s.Cardinality() == 0 {
		return false
	}

	for _, x := range xs {
		if _, found := s.elements[x]; !found {
			return false
		}
	}
	return true
}

func (s1 Set) Equal(s2 Set) bool {
	if s1.Cardinality() != s2.Cardinality() {
		return false
	}

	for x, _ := range s1.elements {
		if found := s2.Member(x); !found {
			return false
		}
	}
	return true
}

func (s1 Set) Intersect(s2 Set) Set {
	intersection := NewSet()
	small, large := orderSetsByLength(s1, s2)

	for x := range small.elements {
		if large.Member(x) {
			intersection = intersection.Add(x)
		}
	}

	return intersection
}

func (s Set) Copy() Set {
	copy := NewSet()
	for x := range s.elements {
		copy.Add(x)
	}
	return copy
}

func (s1 Set) Difference(s2 Set) Set {
	difference := s1.Copy()
	for x := range s2.elements {
		difference.Delete(x)
	}
	return difference
}

func (s1 Set) SymmetricDifference(s2 Set) Set {
	return s1.Difference(s2).Union(s2.Difference(s1))
}
