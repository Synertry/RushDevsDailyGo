### NR 34

---

#### Aufgabe
> Du erhälst eine Zeichenfolge s und eine Ganzzahl k. Gib die Länge der längsten Teilzeichenfolge in s zurück, die höchstens k verschiedene Zeichen enthält.
>
> Zum Beispiel angesichts der Zeichenfolge:<br>
> aabcdefff und k = 3, dann wäre der längste Teilstring mit 3 verschiedenen Zeichen "defff". Die Antwort sollte 5 sein, da "defff" 3 verschiedene Zeichen hat und 5 Länge.


#### Start
>Hier ist ein Beispiel und etwas Startcode:

```py
def longest_substring_with_k_distinct_characters(s, k):
# Das Ausfüllen

print longest_substring_with_k_distinct_characters('aabcdefff', 3)
# 5 (weil'defff' eine Länge von 5 hat und 3 verschiedene Zeichen -> d, e, f)

print longest_substring_with_k_distinct_characters('aabcdefff', 1)
# 3 (weil 'fff' eine Länge von 3 hat und 1 Zeichen -> f)
```


#### Hinweis
> -> s.length > k<br>
> -> typeof k = Integer<br>
> -> typeof s = String<br>
> -> k > 0<br>