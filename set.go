package intset

type IntSet struct {
	elements map[int]struct{}
}

func NewIntSet(xs ...int) IntSet {
	s := IntSet{map[int]struct{}{}}

	return s.Add(xs...)
}

func (s IntSet) Add(xs ...int) IntSet {
	for _, x := range xs {
		s.elements[x] = struct{}{}
	}
	return s
}

func (s IntSet) Delete(xs ...int) IntSet {
	for _, x := range xs {
		if _, found := s.elements[x]; found {
			delete(s.elements, x)
		}
	}
	return s
}

func (s IntSet) Cardinality() int {
	return len(s.elements)
}

func orderSetsByLength(s1 IntSet, s2 IntSet) (IntSet, IntSet) {
	if s1.Cardinality() < s2.Cardinality() {
		return s1, s2
	}
	return s2, s1
}

func (s1 IntSet) Union(s2 IntSet) IntSet {
	union := s1.Copy()

	for x, _ := range s2.elements {
		union.Add(x)
	}

	return union
}

func (s IntSet) Member(xs ...int) bool {
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

func (s1 IntSet) Equal(s2 IntSet) bool {
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

func (s1 IntSet) Intersect(s2 IntSet) IntSet {
	intersection := NewIntSet()
	small, large := orderSetsByLength(s1, s2)

	for x := range small.elements {
		if large.Member(x) {
			intersection = intersection.Add(x)
		}
	}

	return intersection
}

func (s IntSet) Copy() IntSet {
	copy := NewIntSet()
	for x := range s.elements {
		copy.Add(x)
	}
	return copy
}

func (s1 IntSet) Difference(s2 IntSet) IntSet {
	difference := s1.Copy()
	for x := range s2.elements {
		difference.Delete(x)
	}
	return difference
}

func (s1 IntSet) SymmetricDifference(s2 IntSet) IntSet {
	return s1.Difference(s2).Union(s2.Difference(s1))
}
