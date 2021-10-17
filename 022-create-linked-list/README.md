# create linked list

Write a function, createLinkedList, that takes in an array of values as an argument. The function should create a linked list containing each element of the array as the values of the nodes. The function should return the head of the linked list.

test_00:
```js
createLinkedList(["h", "e", "y"]);
// h -> e -> y
```

test_01:
```js
createLinkedList([1, 7, 1, 8]);
// 1 -> 7 -> 1 -> 8
```

test_02:
```js
createLinkedList(["a"]);
// a
```

test_03:
```js
createLinkedList([]);
// null
```