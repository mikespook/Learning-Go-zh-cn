#!/usr/bin/perl -l
my (%child, $pid, $parent);
my @ps=`ps -e -opid,ppid,comm`;	    # Capture the output from `ps`
foreach (@ps[1..$#ps]) {	    # Discard the header line
    ($pid, $parent, undef) = split; # Split the line, discard 'comm'
    push @{$child{$parent}}, $pid;  # Save the child PIDs on a list
}
# Walk through the sorted PPIDs
foreach (sort { $a <=> $b } keys %child) {  
    print "Pid ", $_, " has ", @{$child{$_}}+0, " child",  # Print them
	@{$child{$_}} == 1 ? ": " : "ren: ", "[@{$child{$_}}]";
}
