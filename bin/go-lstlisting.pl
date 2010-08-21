#!/usr/bin/perl
# extract Go code and compile it
use strict;
use warnings;

sub removeremark(@) {
    my @go;
    foreach(@_) {
	s/\|\\coderemark.*?\|//;
	s/\|\\longremark.*?\|//;
	push @go, $_;
    }
    @go;
}

my $inlisting = 0;
my @listing;
while(<>) {
    if (m|\\begin{lstlisting}|) {
	    $inlisting = 1;
	    next;
    }
    if (m|\\end{lstlisting}|) {
	    $inlisting = 0;
	    
	    if ( grep { /package main/ } @listing ) {
		print "// Full program\n";
	    } else {
		print "// Snippet\n";
	    }
	    @listing = removeremark(@listing);

	    print @listing;

	    @listing = ();
	    next;
    }
    if ($inlisting == 1) {
	push @listing, $_;
    }
}

