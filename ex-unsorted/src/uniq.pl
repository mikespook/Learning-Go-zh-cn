#!/usr/bin/perl
my @a = qw/a b a a a c d e f g/;
my $first = shift @a; print $first;

foreach (@a) {
    if ($first ne $_) {
	print $_;
	$first = $_;
    }
}
