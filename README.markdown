gotweet
=======

Gotweet is a simple command-line Twitter client in [Go programming language](http://golang.org).

Building from Sources
---------------------

To make (you have to install [Go](http://golang.org) first):

	$ cd src
	$ make

(Note: currently it builds using 6g compiler. If you change Makefile to be compiler-independent, please fork and submit your pull request.)

Binaries
--------

* [gotweet-osx-amd64.zip](http://c0043312.cdn.cloudfiles.rackspacecloud.com/gotweet-osx-amd64.zip)
(Mac OS X 10.5 x86-64, 370 KB, Version [b975035f58], 2009-11-12 22:07:02)

* [gotweet-linux-amd64.zip](http://c0043312.cdn.cloudfiles.rackspacecloud.com/gotweet-linux-amd64.zip)
(Linux amd64, 370 KB, Version [b975035f58], 2009-11-12 22:07:02)

Usage
------

	Usage: ./gotweet [options...] action ...
	Options:
	  -p="": password
	  -u="": username (Twitter login)
	Actions:
	  post		Post status update (followed by status). Alias: p
	  user		Show user timeline. Alias: u
	  friends	Show friends timeline. Alias: (nothing)
	  mentions	Show mentions. Alias: @
	  public	Show public timeline

Examples
--------

Post status update:
	
	gotweet -u=username -p=password p "My status update"

Show mentions:

	gotweet -u=username -p=password @

Show friends timeline:

	gotweet -u=username -p=password

* * *

Made by [Coding Robots](http://www.codingrobots.com).