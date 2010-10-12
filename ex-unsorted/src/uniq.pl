#!/usr/bin/perl
my @a = qw/a b a a a c d e f g/;
my $last = shift @a; print $last;

foreach (@a) {
    if ($last ne $_) {
	print $_;
	$last = $_;
    }
}
