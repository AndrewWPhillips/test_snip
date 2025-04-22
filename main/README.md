This main directory is just for single file programs, unlike all the other directories that have standalone tests.  This is done for various reasons - eg to run pprof we need a .exe.

Each .go file is a complete, simple program - ie with a main() func.  So you can't build all the file, as the main package, since there will be duplicate functions called "main".

Here are some examples working with flame.go, which also show how to use pprof.

# Create flame.exe from flame.go
$ go build flame.go

# Build and run flame.exe
$ go run flame.go

# Generate profile graph (SVG format) in flame.svg
$ go tool pprof -svg -output flame.svg flame.exe cpu.pprof

# View profile CLI - commands: top, tree/peek
$ go tool pprof cpu.pprof

# View profile Web
$ go tool pprof -http : cpu.pprof

