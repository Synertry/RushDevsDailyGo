### NR 5

---

#### Aufgabe
> Gegeben ist ein String s und ein Zeichen c, finde den Abstand für alle Zeichen zum Zeichen c im String s. Du kannst davon ausgehen, dass das Zeichen c mindestens einmal in der Zeichenfolge vorkommt. Bsp: string = 'el' c='l', output: [1, 0] -> 1: Distanz zwischen e und l ist 1, zwischen l und l ist 0.

#### Start
> Hier ist ein Beispiel und etwas Startcode:

```py
def shortest_dist(s, c):
# Fill this in.

print(shortest_dist('helloworld', 'l'))
# [2, 1, 0, 0, 1, 2, 2, 1, 0, 1]
# -> h zu l: 2, e zu l: 1, l zu l: 0, l zu l: 0, o zu l: 1, usw...
```

#### Hinweis
> -> c kommt mindestens 1 mal im String vor
> -> String.length > 0
> -> Gebe alle Distanzen in einem Array aus (von c bis zu jeden Zeichen den Abstand)