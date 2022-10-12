### NR 27

---

#### Aufgabe
> Ein Krimineller erstellt eine Lösegeldforderung. Um seine Handschrift zu verschleiern, schneidet er Buchstaben aus einer Zeitschrift aus.
> 
> Bestimme anhand eines Magazins mit Briefen und der Notiz, die er schreiben möchte, ob er das Wort bilden kann.


#### Start
> Hier ist ein Beispiel und etwas Startcode:

```
class Solution(object):
def canSpell(self, magazine, note):
# Das Ausfüllen

print(Solution().canSpell(['a', 'b', 'c', 'd', 'e', 'f'], 'bed'))
# True

print(Solution().canSpell(['a', 'b', 'c', 'd', 'e', 'f'], 'cat'))
# False
```


#### Hinweis
> -> typeof magazine = Array<String><br>
> -> typeof note = String<br>
> -> note.length > 1<br>
> -> output: Boolean<br>