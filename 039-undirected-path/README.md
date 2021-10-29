# undirected path

Write a function, undirectedPath, that takes in an array of edges for an undirected graph and two nodes (nodeA, nodeB). The function should return a boolean indicating whether or not there exists a path between nodeA and nodeB.

test_00:
```js
const edges = [
['i', 'j'],
['k', 'i'],
['m', 'k'],
['k', 'l'],
['o', 'n']
];

undirectedPath(edges, 'j', 'm'); // -> true
```

test_01:
```js
const edges = [
['i', 'j'],
['k', 'i'],
['m', 'k'],
['k', 'l'],
['o', 'n']
];

undirectedPath(edges, 'm', 'j'); // -> true
```

test_02:
```js
const edges = [
['i', 'j'],
['k', 'i'],
['m', 'k'],
['k', 'l'],
['o', 'n']
];

undirectedPath(edges, 'l', 'j'); // -> true
```

test_03:
```js
const edges = [
['i', 'j'],
['k', 'i'],
['m', 'k'],
['k', 'l'],
['o', 'n']
];

undirectedPath(edges, 'k', 'o'); // -> false
```

test_04:
```js
const edges = [
['i', 'j'],
['k', 'i'],
['m', 'k'],
['k', 'l'],
['o', 'n']
];

undirectedPath(edges, 'i', 'o'); // -> false
```

test_05:
```js
const edges = [
['b', 'a'],
['c', 'a'],
['b', 'c'],
['q', 'r'],
['q', 's'],
['q', 'u'],
['q', 't'],
];


undirectedPath(edges, 'a', 'b'); // -> true
```

test_06:
```js
const edges = [
['b', 'a'],
['c', 'a'],
['b', 'c'],
['q', 'r'],
['q', 's'],
['q', 'u'],
['q', 't'],
];

undirectedPath(edges, 'a', 'c'); // -> true
```

test_07:
```js
const edges = [
['b', 'a'],
['c', 'a'],
['b', 'c'],
['q', 'r'],
['q', 's'],
['q', 'u'],
['q', 't'],
];

undirectedPath(edges, 'r', 't'); // -> true
```

test_08:
```js
const edges = [
['b', 'a'],
['c', 'a'],
['b', 'c'],
['q', 'r'],
['q', 's'],
['q', 'u'],
['q', 't'],
];

undirectedPath(edges, 'r', 'b'); // -> false
```
