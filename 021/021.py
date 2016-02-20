# coding: utf-8
import itertools
import math

def get_sum_of_divisors(n):
    sqrt = math.sqrt(n)
    if math.trunc(sqrt)**2 == n:
        _sum = math.trunc(sqrt)
    else:
        _sum = 0
    for i in range(1, math.trunc(sqrt)):
        if i == 1:
            _sum += 1
        elif n%i == 0:
            _sum += i + n/i
    return _sum

ans = 0

for i in range(2, 10000):
    a = get_sum_of_divisors(i)
    b = get_sum_of_divisors(a)
    if i != a and i == b:
        ans += i

print(ans)
