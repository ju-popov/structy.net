# min change

Write a function minChange that takes in an amount and an array of coins. The function should return the minimum number of coins required to create the amount. You may use each coin as many times as necessary.

If it is not possible to create the amount, then return -1.

test_00:
```js
minChange(8, [1, 5, 4, 12]); // -> 2, because 4+4 is the minimum coins possible
```

test_01:
```js
minChange(13, [1, 9, 5, 14, 30]); // -> 5
```

test_02:
```js
minChange(23, [2, 5, 7]); // -> 4
```

test_03:
```js
minChange(102, [1, 5, 10, 25]); // -> 6
```

test_04:
```js
minChange(200, [1, 5, 10, 25]); // -> 8
```

test_05:
```js
minChange(2017, [4, 2, 10]); // -> -1
```

test_06:
```js
minChange(271, [10, 8, 265, 24]); // -> -1
```

test_07:
```js
minChange(0, [4, 2, 10]); // -> 0
```

test_08:
```js
minChange(0, []); // -> 0
```
