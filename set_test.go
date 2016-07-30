package intset

import (
	"testing"
)

func TestAdd(t *testing.T) {
	s := NewIntSet().Add(4, 4, 5).Add(4).Delete(4, 4, 4, 3).Delete().Add()
	if s.Len() != 1 {
		t.Errorf("expected 1 element got %d", s.Len())
	}
}

func TestUnion(t *testing.T) {
	s := NewIntSet(1, 2, 3, 4).Union(NewIntSet(2, 3, 4, 5))
	if s.Len() != 5 {
		t.Errorf("expected 5 element got %d", s.Len())
	}
}

func TestMember(t *testing.T) {
	s := NewIntSet(1, 2, 3, 4, 5)
	if !s.Member(1, 2, 3, 4) {
		t.Error("missing member")
	}
}

func TestEqual(t *testing.T) {
	s1 := NewIntSet(1, 2, 3, 4)
	s2 := NewIntSet(1, 2, 3, 4)
	if !s1.Equal(s2) {
		t.Error("not equal")
	}
}
