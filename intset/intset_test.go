package intset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		actual      IntSet
		expected    IntSet
		description string
	}{

		{
			NewIntSet().Add(4, 4),
			NewIntSet(4),
			"should handle adding dublicate",
		},
		{
			NewIntSet().Add(4, 4, 5).Add(6),
			NewIntSet(4, 5, 6),
			"should handle chaining multiple calls",
		},
		{
			NewIntSet().Add(),
			NewIntSet(),
			"should handle empty",
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.actual, tt.description)
	}
}

func TestDelete(t *testing.T) {
	var tests = []struct {
		actual       IntSet
		expected     IntSet
		desctription string
	}{
		{
			NewIntSet(4, 5, 6, 7).Delete(4),
			NewIntSet(5, 6, 7),
			"delete 4 from a set containging 4 should return a set without 4",
		},
		{
			NewIntSet(4, 5, 6, 7).Delete(),
			NewIntSet(4, 5, 6, 7),
			"should handle empty delete",
		},
		{
			NewIntSet(4, 5, 6, 7).Delete(3, 4, 5, 6, 7, 8, 9),
			NewIntSet(),
			"should handle multiple delete",
		},
		{
			NewIntSet(4, 5, 6, 7).Delete(4).Delete(5),
			NewIntSet(6, 7),
			"should handle chained deletes",
		},
		{
			NewIntSet(4, 5, 6, 7).Delete(1),
			NewIntSet(4, 5, 6, 7),
			"should handle delete of element not in set",
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.actual, tt.desctription)
	}
}

func TestUnion(t *testing.T) {
	var tests = []struct {
		actual   IntSet
		expected IntSet
	}{
		{
			NewIntSet(4, 5, 6, 7).Union(NewIntSet(4, 5, 6, 7)),
			NewIntSet(4, 5, 6, 7),
		},
		{
			NewIntSet(4, 5, 6, 7).Union(NewIntSet(8, 9)),
			NewIntSet(4, 5, 6, 7, 8, 9),
		},
		{
			NewIntSet(4, 5, 6, 7).Union(NewIntSet()),
			NewIntSet(4, 5, 6, 7),
		},
		{
			NewIntSet().Union(NewIntSet(4, 5, 6, 7)),
			NewIntSet(4, 5, 6, 7),
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.actual)
	}
}

func TestMember(t *testing.T) {
	var tests = []struct {
		actual   bool
		expected bool
	}{
		{
			NewIntSet(4, 5, 6, 7).Member(4),
			true,
		},
		{
			NewIntSet(4, 5, 6, 7).Member(1),
			false,
		},
		{
			NewIntSet(4, 5, 6, 7).Member(4, 5, 6, 7),
			true,
		},
		{
			NewIntSet(4, 5, 6, 7).Member(4, 5, 6, 7, 8),
			false,
		},
		{
			NewIntSet().Member(4),
			false,
		},
		{
			NewIntSet(3).Member(4),
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.actual)
	}
}

func TestEqual(t *testing.T) {
	var tests = []struct {
		s1       IntSet
		s2       IntSet
		expected bool
	}{
		{
			NewIntSet(4, 5, 6, 7),
			NewIntSet(4, 5, 6, 7),
			true,
		},
		{
			NewIntSet(4, 5, 6),
			NewIntSet(4, 5, 6, 7),
			false,
		},
		{
			NewIntSet(),
			NewIntSet(4, 5, 6, 7),
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.s1.Equal(tt.s2))
	}
}

func TestIntersect(t *testing.T) {
	var tests = []struct {
		s1       IntSet
		s2       IntSet
		expected IntSet
	}{
		{
			NewIntSet(4, 5, 6, 7),
			NewIntSet(4, 5, 6, 7),
			NewIntSet(4, 5, 6, 7),
		},
		{
			NewIntSet(4, 5, 6),
			NewIntSet(4, 5, 6, 7),
			NewIntSet(4, 5, 6),
		},
		{
			NewIntSet(),
			NewIntSet(4, 5, 6, 7),
			NewIntSet(),
		},
		{
			NewIntSet(1),
			NewIntSet(2),
			NewIntSet(),
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.s1.Intersect(tt.s2))
	}
}

func TestDifference(t *testing.T) {
	var tests = []struct {
		s1       IntSet
		s2       IntSet
		expected IntSet
	}{
		{
			NewIntSet(4, 5, 6, 7),
			NewIntSet(4, 5, 6, 7),
			NewIntSet(),
		},
		{
			NewIntSet(4, 5, 6),
			NewIntSet(4, 5, 6, 7),
			NewIntSet(),
		},
		{
			NewIntSet(1, 2, 3),
			NewIntSet(3, 4, 5),
			NewIntSet(1, 2),
		},
		{
			NewIntSet(1),
			NewIntSet(2),
			NewIntSet(1),
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.s1.Difference(tt.s2))
	}
}

func TestSymmetricDifference(t *testing.T) {
	var tests = []struct {
		s1       IntSet
		s2       IntSet
		expected IntSet
	}{
		{
			NewIntSet(4, 5, 6, 7),
			NewIntSet(4, 5, 6, 7),
			NewIntSet(),
		},
		{
			NewIntSet(4, 5, 6),
			NewIntSet(4, 5, 6, 7),
			NewIntSet(7),
		},
		{
			NewIntSet(1, 2, 3),
			NewIntSet(3, 4, 5),
			NewIntSet(1, 2, 4, 5),
		},
		{
			NewIntSet(1),
			NewIntSet(2),
			NewIntSet(1, 2),
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.s1.SymmetricDifference(tt.s2))
	}
}
