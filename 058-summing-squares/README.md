# summing squares

Write a function, summingSquares, that takes a target number as an argument. The function should return the minimum number of perfect squares that sum to the target. A perfect square is a number of the form (i*i) where i >= 1.

For example: 1, 4, 9, 16 are perfect squares, but 8 is not perfect square.

```
Given 12:

summingSquares(12) -> 3

The minimum squares required for 12 is three, by doing 4 + 4 + 4.

Another way to make 12 is 9 + 1 + 1 + 1, but that requires four perfect squares.
```

test_00:
```js
summingSquares(8); // -> 2
```

test_01:
```js
summingSquares(9); // -> 1
```

test_02:
```js
summingSquares(12); // -> 3
```

test_03:
```js
summingSquares(1); // -> 1
```

test_04:
```js
summingSquares(31); // -> 4
```

test_05:
```js
summingSquares(50); // -> 2
```

test_06:
```js
summingSquares(68); // -> 2
```

test_07:
```js
summingSquares(87); // -> 4
```
