import itertools
import math

location = 1000000
ans = 0
digits = [1, 1, 1, 1, 1, 1, 1, 1, 1, 1]
targets= [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

for i in range(1, 11):
    x = 1
    for j in range(1, 11-i):
        x *= j
    digits[i-1] = targets[location//x]
    del targets[location//x]
    location = location%x

for i in range(0, 10):
    print(digits[i], end = '')
print()
