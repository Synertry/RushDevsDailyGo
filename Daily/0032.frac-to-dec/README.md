### NR 32

---

#### Aufgabe
> Finde bei gegebenem Zähler und Nenner heraus, was die äquivalente Dezimaldarstellung als Zeichenfolge ist. Wenn die Dezimaldarstellung wiederkehrende Ziffern hat, setzen Sie diese Ziffern in Klammern (d.h 4/3 sollte durch 1.(3) dargestellt werden, um 1,333 darzustellen ...). Verwende keine eingebauten Evaluator-Funktionen wie eval von Python. Du kannst auch davon ausgehen, dass der Nenner nicht Null ist.


#### Start
> Hier ist ein Beispiel und etwas Startcode:

```py
def frac_to_dec(numerator, denominator):
  # Das Ausfüllen

print(frac_to_dec(-3, 2))
# -1.5 

print(frac_to_dec(4, 3))
# 1.(3) -> 1,333333333 (Periode)

print(frac_to_dec(1, 6))
# 0.1(6)
```


#### Hinweis
> -> denominator != 0<br>
> -> typeof denominator = Integer<br>
> -> typeof numerator = Integer<br>
