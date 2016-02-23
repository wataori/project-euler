# coding: utf-8
import itertools
import math

def get_sum_of_divisors(n):
    sqrt = math.sqrt(n)
    _sum = 0
    for i in range(1, math.trunc(sqrt)+1):
        if i == 1:
            _sum += 1
        elif n%i == 0:
            if i**2 == n:
                _sum += i
            else:
                _sum += i + n/i
    return _sum

ans = 0
for i in range(2, 28123):
    a = get_sum_of_divisors(i)
    if i < a:
        ans += i

print(ans)
