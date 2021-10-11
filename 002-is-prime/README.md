# is prime

Write a function, isPrime, that takes in a number as an argument. The function should return a boolean indicating whether or not the given number is prime.

A prime number is a number that is only divisible two distinct numbers: 1 and itself.

For example, 7 is a prime because it is only divisible by 1 and 7. For example, 6 is not a prime because it is divisible by 1, 2, 3, and 6.

You can assume that the input number is a positive integer.

test_00:
```js
isPrime(2); // -> true
```

test_01:
```js
isPrime(3); // -> true
```

test_02:
```js
isPrime(4); // -> false
```

test_03:
```js
isPrime(5); // -> true
```

test_04:
```js
isPrime(6); // -> false
```

test_05:
```js
isPrime(7); // -> true
```

test_06:
```js
isPrime(8); // -> false
```

test_07:
```js
isPrime(25); // -> false
```

test_08:
```js
isPrime(31); // -> true
```

test_09:
```js
isPrime(2017); // -> true
```

test_10:
```js
isPrime(2048); // -> false
```

test_11:
```js
isPrime(1); // -> false
```