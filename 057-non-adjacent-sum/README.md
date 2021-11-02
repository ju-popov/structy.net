# non adjacent sum

Write a function, nonAdjacentSum, that takes in an array of numbers as an argument. The function should return the maximum sum of non-adjacent elements in the array. There is no limit on how many elements can be taken into the sum as long as they are not adjacent.

```
For example, given:

[2, 4, 5, 12, 7]

The maximum non-adjacent sum is 16, because 4 + 12. 
4 and 12 are not adjacent in the array.
```

test_00:
```js
const nums = [2, 4, 5, 12, 7];
nonAdjacentSum(nums); // -> 16
```

test_01:
```js
const nums = [7, 5, 5, 12];
nonAdjacentSum(nums); // -> 19
```

test_02:
```js
const nums = [7, 5, 5, 12, 17, 29];
nonAdjacentSum(nums); // -> 48
```

test_03:
```js
const nums = [
  72, 62, 10,  6, 20, 19, 42,
  46, 24, 78, 30, 41, 75, 38,
  23, 28, 66, 55, 12, 17, 9,
  12, 3, 1, 19, 30, 50, 20
];
nonAdjacentSum(nums); // -> 488
```

test_04:
```js
const nums = [
  72, 62, 10,  6, 20, 19, 42, 46, 24, 78,
  30, 41, 75, 38, 23, 28, 66, 55, 12, 17,
  83, 80, 56, 68,  6, 22, 56, 96, 77, 98,
  61, 20,  0, 76, 53, 74,  8, 22, 92, 37,
  30, 41, 75, 38, 23, 28, 66, 55, 12, 17,
  72, 62, 10,  6, 20, 19, 42, 46, 24, 78,
  42
];
nonAdjacentSum(nums); // -> 1465
```
