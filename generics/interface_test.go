package __

type Reader[T any] interface {
	Read(p []T) int
}
