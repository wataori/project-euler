<?php
// $input = (int)trim(fgets(STDIN));
$input = 1000;
$sqrt = sqrt($input);
$a = $b = $c = 0;

for ($i = 1; $i <= $sqrt - 2; $i++) {
  $a = $i * $i;
  for ($j = 2; $j <= $sqrt - 1; $j++) {
    $b = $j * $j;
    for ($k = 3; $k <= $sqrt; $k++) {
      $c = $k * $k;
      if ($a + $b + $c === $input) {
        break 3;
      }
    }
  }
}

$result = $a * $b * $c;
echo "$result\n";
