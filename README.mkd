# Learning Go - a free E-Book for learning the Go language.

This is an introduction in the language Go from Google.

The book currently consists out +/- 120 (A4 sized) pages and the following chapters:

Introduction

:   Show how to install Go and details the lineage of the language Go.

Basics

:   Types, variables and control structures.

Functions

:   How to make and use functions.

Packages

:   Functions and data is grouped together in packages. Here you will see how to make your own package. 
    How to unit test your package is also described.

Beyond the basics

:   Learn how to create your own data types and define function on them (called methods in Go).

Interfaces

:   Go does not support Object Orientation in the traditional sense. In Go the central concept is interfaces.

Concurrency

:   With the go keyword function can be started in separate routines 
    (called goroutines). Communication with those goroutines is done via channels.

Communication

:   How to create/read/write from and to files. And how to do networking.

Each chapter concludes with a number of exercises (and answers) to may help you to get some hands on experience. Currently it has more
than 30 exercises.

* See http://www.golang.org for the homepage of Go;
* Download the `pdf`s from: http://miek.nl/files/go/ ;
* Project page: http://www.miek.nl/projects/learninggo/index.html .

## Building
When building this from scratch you will need the
following packages on Debian/Ubuntu:

* `inkscape`
* `gnumeric`
* `ttf-droid`
* `ttf-dejavu `
* `ttf-sazanami-gothic`
* `texlive-fonts-recommended`
* `texlive-extra-utils`
* `texlive-xetex`
* `texlive-latex-extra`
* `texlive-latex-recommended`
* `latex-cjk-xcjk`
* `git-core`
* `make`

After that, you will hopefully only have to give `make`.
If not, so the pre-build `pdf`s on http://www.miek.nl/files/go/
