#!perl
use strict;
use warnings;

use Test::More;

use Sereal::Merger;

{
    my $test = Sereal::Merger::test_me();
    print "--[$test]\n";
}

{
    my $test = Sereal::Merger->new({});
    print Dumper($test); use Data::Dumper;
}

pass;
done_testing;