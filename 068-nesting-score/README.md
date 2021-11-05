# nesting score

Write a function, nestingScore, that takes in a string of brackets as an argument. The function should return the score of the string according to the following rules:

[] is worth 1 point
XY is worth m + n points where X, Y are substrings of well-formed brackets and m, n are their respective scores
[S] is worth 2 * k points where S is a substring of well-formed brackets and k is the score of that substring
You may assume that the input only contains well-formed square brackets.

test_00:
```js
nestingScore("[]"); // -> 1
```

test_01:
```js
nestingScore("[][][]"); // -> 3
```

test_02:
```js
nestingScore("[[]]"); // -> 2
```

test_03:
```js
nestingScore("[[][]]"); // -> 4
```

test_04:
```js
nestingScore("[[][][]]"); // -> 6
```

test_05:
```js
nestingScore("[[][]][]"); // -> 5
```

test_06:
```js
nestingScore("[][[][]][[]]"); // -> 7
```

test_07:
```js
nestingScore("[[[[[[[][]]]]]]][]"); // -> 129
```
