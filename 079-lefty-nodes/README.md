# lefty nodes

Write a function, leftyNodes, that takes in the root of a binary tree. The function should return an array containing the left-most value on every level of the tree. The array must be ordered in a top-down fashion where the root is the first element.

test_00
```js
const a = new Node('a');
const b = new Node('b');
const c = new Node('c');
const d = new Node('d');
const e = new Node('e');
const f = new Node('f');
const g = new Node('g');
const h = new Node('h');

a.left = b;
a.right = c;
b.left = d;
b.right = e;
c.right = f;
e.left = g;
e.right = h;

//      a
//    /    \
//   b      c
//  / \      \
// d   e      f
//    / \
//    g  h

leftyNodes(a); // [ 'a', 'b', 'd', 'g' ]
```

test_01
```js
const u = new Node('u');
const t = new Node('t');
const s = new Node('s');
const r = new Node('r');
const q = new Node('q');
const p = new Node('p');

u.left = t;
u.right = s;
s.right = r;
r.left = q;
r.right = p;

//     u
//  /    \
// t      s
//         \
//         r
//        / \
//        q  p

leftyNodes(u); // [ 'u', 't', 'r', 'q' ]
```

test_02
```js
const l = new Node('l');
const m = new Node('m');
const n = new Node('n');
const o = new Node('o');
const p = new Node('p');
const q = new Node('q');
const r = new Node('r');
const s = new Node('s');
const t = new Node('t');

l.left = m;
l.right = n;
n.left = o;
n.right = p;
o.left = q;
o.right = r;
p.left = s;
p.right = t;

//        l
//     /     \
//    m       n
//         /    \
//         o     p
//        / \   / \
//       q   r s   t

leftyNodes(l); // [ 'l', 'm', 'o', 'q' ]
```

test_03
```js
const n = new Node('n');
const y = new Node('y');
const c = new Node('c');

n.left = y;
n.right = c;

//       n
//     /   \
//    y     c

leftyNodes(n); // [ 'n', 'y' ]
```

test_04
```js
const i = new Node('i');
const n = new Node('n');
const s = new Node('s');
const t = new Node('t');

i.right = n;
n.left = s;
n.right = t;

//       i
//        \
//         n
//        / \
//       s   t

leftyNodes(i); // [ 'i', 'n', 's' ]
```

test_05
```js
leftyNodes(null); // [ ]
```
