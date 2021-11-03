# can concat

Write a function, canConcat, that takes in a string and an array of words as arguments. The function should return boolean indicating whether or not it is possible to concatenate words of the array together to form the string.

You may reuse words of the array as many times as needed.

test_00:
```js
canConcat("oneisnone", ["one", "none", "is"]); // -> true
```

test_01:
```js
canConcat("oneisnone", ["on", "e", "is"]); // -> false
```

test_02:
```js
canConcat("oneisnone", ["on", "e", "is", "n"]); // -> true
```

test_03:
```js
canConcat("foodisgood", ["is", "g", "ood", "f"]); // -> true
```

test_04:
```js
canConcat("santahat", ["santah", "hat"]); // -> false
```

test_05:
```js
canConcat("santahat", ["santah", "san", "hat", "tahat"]); // -> true
```

test_06:
```js
canConcat("rrrrrrrrrrrrrrrrrrrrrrrrrrx", ["r", "rr", "rrr", "rrrr", "rrrrr", "rrrrrr"]); // -> false
```
