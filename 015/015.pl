#!/usr/bin/perl

use strict;
use warnings;
use utf8;
use bigint;

use Data::Dumper;

my @grid = ();

for (my $i = 0; $i < 21; $i++) {
  for (my $j = 0; $j < 21; $j++) {
    if ($i > 0 && $j > 0) {
      $grid[$i][$j] = $grid[$i-1][$j] + $grid[$i][$j-1];
    } else {
      $grid[$i][$j] = 1;
    }
  }
}

printf("$grid[20][20]\n");
