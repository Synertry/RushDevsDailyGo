### NR 26

---

#### Aufgabe
> Bei einer gegebenen sortierten Liste der Größe n mit m eindeutigen Elementen (also m < n), modifiziere die Liste so, dass die ersten m eindeutigen Elemente in der Liste sortiert werden, wobei der Rest der Liste ignoriert wird. Deine Lösung sollte eine Raumkomplexität von O(1) haben. Anstatt die Liste zurückzugeben (da Sie nur die ursprüngliche Liste ändern), solltest du zurückgeben, was m ist.


#### Start
> Hier ist ein Beispiel und etwas Startcode:

```py
def remove_dups(nums):
# Das Ausfüllen

nums = [1, 1, 2, 3, 4, 4, 4, 4, 4, 5, 5, 6, 7, 9]
print(remove_dups(nums))
# 8
print(nums)
# [1, 2, 3, 4, 5, 6, 7, 9]

nums = [1, 1, 1, 1, 1, 1]
print(remove_dups(nums))
print(nums)
# 1
# [1]
```


#### Hinweis
> -> nums.length > 1<br>
> -> nums.entries = Integer<br>
> -> m = nums.length<br>
> -> Space-Complexity is O(1)<br>