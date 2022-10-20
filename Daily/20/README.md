### NR 20

---

#### Aufgabe

> Finde bei einer sortierten Liste mit Duplikaten und einer Zielzahl n den Bereich, in dem die Zahl vorhanden ist (dargestellt als Tupel (niedrig, hoch), beide einschließlich. Wenn die Zahl nicht in der Liste vorhanden ist, gebe (-1, -1)) zurück.

#### Start

> Hier ist ein Beispiel und etwas Startcode:

```py
def find_num(nums, target):
# Das Ausfüllen

print(find_num([1, 1, 3, 5, 7], 1))
# (0, 1) -> bei Index 0 und Index 1 sind die Duplikate

print(find_num([1, 2, 3, 4], 5))
# (-1, -1)
```

#### Hinweis

> -> nums.entries = Integer<br>
> -> nums.length > 2<br>
