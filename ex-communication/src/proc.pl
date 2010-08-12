#!/usr/bin/perl -l
use strict;
use warnings;

my (%child, $pid, $parent);
my @ps=`ps -e -opid,ppid,comm`;
foreach (@ps[1..$#ps]) {
    ($pid, $parent, undef) = split;
    push @{$child{$parent}}, $pid;
}

foreach (sort { $a <=> $b } keys %child) {
    print "Pid ", $_, " has ", @{$child{$_}}+0, " child", 
	@{$child{$_}} == 1 ? ": " : "eren: ", "@{$child{$_}}";
}
