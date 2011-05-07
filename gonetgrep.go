package main

import (
	"fmt"
	"flag"
	"os"
	"strings"
	"go-charset.googlecode.com/hg/charset"
	"bytes"
	"io"
)

// for more command line options example, ref: /opt/go/src/cmd/gofmt/gofmt.go
var gIsGrepTable = flag.Int("t", 0,
	"scan <TABLE>: 0:disable, 1:scan first table, 2:scan 2nd")

var gExitCode = 0

var gChTransTo charset.Translator

func init() {
	cs := charset.Info("big5")
	if cs == nil {
		println("Err: big5 charset not found")
		os.Exit(1)
	}
	var err os.Error
	if cs.TranslatorTo == nil {
		println("Err: utf8 to big5 translator not implement")
		os.Exit(1)
	}
	gChTransTo, err = cs.TranslatorTo()
        if err != nil {
                println("Err: making translator from %q: %v", cs, err)
		os.Exit(1)
        }
}

func utf82big5(in string) string {
        var buf bytes.Buffer
        r := charset.NewTranslatingReader(strings.NewReader(in), gChTransTo)
        _, err := io.Copy(&buf, r)
        if err != nil {
		println("Err: translating from %s: %v", in, err)
		os.Exit(1)
        }
        out := string(buf.Bytes())
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
