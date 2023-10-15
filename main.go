package main

// This (main.go file) is here because some things like "go build -gcflags -help" don't work unless there is a non-test .go file present.

func main() {
	test_cgo()

	test_go_release()
}

/*
Useful options:
build   -a -n -v -x -race -tags
gcflags -l -m -S
ldflags -c -s -w

BTW here are some command line options

$ go build -gcflags -help

-B	disable bounds checking
-D path
  	set relative path for local imports
-I directory
  	add directory to import search path
-N	disable optimizations
-S	print assembly listing
-V	print version and exit
-W	debug parse tree after type checking
-buildid id
  	record id as the build id in the export metadata
-c int
  	concurrency during compilation (1 means no concurrency) (default 12)
-e	no limit on number of errors reported
-l	disable inlining
-l -l enable inlining
-l -l -l more inlining
-m	print optimization decisions
-o file
  	write output to file
-race
  	enable race detector
-smallframes
  	reduce the size limit for stack allocated objects
-t	enable tracing for debugging the compiler
-traceprofile file
  	write an execution trace to file
-trimpath prefix
  	remove prefix from recorded source file paths
-v	increase debug verbosity
-wb
  	enable write barrier (default true)

$ go build -ldflags -help

# github.com/andrewwphillips/test_snip
usage: link [options] main.o
  -B note
        add an ELF NT_GNU_BUILD_ID note when using ELF
  -E entry
        set entry symbol name
  -H type
        set header type
  -I linker
        use linker as ELF dynamic linker
  -L directory
        add specified directory to library path
  -R quantum
        set address rounding quantum (default -1)
  -T address
        set text segment address (default -1)
  -V    print version and exit
  -X definition
        add string value definition of the form importpath.name=value
  -a    no-op (deprecated)
  -asan
        enable ASan interface
  -aslr
        enable ASLR for buildmode=c-shared on windows (default true)
  -benchmark string
        set to 'mem' or 'cpu' to enable phase benchmarking
  -benchmarkprofile base
        emit phase profiles to base_phase.{cpu,mem}prof
  -buildid id
        record id as Go toolchain build id
  -buildmode mode
        set build mode
  -c    dump call graph
  -compressdwarf
        compress DWARF if possible (default true)
  -cpuprofile file
        write cpu profile to file
  -d    disable dynamic executable
  -debugnosplit
        dump nosplit call graph
  -debugtextsize int
        debug text section max size
  -debugtramp int
        debug trampolines
  -dumpdep
        dump symbol dependency graph
  -extar string
        archive program for buildmode=c-archive
  -extld linker
        use linker when linking in external mode
  -extldflags flags
        pass flags to external linker
  -f    ignore version mismatch
  -g    disable go package data checks
  -h    halt on error
  -importcfg file
        read import configuration from file
  -installsuffix suffix
        set package directory suffix
  -k symbol
        set field tracking symbol
  -libgcc string
        compiler support lib for internal linking; use "none" to disable
  -linkmode mode
        set link mode
  -linkshared
        link against installed Go shared libraries
  -memprofile file
        write memory profile to file
  -memprofilerate rate
        set runtime.MemProfileRate to rate
  -msan
        enable MSan interface
  -n    dump symbol table
  -o file
        write output to file
  -pluginpath string
        full path name for plugin
  -r path
        set the ELF dynamic linker search path to dir1:dir2:...
  -race
        enable race detector
  -s    disable symbol table
  -strictdups int
        sanity check duplicate symbol contents during object file reading (1=warn 2=err).
  -tmpdir directory
        use directory for temporary files
  -v    print link trace
  -w    disable DWARF generation


*/
