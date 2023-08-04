package __

import (
	"fmt"
	"testing"
)

type S8 struct {
	a int
}

func BenchmarkAssignS8(b *testing.B) { // 0.42 ns/op
	outer1 := S8{a: 42}
	for i := 0; i < b.N; i++ {
		inner := S8{a: i}
		outer1 = inner
	}
	fmt.Println(outer1.a)
}

func BenchmarkAssignS8_2(b *testing.B) { // 0.43
	outer1 := S8{a: 42}
	outer2 := S8{a: 43}
	for i := 0; i < b.N; i++ {
		inner := S8{a: i}
		outer1 = inner
		outer2 = inner
	}
	fmt.Println(outer1.a, outer2.a)
}

// BenchmarkAssignClearS8 is an (unsuccessful) attempt to work out the cost of clearing a struct (diff with above test)
func BenchmarkAssignClearS8(b *testing.B) { // 0.35 - faster due to eliminating outer2 assignment from the loop?
	empty := S8{}
	outer1 := S8{a: 42}
	outer2 := S8{a: 43}
	for i := 0; i < b.N; i++ {
		inner := S8{a: i}
		outer1 = inner
		inner = empty
		outer2 = inner
	}
	fmt.Println(outer1.a, outer2.a)
}

func BenchmarkAssignClearS8_3(b *testing.B) { // 0.51
	empty := S8{}
	outer1 := S8{a: 42}
	outer2 := S8{a: 43}
	for i := 0; i < b.N; i++ {
		inner := S8{a: i}
		outer1 = inner
		inner = empty
		inner.a = i / 2
		outer2 = inner
	}
	fmt.Println(outer1.a, outer2.a)
}

func BenchmarkAssignS8_4(b *testing.B) { // 0.54
	//empty := S8{}
	outer1 := S8{a: 42}
	outer2 := S8{a: 43}
	for i := 0; i < b.N; i++ {
		inner := S8{a: i}
		outer1 = inner
		//inner = empty
		inner.a = i / 2
		outer2 = inner
	}
	fmt.Println(outer1.a, outer2.a)
}

type S32 struct {
	a, b, c, d int
}

func BenchmarkAssignS32(b *testing.B) { // 0.43
	outer1 := S32{a: 42}
	outer2 := S32{a: 43}
	for i := 0; i < b.N; i++ {
		inner := S32{a: i}
		outer1 = inner
		outer2 = inner
	}
	fmt.Println(outer1.a, outer2.a)
}

func BenchmarkAssignClearS32(b *testing.B) { // 0.35
	empty := S32{}
	outer1 := S32{a: 42}
	outer2 := S32{a: 43}
	for i := 0; i < b.N; i++ {
		inner := S32{a: i}
		outer1 = inner
		inner = empty
		outer2 = inner
	}
	fmt.Println(outer1.a, outer2.a)
}
