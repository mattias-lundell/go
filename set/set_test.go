package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		actual      Set
		expected    Set
		description string
	}{

		{
			NewSet().Add(4, 4),
			NewSet(4),
			"should handle adding dublicate",
		},
		{
			NewSet().Add(4, 4, 5).Add(6),
			NewSet(4, 5, 6),
			"should handle chaining multiple calls",
		},
		{
			NewSet().Add(),
			NewSet(),
			"should handle empty",
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.actual, tt.description)
	}
}

func TestDelete(t *testing.T) {
	var tests = []struct {
		actual       Set
		expected     Set
		desctription string
	}{
		{
			NewSet(4, 5, 6, 7).Delete(4),
			NewSet(5, 6, 7),
			"delete 4 from a set containging 4 should return a set without 4",
		},
		{
			NewSet(4, 5, 6, 7).Delete(),
			NewSet(4, 5, 6, 7),
			"should handle empty delete",
		},
		{
			NewSet(4, 5, 6, 7).Delete(3, 4, 5, 6, 7, 8, 9),
			NewSet(),
			"should handle multiple delete",
		},
		{
			NewSet(4, 5, 6, 7).Delete(4).Delete(5),
			NewSet(6, 7),
			"should handle chained deletes",
		},
		{
			NewSet(4, 5, 6, 7).Delete(1),
			NewSet(4, 5, 6, 7),
			"should handle delete of element not in set",
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.actual, tt.desctription)
	}
}

func TestUnion(t *testing.T) {
	var tests = []struct {
		actual   Set
		expected Set
	}{
		{
			NewSet(4, 5, 6, 7).Union(NewSet(4, 5, 6, 7)),
			NewSet(4, 5, 6, 7),
		},
		{
			NewSet(4, 5, 6, 7).Union(NewSet(8, 9)),
			NewSet(4, 5, 6, 7, 8, 9),
		},
		{
			NewSet(4, 5, 6, 7).Union(NewSet()),
			NewSet(4, 5, 6, 7),
		},
		{
			NewSet().Union(NewSet(4, 5, 6, 7)),
			NewSet(4, 5, 6, 7),
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
			NewSet(4, 5, 6, 7).Member(4),
			true,
		},
		{
			NewSet(4, 5, 6, 7).Member(1),
			false,
		},
		{
			NewSet(4, 5, 6, 7).Member(4, 5, 6, 7),
			true,
		},
		{
			NewSet(4, 5, 6, 7).Member(4, 5, 6, 7, 8),
			false,
		},
		{
			NewSet().Member(4),
			false,
		},
		{
			NewSet(3).Member(4),
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.actual)
	}
}

func TestEqual(t *testing.T) {
	var tests = []struct {
		s1       Set
		s2       Set
		expected bool
	}{
		{
			NewSet(4, 5, 6, 7),
			NewSet(4, 5, 6, 7),
			true,
		},
		{
			NewSet(4, 5, 6),
			NewSet(4, 5, 6, 7),
			false,
		},
		{
			NewSet(),
			NewSet(4, 5, 6, 7),
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.s1.Equal(tt.s2))
	}
}

func TestIntersect(t *testing.T) {
	var tests = []struct {
		s1       Set
		s2       Set
		expected Set
	}{
		{
			NewSet(4, 5, 6, 7),
			NewSet(4, 5, 6, 7),
			NewSet(4, 5, 6, 7),
		},
		{
			NewSet(4, 5, 6),
			NewSet(4, 5, 6, 7),
			NewSet(4, 5, 6),
		},
		{
			NewSet(),
			NewSet(4, 5, 6, 7),
			NewSet(),
		},
		{
			NewSet(1),
			NewSet(2),
			NewSet(),
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.s1.Intersect(tt.s2))
	}
}

func TestDifference(t *testing.T) {
	var tests = []struct {
		s1       Set
		s2       Set
		expected Set
	}{
		{
			NewSet(4, 5, 6, 7),
			NewSet(4, 5, 6, 7),
			NewSet(),
		},
		{
			NewSet(4, 5, 6),
			NewSet(4, 5, 6, 7),
			NewSet(),
		},
		{
			NewSet(1, 2, 3),
			NewSet(3, 4, 5),
			NewSet(1, 2),
		},
		{
			NewSet(1),
			NewSet(2),
			NewSet(1),
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.s1.Difference(tt.s2))
	}
}

func TestSymmetricDifference(t *testing.T) {
	var tests = []struct {
		s1       Set
		s2       Set
		expected Set
	}{
		{
			NewSet(4, 5, 6, 7),
			NewSet(4, 5, 6, 7),
			NewSet(),
		},
		{
			NewSet(4, 5, 6),
			NewSet(4, 5, 6, 7),
			NewSet(7),
		},
		{
			NewSet(1, 2, 3),
			NewSet(3, 4, 5),
			NewSet(1, 2, 4, 5),
		},
		{
			NewSet(1),
			NewSet(2),
			NewSet(1, 2),
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.s1.SymmetricDifference(tt.s2))
	}
}
