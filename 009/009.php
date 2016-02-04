<?php
// $input = (int)trim(fgets(STDIN));
$input = 1000;

for ($i=1; $i<=$input-2; $i++) {
  for ($j=$i+1; $j<=$input-$i-1; $j++) {
    $k = $input-($i+$j);
    if (($i*$i)+($j*$j) === $k*$k) {
      $result = $i*$j*$k;
      echo "$result\n";
      break 2;
    }
  }
}
