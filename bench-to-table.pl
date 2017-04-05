#!/usr/bin/env perl

# Takes the "go test -bench B -benchmem" output and format in GitHub-flavored Markdown

use 5.010;
use strict;
use warnings;


my $bench = '';
my @benchs;
my ($bests, $results);

while (<>) {
    next unless /^B/;
    chomp;
    my @cols = split /\s\s+/;

    # Split benchmark name in 2 columns
    unshift @cols, (shift @cols) =~ m{^(.+)/\[([^]]+)};
    # Delete column with operations count
    splice @cols, 2, 1;

    my $name = $cols[0];
    my $impl = $cols[1];

    if ($name ne $bench) {
        push @benchs, {
            name => $name,
            results => ($results = {}),
            bests => ($bests = [ map { (1e1000, '') } 0..2 ]),
        };
        $bench = $name;
    }
    for my $i (0..2) {
        (my $score = $cols[2+$i]) =~ s/ .*$//;
        if ($score < $bests->[2*$i]) {
            @{$bests}[2*$i, 2*$i+1] = ($score, $cols[2+$i])
        }
    }

    $results->{$impl} = [ @cols[2..4] ];
}

my $header = <<EOF;
| Impl | speed | allocs bytes | allocs count |
| --- | ---: | ---: | ---: |
EOF

foreach $bench (@benchs) {
    my ($name, $results, $bests) = @{$bench}{qw<name results bests>};
    print "#### $name\n\n$header";
    $name =~ s/~/\\~/g;
    my $min_score = 2;
    foreach my $impl (sort keys %$results) {
        my $result = $results->{$impl};
        my $score = 0;
        my @cols;
        for (0..2) {
            my $bold = '';
            if ($result->[$_] eq $bests->[2*$_+1]) {
                $bold = '**';
                $score++;
            }
            push @cols, "$bold$result->[$_]$bold";
        }
        print '| ', join(' | ', ($score >= $min_score ? "**$impl**" : $impl), @cols), " |\n";
        $min_score = $score if $score > $min_score;
    }
    print "\n";
}
