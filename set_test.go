package set_test

import (
	"math/rand"
	"testing"

	"github.com/mattias-lundell/go-set/intset"
	"github.com/mattias-lundell/go-set/set"
)

func benchmarkIntSet(b *testing.B, size int, fill int) {
	for i := 0; i < b.N; i++ {
		set := intset.NewIntSet()
		for j := 0; j < size; j++ {
			set.Add(rand.Int() % fill)
		}
	}
}

func BenchmarkSmallIntSetWithFewCollisions(b *testing.B) {
	benchmarkIntSet(b, 100, 700)
}
func BenchmarkSmallIntSetWithMoreCollisions(b *testing.B) {
	benchmarkIntSet(b, 100, 100)
}
func BenchmarkSmallIntSetWithManyCollisions(b *testing.B) {
	benchmarkIntSet(b, 100, 25)
}
func BenchmarkMediumIntSetWithFewCollisions(b *testing.B) {
	benchmarkIntSet(b, 5000, 35000)
}
func BenchmarkMediumIntSetWithMoreCollisions(b *testing.B) {
	benchmarkIntSet(b, 5000, 5000)
}
func BenchmarkMediumIntSetWithManyCollisions(b *testing.B) {
	benchmarkIntSet(b, 5000, 1250)
}
func BenchmarkLargeIntSetWithFewCollisions(b *testing.B) {
	benchmarkIntSet(b, 100000, 700000)
}
func BenchmarkLargeIntSetWithMoreCollisions(b *testing.B) {
	benchmarkIntSet(b, 100000, 100000)
}
func BenchmarkLargeIntSetWithManyCollisions(b *testing.B) {
	benchmarkIntSet(b, 100000, 25000)
}

func benchmarkSet(b *testing.B, size int, fill int) {
	for i := 0; i < b.N; i++ {
		set := set.NewSet()
		for j := 0; j < size; j++ {
			set.Add(rand.Int() % fill)
		}
	}
}

func BenchmarkSmallSetWithFewCollisions(b *testing.B) {
	benchmarkSet(b, 100, 700)
}
func BenchmarkSmallSetWithMoreCollisions(b *testing.B) {
	benchmarkSet(b, 100, 100)
}
func BenchmarkSmallSetWithManyCollisions(b *testing.B) {
	benchmarkSet(b, 100, 25)
}
func BenchmarkMediumSetWithFewCollisions(b *testing.B) {
	benchmarkSet(b, 5000, 35000)
}
func BenchmarkMediumSetWithMoreCollisions(b *testing.B) {
	benchmarkSet(b, 5000, 5000)
}
func BenchmarkMediumSetWithManyCollisions(b *testing.B) {
	benchmarkSet(b, 5000, 1250)
}
func BenchmarkLargeSetWithFewCollisions(b *testing.B) {
	benchmarkSet(b, 100000, 700000)
}
func BenchmarkLargeSetWithMoreCollisions(b *testing.B) {
	benchmarkSet(b, 100000, 100000)
}
func BenchmarkLargeSetWithManyCollisions(b *testing.B) {
	benchmarkSet(b, 100000, 25000)
}
