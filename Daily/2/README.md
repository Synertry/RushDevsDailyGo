### NR 2

---

#### Aufgabe

> Finde anhand einer Liste von Meetings, die an einem Tag stattfinden, die Mindestanzahl an Besprechungsräumen, die für alle Meetings nötig sind.
>
> Jedes Meeting wird durch ein Tupel von (start_time,end_time) dargestellt, wobei sowohl start_time als auch end_time durch eine ganze Zahl dargestellt werden, um die Zeit anzugeben. Das bedeutet, dass für ein Meeting von (0, 10) und (10, 20) nur 1 Meetingraum benötigt wird, da sich die Meetings nicht überlappen.

#### Start

> Hier ist ein Beispiel und etwas Startcode:

```py
def meeting_rooms(meetings):
# Das ausfüllen

print(meeting_rooms([(0, 10), (10, 20)]))
# 1 (Überlappen sich nicht, da sie nacheinander sind)

print(meeting_rooms([(20, 30), (10, 21), (0, 50)]))
# 3 (Alle Meetings überschneiden sich zum Zeitpunkt 20)
```

#### Hinweis

> -> Die Zeiten müssen nicht Real sein. Sprich es kann auch von 0 Uhr bis 100 Uhr gehen.<br>
> -> Die Tuples sind nicht nach den Startzeiten / Endzeiten sortiert<br>
