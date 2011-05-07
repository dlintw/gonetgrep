package main

import (
	"fmt"
	"flag"
	"os"
	"strings"
	iconv "github.com/sloonz/go-iconv/src"
)

// for more command line options example, ref: /opt/go/src/cmd/gofmt/gofmt.go
var gIsGrepTable = flag.Int("t", 0,
	"scan <TABLE>: 0:disable, 1:scan first table, 2:scan 2nd")

var gExitCode = 0

func utf82big5(s string) string {
	codec, err := iconv.Open("big5", "utf-8") // from utf8 to big5
	if err != nil {
		println(err)
		os.Exit(1)
	}
	defer codec.Close()
	out, err := codec.Conv(s)
	return out
}

func usage() {
	if strings.Contains(os.Getenv("LC_CTYPE"), "zh_TW") {
		s := "用法: gonetgrep [選項] <查詢字> 網址 [網址...]" +
			"\n抓取多個網站上的關鍵字"
		if strings.Contains(os.Getenv("LC_CTYPE"), "Big5") {
			fmt.Fprintln(os.Stderr, utf82big5(s))
		} else {
			fmt.Fprintln(os.Stderr, s)
		}
	} else {
	// instead of Fprintf, try to use Fprintln, it is easier and powerful
	fmt.Fprintln(os.Stderr, "usage: gonetgrep [flags] <key> url [url...]"+
		"\nGrep keyword in multiple web pages.")
	}
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
