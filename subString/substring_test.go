package subString

import (
	"testing"
	"unicode/utf8"
)

var benchMarkString = "Go语言是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。为了方便搜索和识别，有时会将其称为Golang。"
var benchMarkStringLen = 20

func SubStrRange(s string, length int) string {
	var n, i int
	for i = range s {

		if n == length {
			break
		}
		n++
	}

	return s[:i]
}

func SubStrDecodeRuneInString(s string, length int) string {
	var size, n int
	for i := 0; i < length && n < len(s); i++ {
		_, size = utf8.DecodeLastRuneInString(s[n:])
		n += size
	}
	return s[:n]
}

func BenchmarkSubStrRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubStrRange(benchMarkString, benchMarkStringLen)
	}
}

func BenchmarkDecodeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubStrDecodeRuneInString(benchMarkString, benchMarkStringLen)
	}
}
