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
on **Dummy Days** sections and github.com.

BTW, I'm not native English speaker.  I'm a newbie of Go.
There are many bugs exist on my code or grammars, please notice me.

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

Day 1
=====

Subject: install Go's develop environment

In fact, I don't like the name of this language.
I would like it named as 'golang' which is more searchable.

I use archlinux.  It is easy to install multiple newest packages::

  yaourt -S go-hg  # Google Go

  # optional packages
  yaourt -S gocode-git # suggest strongly for vim users
  # if gocode-git install failed, just skip it, I'll explain how to fix it.
  yaourt -S git mercurial # source code version control for goinstall
  yaourt -S vi vim-diff # my favorite editor
  yaourt -S docutils # convert this document into web page form (HTML).

By the way, I suggest to open a github_ account, and learn how to use 
git on github. And try to write document in rst_ format.

.. _github: https://github.com
.. _rst: http://docutils.sourceforge.net/docs/user/rst/quickref.html

Day 2
=====
Suject: compile hello world by Make.inc (cmd:gofmt pkg:flag)

Golang provide a **gofmt** utility to make same coding style.
We could try to copy the Makefile from gofmt::
  
  find /opt/go |grep gofmt

  mkdir gonetgrep
  cd gonetgrep
  cp /opt/go/src/cmd/gofmt/Makefile .

time: Wed May  4 06:26:54 CST 2011 ~

Subjects in the Future
======================

* display usage depend on locale 
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
