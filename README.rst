gonetgrep
#########

.. contents::

Introduction
^^^^^^^^^^^^

I found Google GO is a interesting programming language.
I guess, it is very useful in network data mining.
This project is just for a sample to show Go's power.

The best method to learn a new language is to try to teach
others.  So, this project won't show up a finished code 
at once.  Intead, I'll keep the my learning experiences
on `Dummy Days`_ section of this document and the progress of this program 
on https://github.com/dlintw/gonetgrep.

BTW, I'm not a native English speaker.  I'm a newbie of Go.
There are many bugs exist on my code or grammars, please notice me.

This document could be converted to html by docutils, that let the content 
could be clickable. See `ReStructuredText`_ for more information.

Usage
^^^^^

I define the usage here::

  gonetgrep [options] <keyword> <url> [<url> ...]

  Grep keyword in multiple web pages.

  options:
    -t <num> 0: without parse table, 1:first table only

Dummy Days
^^^^^^^^^^

This is the section of my experience.
I stduy Golang by following methods:

1. Read document_ . (note: the Specification is also must read document)
2. Search question on golang-nuts_.
3. Find similar package on `project dashboard`_ and `packages dashboard`_.
4. Ask questions on golang-nuts.

.. _document: http://golang.org/doc/docs.html
.. _golang-nuts: http://groups.google.com/group/golang-nuts
.. _project dashboard: http://godashboard.appspot.com/project
.. _packages dashboard: http://godashboard.appspot.com/package

You could check out the source code in different stage. The method is describe
in `Version Control System`_.

D1 - Setup develop environment (cmd:godoc)
==========================================

In fact, I don't like the name of this language.
I would like it named as 'golang' which is more searchable.

I use archlinux_.  It is easy to install multiple newest packages by yaourt_::

  yaourt -S go-hg  # Google Go

  # optional packages
  yaourt -S gocode-git # suggest strongly for vim users
  # if gocode-git install failed, just skip it, I'll explain how to fix it.
  yaourt -S git mercurial # source code version control for goinstall
  yaourt -S vi vim-diff # my favorite editor
  yaourt -S docutils # convert this document into web page form (HTML).

.. _archlinux: http://www.archlinux.org
.. _yaourt: https://wiki.archlinux.org/index.php/Yaourt

You can read Golang's document without network by godoc_ ::

  godoc -http=:6060  # and launch in browser by http://localhost:6060
  godoc godoc # read more usage by this builtin document tool

In fact, if you use brand new version of Go.  You should reference package 
manual by this method instead of just read the offical site's manual.  
Because official site only keep stable version's package document.

By the way, I suggest to open a github_ account, and learn how to use 
git on github. And try to write document in rst_ format.

.. _github: https://github.com
.. _rst: http://docutils.sourceforge.net/docs/user/rst/quickref.html
.. _godoc: http://golang.org/cmd/godoc

D2 - Compile hello world (cmd:gofmt pkg:fmt,flag,os)
====================================================

Suject: compile hello world by Make.inc (cmd:gofmt pkg:fmt,flag,os)

Golang provide a **gofmt** utility to make same coding style.
We could try to copy the Makefile from gofmt::
  
  find /opt/go |grep gofmt

  mkdir gonetgrep
  cd gonetgrep
  cp /opt/go/src/cmd/gofmt/Makefile .

We could reference /opt/go/src/cmd/gofmt/gofmt.go to build our gonetgrep.go.
If gocode installed.  Here is the most important tip when edit by vim editor::

  package main
  import "flag"  // after declare the flag
  func main() {  // note Go's style, brace { should not in next line
    flag. // press Ctrl-X, Ctrl-O then Ctrl-P or Ctrl-N here
  }

Press Ctrl-X then Ctrl-O after type **flag.**, you could see flag's members.
If you want to know the usage of member functions, just look godoc.
To clear the automatic typing code, you could try Ctrl-P again.

A successful programming language should come with a powerful and useful 
library. 

We use the following package functions.

======== =============
C        Golang
printf() fmt.Println()
getopt() flag.Parse()
argv()   flag.Args()
exit()   os.Exit()
======== =============


Q&A
---

1. How to write long line string in [`fd2a code`_]?  

Ans. use back single quote or **+** operator (`Thank Arlen and PeterGo`_), this bug will cause the 
following warning::

  gonetgrep.go:17: syntax error: unexpected semicolon or newline, expecting )

This bug is fixed in [15dd_].

.. _fd2a code: https://github.com/dlintw/gonetgrep/blob/fd2a/gonetgrep.go#L18
.. _Thank Arlen and PeterGo: http://groups.google.com/group/golang-nuts/browse_thread/thread/a995c49934392b27
.. _15dd: https://github.com/dlintw/gonetgrep/commit/15dd

.. code time: Wed May  4 06:26:54 ~ 07:33:02 CST 2011

D3- Display Usage Depend on Locale (pkg:utf8,go-iconv)
======================================================

Golang suggest we use utf-8 as default.  So, if we want to display string,
we should code in utf-8.  For different terminal codec environment, we require
convert from utf-8 to encoding locale. There is no default convert package
in go package, so, I searched in http://godashboard.appspot.com/package.
I found there is two go-iconv package, choose the max count package and install::

  goinstall "github.com/sloonz/go-iconv"  # this line failed
  goinstall "github.com/sloonz/go-iconv/src" # it works

The finished code in

.. code time: Sat May  7 12:16:15~14:00 CST 2011

Appendix
^^^^^^^^

Version Control System
======================

You may see hex deciaml numbers like this [fd2a_].
That's the snapshot of source code at the moment with git version fd2a.

* To read the version's source tree in browser, just click the version.
* To read changes of this version, just click the **commit** on right side 
  after click the link.
* TO read commit log, click on github's **commit** button on upper bar.

.. _fd2a: https://github.com/dlintw/gonetgrep/tree/fd2a

To check the source code in your linux box, here are sample commands::

  # initial copy
  git clone git://github.com/dlintw/gonetgrep.git 
  cd gonetgrep

  # get update source
  git pull 

  # show commit log
  git log --all
  git log    # show current checkout version's log only.

  # update to special version, for example fd2a
  git checkout fd2a

  # back to newest version.
  git checkout HEAD

  # compare the differences of version fd2a and previous version(fd2a^)
  git diff fd2a fd2a^

Q&A
----

1. Why 'git ci' can not check in but 'git ci -a' can?  

Ans.  git's process force you separate a large patch into small pieces by manual add any 'add' or 'modify' patch. [#ga]_

.. [#ga] http://plasmasturm.org/log/gitidxpraise>
 
ReStructuredText
================

This document is written by ReStructuredText format which is used by python language.

This document could be converted to html by `docutils <http://docutils.sourceforge.net>`_.::
  rst2html README.rst README.html

Q&A
---

1. How to hightlight Go's syntax in rst format?


TODO
^^^^

I require help to finish all these jobs. If you can help me. Just fork my source, and notice me to pull your code and document.

* debug code (pkg:log,runtime)
* read test code of official packages (pkg:testing)
* read file line by line (pkg:io)
* find keyword and display line number (pkg:bytes,regexp)
* get web page (pkg:http)
* store to file (pkg:path)
* get multiple web page by goroutine (pkg:sync)
* store history into database (pkg:sqlite)
* get web pages through multiple agents (pkg:gob)
* show web robots's status on web
* build test case (pkg:testing)
* benchmark the code
* balance load of bottleneck
* prevent hardware fail by architecture

.. vim:set sw=2 ts=2 et sta:
