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

gonetgrep [options] <keyword> <url> [<url> ...]

Grep keyword in multiple web pages.

Dummy Days
^^^^^^^^^^

This is the section of my experience.

Day 1
=====

Subject: install Go

In fact, I don't like the name of this language.
I would like it named as 'golang' which is more searchable.

I use archlinux.  It is easy::

  yaourt -S go-hg  # Google Go
  yaourt -S gocode-git # suggest strongly for vim users
  # if gocode-git install failed, just forgive it.

Subjects in the Future
======================

* compile hello world by Make.inc (pkg:flag)
* display usage depend on locale 
* read file line by line (pkg:io)
* find keyword and display line number (pkg:bytes,regexp)
* get web page (pkg:http)
* store to file (pkg:path)
* get multiple web page by goroutine (pkg:sync)
* store history into database (pkg:sqlite)
* get web pages through multiple agents (pkg:gob)
* build test case (pkg:testing)
* benchmark the code
* balance load of bottleneck
* prevent hardware fail by architecture

.. vim:set sw=2 ts=2 et sta:
