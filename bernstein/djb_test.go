package djb2

import (
"testing"
)

func BenchmarkHash(b *testing.B) {
	str:= "hello world"

for i := 0; i < b.N; i++ {
		Hash(str)
	}
}
