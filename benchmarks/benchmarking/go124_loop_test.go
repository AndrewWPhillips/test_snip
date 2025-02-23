//go:build go1.24

package __

import (
	"testing"
)

func BenchmarkNotAliveInlineGo124(b *testing.B) { // 4
	for b.Loop() {
		_ = f()
	}
}

func BenchmarkNotAliveNotInlineGo124(b *testing.B) { // 4
	for b.Loop() {
		_ = g()
	}
}
