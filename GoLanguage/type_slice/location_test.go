package __

import (
	"testing"
)

// Check if a slice's backing array is on stack or heap

func TestSliceMemoryLocation(t *testing.T) {
	println(make([]byte, 1024))    // 0x1a46963fb57 stack
	println(make([]byte, 64*1024)) // stack
	// Allocate a slice bigger than MaxImplicitStackVarSize
	println(make([]byte, 64*1024+1)) // 0xc000126000 heap
	println([]byte{1024 * 1024: 10}) // 0x1a46963ff57 stack!!
}
