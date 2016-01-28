var results = [];
const start = 100 * 100;
const end = 999 * 999;
for (var i = start; i < end; i++) {
  var tmp = i.toString();
  if (tmp[0] === tmp[tmp.length-1] && tmp[1] === tmp[tmp.length-2]) {
    results.push(i);
  }
}
console.log(results[results.length-1]);
