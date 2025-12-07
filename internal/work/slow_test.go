package work

import "testing"

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Fib(30) // поменьше, чтобы бенч шел быстрее
	}
}
