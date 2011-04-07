#!/usr/bin/perl
# extract Go code and compile it
use strict;
use warnings;

my $quiet = 0;

sub removeremark(@) {
    my @go;
    my $skip = 0;
    foreach(@_) {
	if ($skip) {
	    # check if we have % again, if so, skip again
	    if (/\%$/) { next; }
	    $skip = 0;
	    next;
	}

	if (/\%$/) { $skip = 1; next; }
	if (/\|\\begin/) { $skip = 1; next; }
	if (/\|\\draw/) { $skip = 1; next; }
	if (/\\end/) { $skip = 0; next; }


	s/\%[^a-zA-Z].*$//;
	s/\\draw.*$//;
	s/\|\\coderemark.*?(\||\%$|$)//;
	s/\|\\longremark.*?(\||\%$|$)//;
	s/\\newline//;
	s/\|.*\|//;
	push @go, $_;
    }
    @go;
}

sub nl(@) {
    my $i = 0;
    foreach(@_) {
	print ++$i, "\t", $_;
    }
}

sub gofmt(@) {
    open FMT, "|-", "gofmt > /dev/null";
	foreach(@_) {
	    print FMT  $_;
	}
    close FMT;
    $?;
}

my $inlisting = 0;
my @listing;
my ($snip, $func);
my (@func, @snip);
while(<>) {
    if (m|\\begin{lstlisting}|) {
	    $inlisting = 1;
	    next;
    }
    if (m|\\end{lstlisting}|) {
	    $inlisting = 0;
	    
	    @listing = removeremark(@listing);

	    if ( grep { /package main/ } @listing ) {

		unshift @listing, "// Full program\n";
		print "FULL FULL FULL -- " if not $quiet;
		if (gofmt(@listing) != 0) {
		    print "\n";
		    nl @listing;
		    print "$ARGV - NOT OK\n";
		} else {
		    print "OK\n" if not $quiet;
		}

	    } elsif ( grep { /func .*?\(/ } @listing ) {
		push @func, "// $ARGV - Function " . ++$func . "\n";
		grep { s/(import)/\/\/$1/}  @listing;
		@func = (@func, @listing);
	    } else {
		push @snip, "// $ARGV - Snippet " . ++$snip . "\n";
		grep { s/(import)/\/\/$1/}  @listing;
		@snip = (@snip, @listing);
	    }
	    @listing = ();
	    next;
    }
    if ($inlisting == 1) {
	push @listing, $_;
    }
}
# snippets
unshift @snip, <<EOF;
package main

import (
    "fmt"
)

func main() {	// START

EOF
push @snip, "}   // END\n";
unshift @func, <<EOF;
package main
func main () { }
EOF

print "SNIP SNIP SNIP -- " if not $quiet;
if (gofmt(@snip) != 0) {
    print "\n";
    nl @snip;
    print "NOT OK\n";
} else {
    print "OK\n" if not $quiet;
}
print "FUNC FUNC FUNC -- " if not $quiet;

if (gofmt(@func) != 0) {
    print "\n";
    nl @func;
    print "NOT OK\n";
} else {
    print "OK\n" if not $quiet;
}
