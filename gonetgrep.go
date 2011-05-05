package main

import (
	"fmt"
	"flag"
	"os"
)

// for more command line options example, ref: /opt/go/src/cmd/gofmt/gofmt.go
var gIsGrepTable = flag.Int("t", 0,
	"scan <TABLE>: 0:disable, 1:scan first table, 2:scan 2nd")

var gExitCode = 0

func usage() {
	// instead of Fprintf, try to use Fprintln, it is easier and powerful
	fmt.Fprintln(os.Stderr, "usage: gonetgrep [flags] <key> url [url...]"+
		"\nGrep keyword in multiple web pages.")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		usage()
	}
	println("This is first code Go support utf-8, 也可以用中文寫")
	os.Exit(gExitCode) // go can not return code to shell directly
}
