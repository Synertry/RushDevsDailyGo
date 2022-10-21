### NR 9

---

#### Aufgabe
> Finde bei einer gegebenen Zahlenliste und einer Zielzahl alle möglichen eindeutigen Teilmengen der Zahlenliste, die sich zu der Zielzahl summieren. Die Zahlen sind alle positive und ganze Zahlen.


#### Start
> Hier ist ein Beispiel und etwas Startcode:

```py
def sum_combinations(nums, target):
# Fill this in.

print(sum_combinations([10, 1, 2, 7, 6, 1, 5], 8))
# [(2, 6), (1, 1, 6), (1, 2, 5), (1, 7)]
# 2 + 6 = 8, 1 + 1 + 6 = 8, usw...
```


#### Hinweis
> -> typeof nums = Integer<br>
> -> target > 0<br>
> -> nums.length > 1<br>