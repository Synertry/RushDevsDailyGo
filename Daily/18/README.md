### NR 18

---

#### Aufgabe

> Bei einer Matrix, die so organisiert ist, dass die Zahlen immer von links nach rechts sortiert werden und die erste Zahl jeder Zeile immer größer ist als das letzte Element der letzten Zeile (mat[i][0] > mat[i - 1][-1]), nach einem bestimmten Wert in der Matrix suchen und zurückgeben, ob er existiert.

#### Start

> Hier ist ein Beispiel und etwas Startcode:

```py
def searchMatrix(mat, value):
# Das Ausfüllen

mat = [
[1, 3, 5, 8],
[10, 11, 15, 16],
[24, 27, 30, 31],
]

print(searchMatrix(mat, 4))
# False

print(searchMatrix(mat, 10))
# True -> 10 ist in der Matrix
```

#### Hinweis

> -> mat[i][0] > mat[i - 1][-1]<br>
> -> mat.entries = Integer<br>
