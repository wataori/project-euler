var sum_of_square = 0;
var square_of_sum = 0;
var sum = 0;

for (var i = 1; i <= 100; i++) {
  sum += i;
  sum_of_square += i * i;
}

square_of_sum = sum * sum;

console.log(square_of_sum - sum_of_square);
