package median

import (
	"testing"
)

var sink float64

func BenchmarkMedian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = Median([]int{1, 2, 3, 5, 6, 7, 8, 9, 10}, []int{4, 11})
	}
}

func BenchmarkMedian2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = Median2([]int{1, 2, 3, 5, 6, 7, 8, 9, 10}, []int{4, 11})
	}
}

func BenchmarkMedian3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = Median3([]int{1, 2, 3, 5, 6, 7, 8, 9, 10}, []int{4, 11})
	}
}
