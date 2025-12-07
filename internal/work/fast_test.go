package work

import "testing"

func BenchmarkFibFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FibFast(30) // поменьше, чтобы бенч шел быстрее
	}
}
