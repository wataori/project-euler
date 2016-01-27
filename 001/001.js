var sum = 0;

(function(cb) {
  for (var i = 1; i < 1000; i++) {
    if (i % 3 === 0) {
      sum += i;
    } else if (i % 5 === 0) {
      sum += i;
    }
  }
  cb();
})(function() {
  console.log(sum);
});
