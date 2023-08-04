package main

import (
	"fmt"
)

/* This is here because some things like "go build -gcflags -help" don't work unless there is a non-test .go file present.
 HELP:gcflags

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
*/

func main() {
	_, _ = fmt.Println("hello world")
}
