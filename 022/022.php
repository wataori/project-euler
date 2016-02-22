<?php
$fp = fopen("./data.txt", "r");
$data = fgetcsv($fp, 20);
sort($data);
$keys = str_split("ABCDEFGHIJKLMNOPQRSTUVWXYZ");

$score_map = array();
for ($i=0; $i<count($keys); $i++) {
  $score_map[$keys[$i]] = $i+1;
}

$ans = 0;
for ($i=0; $i < count($data); $i++) {
  $chars = str_split($data[$i]);
  $score = 0;
  for ($j=0; $j < count($chars); $j++) {
    $score += $score_map[$chars[$j]];
  }
  $score *= ($i+1);
  $ans += $score;
}
print("$ans\n");
