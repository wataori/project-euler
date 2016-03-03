limit = 1000
a = 1
b = 1
i = 2

while len(str(b)) < limit:
    a, b = b , a + b
    i += 1

print(i)
