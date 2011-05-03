# this makefile is copied from /opt/go/src/cmd/gofmt/Makefile

include $(GOROOT)/src/Make.inc

TARG=gonetgrep
GOFILES=\
	gonetgrep.go

include $(GOROOT)/src/Make.cmd

test: $(TARG)
	./test.sh

testshort:
	gotest -test.short
