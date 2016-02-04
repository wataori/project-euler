var count = 0;
var prime = 2;

var isPrime = function(num) {
  if (num <= 2) {
    return true;
  }
  var sqrt = Math.floor(Math.sqrt(num));
  for (var i = 2; i < sqrt; i++) {
    if (num % i === 0) {
      return false;
    }
  }
  return true;
};

function generatePrime(num) {
  if (isPrime(num)) {
    if (++count % 1000 === 0 || count === 10001) {
      return num;
    }
  }
  return generatePrime(num + 1);
}

// use 'for' iteration because stack level is too deep
for(var i = 0; count !== 10001; i++) {
  prime = generatePrime(prime);
}

console.log(prime);
