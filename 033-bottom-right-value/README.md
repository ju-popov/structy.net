# bottom right value

Write a function, bottomRightValue, that takes in the root of a binary tree. The function should return the right-most value in the bottom-most level of the tree.

You may assume that the input tree is non-empty.

test_00:
```js
const a = new Node(3);
const b = new Node(11);
const c = new Node(10);
const d = new Node(4);
const e = new Node(-2);
const f = new Node(1);

a.left = b;
a.right = c;
b.left = d;
b.right = e;
c.right = f;

//       3
//    /    \
//   11     10
//  / \      \
// 4   -2     1

bottomRightValue(a); // -> 1
```

test_01:
```js
const a = new Node(-1);
const b = new Node(-6);
const c = new Node(-5);
const d = new Node(-3);
const e = new Node(-4);
const f = new Node(-13);
const g = new Node(-2);
const h = new Node(6);

a.left = b;
a.right = c;
b.left = d;
b.right = e;
c.right = f;
e.left = g;
e.right = h;

//        -1
//      /   \
//    -6    -5
//   /  \     \
// -3   -4   -13
//     / \       
//    -2  6

bottomRightValue(a); // -> 6
```

test_02:
```js
const a = new Node(-1);
const b = new Node(-6);
const c = new Node(-5);
const d = new Node(-3);
const e = new Node(-4);
const f = new Node(-13);
const g = new Node(-2);
const h = new Node(6);
const i = new Node(7);


a.left = b;
a.right = c;
b.left = d;
b.right = e;
c.right = f;
e.left = g;
e.right = h;
f.left = i;

//        -1
//      /   \
//    -6    -5
//   /  \     \
// -3   -4   -13
//     / \    /   
//    -2  6  7

bottomRightValue(a); // -> 7
```

test_03
```js
const a = new Node('a');
const b = new Node('b');
const c = new Node('c');
const d = new Node('d');
const e = new Node('e');
const f = new Node('f');

a.left = b;
a.right = c;
b.right = d;
d.left = e;
e.right = f;

//      a
//    /   \
//   b     c
//    \
//     d
//    /
//   e
//  /
// f

bottomRightValue(a); // -> 'f'
```

test_04
```js
const a = new Node(42);

//      42

bottomRightValue(a); // -> 42
```
