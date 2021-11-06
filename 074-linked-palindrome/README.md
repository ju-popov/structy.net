# linked palindrome

Write a function, linkedPalindrome, that takes in the head of a linked list as an argument. The function should return a boolean indicating whether or not the linked list is a palindrome. A palindrome is a sequence that is the same both forwards and backwards.

test_00:
```js
const a = new Node(3);
const b = new Node(2);
const c = new Node(7);
const d = new Node(7);
const e = new Node(2);
const f = new Node(3);

a.next = b;
b.next = c;
c.next = d;
d.next = e;
e.next = f;

// 3 -> 2 -> 7 -> 7 -> 2 -> 3
linkedPalindrome(a); // true
```

test_01:
```js
const a = new Node(3);
const b = new Node(2);
const c = new Node(4);

a.next = b;
b.next = c;

// 3 -> 2 -> 4
linkedPalindrome(a); // false
```

test_02:
```js
const a = new Node(3);
const b = new Node(2);
const c = new Node(3);

a.next = b;
b.next = c;

// 3 -> 2 -> 3
linkedPalindrome(a); // true
```

test_03:
```js
const a = new Node(0);
const b = new Node(1);
const c = new Node(0);
const d = new Node(1);
const e = new Node(0);

a.next = b;
b.next = c;
c.next = d;
d.next = e;

// 0 -> 1 -> 0 -> 1 -> 0
linkedPalindrome(a); // true
```

test_04:
```js
const a = new Node(0);
const b = new Node(1);
const c = new Node(0);
const d = new Node(1);
const e = new Node(1);

a.next = b;
b.next = c;
c.next = d;
d.next = e;

// 0 -> 1 -> 0 -> 1 -> 1
linkedPalindrome(a); // false
```

test_05:
```js
const a = new Node(5);

// 5
linkedPalindrome(a); // true
```

test_06:
```js
linkedPalindrome(null); // true
```
