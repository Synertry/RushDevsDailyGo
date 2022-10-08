### NR 24

---

#### Aufgabe
Gegeben ist eine Liste von Zahlen und eine Zielzahl n, finden Sie 3 Zahlen in der Liste, deren Summe der Zielzahl n am nächsten kommt. Es kann mehrere Möglichkeiten geben, die Summe zu bilden, die der Zielzahl am nächsten kommt, Du könntest jede Kombination in beliebiger Reihenfolge zurückgeben.


#### Start
Hier ist ein Beispiel und etwas Startcode:

```py
def closest_3sum(nums, target):
# Das Ausfüllen

print(closest_3sum([2, 1, -5, 4], -1))
# Nächste Summe ist -5+1+2 = -2 ODER -5+1+4 = 0, weil abstand bei beiden 1 ist!
# print [-5, 1, 2]
```


#### Hinweis
> -> nums.length > 2<br>
> -> nums.entries = Integer<br>
> -> typeof target = Integer<br>