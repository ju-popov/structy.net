# paired parentheses

Write a function, pairedParentheses, that takes in a string as an argument. The function should return a boolean indicating whether or not the string has well-formed parentheses.

You may assume the string contains only alphabetic characters, '(', or ')'.

test_00:
```js
pairedParentheses("(david)((abby))"); // -> true
```

test_01:
```js
pairedParentheses("()rose(jeff"); // -> false
```

test_02:
```js
pairedParentheses(")("); // -> false
```

test_03:
```js
pairedParentheses("()"); // -> true
```

test_04:
```js
pairedParentheses("(((potato())))"); // -> true
```

test_05:
```js
pairedParentheses("(())(water)()"); // -> true
```

test_06:
```js
pairedParentheses("(())(water()()"); // -> false
```

test_07:
```js
pairedParentheses(""); // -> true
```
