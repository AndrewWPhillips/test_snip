package __

/*
import "C"

//#include <stdio.h>
//extern int pr(const char *s, const int *a, int *p, double f);  // extern C function (see pr.c)
//int cfunc(const char *s, const int *a) { printf("<%s> %d %d [%d %d %d]\n", s, *a, *(a+2), sizeof(short), sizeof(int), sizeof(long)); }
import "C"

import (
	"testing"
	"unsafe"
)

func TestCGOTest(t *testing.T) {
	C.cfunc(C.CString("test string"), (*C.int)(unsafe.Pointer(&([]int{1, 2, 3, 4}[0]))))
	C.cfunc(C.CString("test string"), (*C.int)(unsafe.Pointer(&([]int32{1, 2, 3, 4}[0])))) // use this: C.int is 32 bits
	C.cfunc(C.CString("test string"), (*C.int)(unsafe.Pointer(&([]int16{1, 2, 3, 4, 5, 6}[0]))))
}
*/
