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

func (xs IntSet) Union(ys IntSet) IntSet {
	union := func(small IntSet, large IntSet) IntSet {
		for x, _ := range small.elements {
			large.Add(x)
		}
		return large
	}

	if xs.Len() < ys.Len() {
		return union(xs, ys)
	}
	return union(ys, xs)
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

	if s1.Len() < s2.Len() {
		return equal(s1, s2)
	}
	return equal(s2, s1)
}
