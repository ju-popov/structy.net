# linked list cycle

Write a function, linkedListCycle, that takes in the head of a linked list as an argument. The function should return a boolean indicating whether or not the linked list contains a cycle.

test_00:
```js
const a = new Node('a');
const b = new Node('b');
const c = new Node('c');
const d = new Node('d');

a.next = b;
b.next = c;
c.next = d;
d.next = b; // cycle

//         _______
//       /        \
// a -> b -> c -> d 

linkedListCycle(a);  // true
```

test_01:
```js
const q = new Node('q');
const r = new Node('r');
const s = new Node('s');
const t = new Node('t');
const u = new Node('u');

q.next = r;
r.next = s;
s.next = t;
t.next = u;
u.next = q; // cycle

//    ________________
//  /                 \
// q -> r -> s -> t -> u 

linkedListCycle(q);  // true
```

test_02
```js
const a = new Node('a');
const b = new Node('b');
const c = new Node('c');
const d = new Node('d');

a.next = b;
b.next = c;
c.next = d;

// a -> b -> c -> d 

linkedListCycle(a);  // false
```

test_03:
```js
const q = new Node('q');
const r = new Node('r');
const s = new Node('s');
const t = new Node('t');
const u = new Node('u');

q.next = r;
r.next = s;
s.next = t;
t.next = u;
u.next = t; // cycle

//                   __
//                 /   \
// q -> r -> s -> t -> u 

linkedListCycle(q);  // true
```

test_04:
```js
const p = new Node('p');

// p

linkedListCycle(p); // false
```

test_05:
```js
linkedListCycle(null); // false
```
