#!/usr/bin/perl
# extract Go code and compile it
use strict;
use warnings;

my $i = 1;
my $basedir=$ARGV[0] if defined $ARGV[0];
if (!defined $basedir) {
    die "No basedir";
}
shift;

# remove |coderemarks|
sub removeremark($) {
    my $file = shift;
    open GO, "+<", $file or die "Cannot open: $file";
    my @go;
    while(<GO>) {
	s/\|\\coderemark.*?\|//;
	s/\|\\longremark.*?\|//;
	push @go, $_;
    }
    truncate GO, 0;
    seek GO, 0, 0;
    foreach(@go) {
	print GO $_;
    }
    close GO;
}

sub compile($) {
    my $file = shift;
    my $target = "/tmp/$$.go.$i";
    $i++;
    system("cp $basedir/$file $target");
    removeremark($target);
    system("6g $target");
    unlink($target);
    $? >> 8;
}

my $inlisting = 0;
my @listing;
while(<>) {
    if (m|\\lstinputlisting(\[.*?\])?{(.*)}|) {
	my $gofile = $2;
	if ($gofile !~ /\.go$/) {
	    print "No Go    : $gofile\n";
	    next;
	}
	print "Compiling: 6g $gofile: ";
	if (compile($gofile) != 0) {
	    print "NOT OK\n";
	} else {
	    print "OK\n";
	}
    }
}

