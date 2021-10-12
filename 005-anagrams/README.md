# anagrams

Write a function, anagrams, that takes in two strings as arguments. The function should return a boolean indicating whether or not the strings are anagrams. Anagrams are strings that contain the same characters, but in any order.

test_00:
```js
anagrams('restful', 'fluster'); // -> true
```

test_01:
```js
anagrams('cats', 'tocs'); // -> false
```

test_02:
```js
anagrams('monkeyswrite', 'newyorktimes'); // -> true
```

test_03:
```js
anagrams('paper', 'reapa'); // -> false
```

test_04:
```js
anagrams('elbow', 'below'); // -> true
```

test_05:
```js
anagrams('tax', 'taxi'); // -> false
```
test_06:
```js
anagrams('taxi', 'tax'); // -> false
```

test_07:
```js
anagrams('night', 'thing'); // -> true
```

test_08:
```js
anagrams('abbc', 'aabc'); // -> false
```
