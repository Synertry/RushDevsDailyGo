### NR 16

---

#### Aufgabe

> Generiere aus einer Liste eindeutiger Nummern (es gibt keine Duplikate) alle möglichen Teilmengen ohne Duplikate. Dazu gehört auch die leere Menge.<br>
> Beispiel für [0, 1]:<br>
> Alle möglichen Teilmengen: [], [1], [0], [1, 0], [0, 1]<br>

#### Start

> Hier ist ein Beispiel und etwas Startcode:

```py
def generateAllSubsets(nums):
# Das Ausfüllen

print(generateAllSubsets([1, 2, 3]))
# [[], [3], [2], [2, 3], [1], [1, 3], [1, 2], [1, 2, 3]]
```

Hinweis
-> nums.length > 0
-> nums.entries = Integer
-> jede num in nums existiert 1 mal (eindeutig)
