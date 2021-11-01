# longest path

Write a function, longestPath, that takes in an adjacency list for a directed acyclic graph. The function should return the length of the longest path within the graph. A path may start and end at any two nodes. The length of a path is considered the number of edges in the path, not the number of nodes.

test_00:
```js
const graph = {
  a: ['c', 'b'],
  b: ['c'],
  c: []
};

longestPath(graph); // -> 2
```

test_01:
```js
const graph = {
  a: ['c', 'b'],
  b: ['c'],
  c: [],
  q: ['r'],
  r: ['s', 'u', 't'],
  s: ['t'],
  t: ['u'],
  u: []
};

longestPath(graph); // -> 4
```

test_02:
```js
const graph = {
  h: ['i', 'j', 'k'],
  g: ['h'],
  i: [],
  j: [],
  k: [],
  x: ['y'],
  y: []
};

longestPath(graph); // -> 2
```

test_03:
```js
const graph = {
  a: ['b'],
  b: ['c'],
  c: [],
  e: ['f'],
  f: ['g'],
  g: ['h'],
  h: []
};

longestPath(graph); // -> 3
```
