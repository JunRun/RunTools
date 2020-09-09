package syncpool

import "testing"

const N = 1024 * 1024

type Element int64

var xForMake = make([]Element, N)
var xForMakeCopy = make([]Element, N)
var xForAppend = make([]Element, N)
var yForMake []Element
var yForMakeCopy []Element
var yForAppend []Element

func Benchmark_PureMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		yForMake = make([]Element, N)
	}
}

func Benchmark_PureMakeCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		yForMakeCopy = make([]Element, N)
		copy(yForMakeCopy, xForMakeCopy)

	}
}
func Benchmark_PureAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		yForMakeCopy = append([]Element(nil), xForMakeCopy...)
	}
}
