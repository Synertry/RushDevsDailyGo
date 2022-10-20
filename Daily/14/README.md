### NR 14

---

#### Aufgabe

> Finde in einer gegebenen Liste positiver Zahlen die größtmögliche Menge, sodass keine Elemente benachbarte Zahlen (adjacent numbers) voneinander sind, das heißt, dass Zahlen, die direkt nebeneinander sind, nicht summiert werden dürfen -> [1, 7, 1] => Höchste Summe 2.

#### Start

> Hier ist ein Beispiel und etwas Startcode:

```py
def maxNonAdjacentSum(nums):
# Das ausfüllen

print(maxNonAdjacentSum([3, 4, 1, 1]))
# 5
# maximale sum ist 4 (index 1) + 1 (index 3)

print(maxNonAdjacentSum([2, 1, 2, 7, 3]))
# 9
# maximale sum ist 2 (index 0) + 7 (index 3)
```

#### Hinweis

> -> nums.length >= 3<br>
> -> nums.entries = Integer<br>
