package __

// TestGlobal is just here for assigning calculated values to.
// This is used in benchmarks to ensure that code is not optimised away, since if you assign a value to
// a variable visible outside the package the compiler doesn't know it's not used anywhere.
var TestGlobal any
