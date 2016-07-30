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

func (s IntSet) Len() int {
	return len(s.elements)
}

func orderSetsByLength(s1 IntSet, s2 IntSet) (IntSet, IntSet) {
	if s1.Len() < s2.Len() {
		return s1, s2
	}
	return s2, s1
}

func (s1 IntSet) Union(s2 IntSet) IntSet {
	union := func(small IntSet, large IntSet) IntSet {
		for x, _ := range small.elements {
			large.Add(x)
		}
		return large
	}

	ss1, ss2 := orderSetsByLength(s1, s2)
	return union(ss1, ss2)
}

func (s IntSet) Member(xs ...int) bool {
	if s.Len() == 0 {
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
	if s1.Len() != s2.Len() {
		return false
	}

	equal := func(small IntSet, large IntSet) bool {
		for x, _ := range small.elements {
			if found := large.Member(x); !found {
				return false
			}
		}
		return true
	}

	ss1, ss2 := orderSetsByLength(s1, s2)
	return equal(ss1, ss2)
}

func (s1 IntSet) Intersect(s2 IntSet) IntSet {
	intersection := NewIntSet()
	ss1, ss2 := orderSetsByLength(s1, s2)

	for x := range ss1.elements {
		if ss2.Member(x) {
			intersection = intersection.Add(x)
		}
	}

	return intersection
}
