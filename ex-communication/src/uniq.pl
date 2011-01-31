#!/usr/bin/perl
my @a = qw/a b a a a c d e f g/;
print my $first = shift @a; 
foreach (@a) {
    if ($first ne $_) { print; $first = $_; }
}
