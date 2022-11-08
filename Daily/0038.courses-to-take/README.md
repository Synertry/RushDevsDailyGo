### NR 38

---

#### Aufgabe
> Du erhältst eine Hash-Tabelle, in der der Schlüssel ein Kurscode ist und der Wert eine Liste aller Kurscodes ist, die Voraussetzungen für den Schlüssel sind. Gebe eine gültige Reihenfolge zurück, in der wir die Kurse abschließen können. Wenn keine solche Reihenfolge vorhanden ist, gebe NULL zurück.


#### Start
> Hier ist ein Beispiel und etwas Startcode:

```py
def courses_to_take(course_to_prereqs):
# Das Ausfüllen.

courses = {
'CSC300': ['CSC100', 'CSC200'],
'CSC200': ['CSC100'],
'CSC100': []
}
print courses_to_take(courses)
# ['CSC100', 'CSC200', 'CSC300']
```

#### Beispiel
> Input:<br>
`{`<br>
&emsp;`'CSC300': ['CSC100', 'CSC200'],`<br>
&emsp;`'CSC200': ['CSC100'],`<br>
&emsp;`'CSC100': []`<br>
`}`<br>
> Output:<br>
>Diese Eingabe sollte die Reihenfolge zurückgeben, die wir benötigen, um diese Kurse zu belegen:<br>
`['CSC100', 'CSC200', 'CSCS300']`


#### Hinweis
> -> typeof courses = Object<Array>
