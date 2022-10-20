### NR 21

---

#### Aufgabe

> Finde alle Wörter, die Verkettungen einer Liste sind (Wörter, die aus anderen Wörtern in der Liste gemacht wurden, es kann aus input.length - 1 Wörtern bestehen).

#### Input:

> ["tech", "lead", "techlead", "cat", "cats", "dog", "catsdog"]

#### Output:

> ['techlead', 'catsdog']

#### Start

Hier ist ein Beispiel und etwas Startcode:

```py
class Solution(object):
def findAllConcatenatedWordsInADict(self, dict):
# Das Ausfüllen

input = ["tech", "lead", "techlead", "cat", "cats", "dog", "catsdog"]

print(Solution().findAllConcatenatedWordsInADict(input))
```

#### Hinweis

> -> typeof Solution = class<br>
> -> dict.length > 3<br>
> -> dict.entries = String<br>
