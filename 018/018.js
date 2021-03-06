"use strict";

const data = [
  [75], // 0
  [95, 64], // 1
  [17, 47, 82], // 2
  [18, 35, 87, 10], // 3
  [20, 4, 82, 47, 65], // 4
  [19, 1, 23, 75, 3, 34], // 5
  [88, 2, 77, 73, 7, 63, 67], // 6
  [99, 65, 4, 28, 6, 16, 70, 92], // 7
  [41, 41, 26, 56, 83, 40, 80, 70, 33], // 8
  [41, 48, 72, 33, 47, 32, 37, 16, 94, 29], // 9
  [53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14], // 10
  [70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57], // 11
  [91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48], // 12
  [63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31], // 13
  [4, 62, 98, 27, 23, 09, 70, 98, 73, 93, 38, 53, 60, 4, 23] // 14
];
let results = [];

const select = function(idx, x, sum) {
  if (idx===0) {
    select(idx+1, x, sum+data[idx][x]);
  } else if (idx===15) {
    return results.push(sum);
  } else {
    select(idx+1, x, sum+data[idx][x]);
    select(idx+1, x+1, sum+data[idx][x+1]);
  }
}

select(0, 0, 0);
console.log(Math.max.apply({}, results));
