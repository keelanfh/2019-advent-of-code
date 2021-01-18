from collections import Counter

p = q = 0
for x in range(123257, 647015):
    v = Counter(s := str(x)).values()
    if all(j >= i for i, j in zip(s, s[1:])):
        p += max(v) >= 2
        q += 2 in v

print(p)
print(q)
