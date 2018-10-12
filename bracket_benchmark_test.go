package bracket

import "testing"

func BenchmarkMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Match("[][]{}{}()(){{[()]}}")
	}
}

func BenchmarkMatch2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Match2("[][]{}{}()(){{[()]}}")
	}
}
