This main directory is just for single file programs - ie, self-contained .go files with a main() func.

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
