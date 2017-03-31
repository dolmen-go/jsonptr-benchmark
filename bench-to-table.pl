#!/usr/bin/env perl

# Takes the "go test -bench B -benchmem" output and format in GitHub-flavored Markdown

use 5.010;
use strict;
use warnings;

print <<EOF;
| Benchmark | impl | op | speed | allocs bytes | allocs count |
| --- | --- | ---: | ---: | ---: | ---: |
EOF

while (<>) {
    next unless /^B/;
    chomp;
    my @cols = split /\s\s+/;
    unshift @cols, (shift @cols) =~ m{^(.+)/\[([^]]+)};
    print '| ', join(' | ', @cols), " |\n";
}
