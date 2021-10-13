# linked list values

Write a function, linkedListValues, that takes in the head of a linked list as an argument. The function should return an array containing all values of the nodes in the linked list.

test_00:
```js
const a = new Node("a");
const b = new Node("b");
const c = new Node("c");
const d = new Node("d");

a.next = b;
b.next = c;
c.next = d;

// a -> b -> c -> d

linkedListValues(a); // -> [ 'a', 'b', 'c', 'd' ]
```

test_01:
```js
const x = new Node("x");
const y = new Node("y");

x.next = y;

// x -> y

linkedListValues(x); // -> [ 'x', 'y' ]
```

test_02:
```js
const q = new Node("q");

// q

linkedListValues(q); // -> [ 'q' ]
```

test_03:
```js
linkedListValues(null); // -> [ ]
```