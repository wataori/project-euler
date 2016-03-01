use strict;
use warnings;
use utf8;
use bigint;

my $i = 2;
my $a = 1;
my $b = 1;
my $tmp = 1;

while (length($b) < 1000) {
  $tmp = $b;
  $b += $a;
  $a = $tmp;
  $i++;
}

printf("$i\n");
