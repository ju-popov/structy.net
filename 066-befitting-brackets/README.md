# befitting brackets

Write a function, befittingBrackets, that takes in a string as an argument. The function should return a boolean indicating whether or not the string contains correctly matched brackets.

You may assume the string contains only characters: ( ) [ ] { }

test_00:
```js
befittingBrackets('(){}[](())'); // -> true
```

test_01:
```js
befittingBrackets('({[]})'); // -> true
```

test_02:
```js
befittingBrackets('[][}'); // -> false
```

test_03:
```js
befittingBrackets('{[]}({}'); // -> false
```

test_04:
```js
befittingBrackets('[]{}(}[]'); // -> false
```

test_05:
```js
befittingBrackets('[]{}()[]'); // -> true
```

test_06:
```js
befittingBrackets(']{}'); // -> false
```

test_07:
```js
befittingBrackets(''); // -> true
```
