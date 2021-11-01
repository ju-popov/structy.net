# has cycle

Write a function, hasCycle, that takes in an object representing the adjacency list of a directed graph. The function should return a boolean indicating whether or not the graph contains a cycle.

test_00:
```js
hasCycle({
  a: ["b"],
  b: ["c"],
  c: ["a"],
}); // -> true
```

test_01:
```js
hasCycle({
  a: ["b", "c"],
  b: ["c"],
  c: ["d"],
  d: [],
}); // -> false
```

test_02:
```js
hasCycle({
  a: ["b", "c"],
  b: [],
  c: [],
  e: ["f"],
  f: ["e"],
}); // -> true
```

test_03:
```js
hasCycle({
  q: ["r", "s"],
  r: ["t", "u"],
  s: [],
  t: [],
  u: [],
  v: ["w"],
  w: [],
  x: ["w"],
}); // -> false
```

test_04:
```js
hasCycle({
  a: ["b"],
  b: ["c"],
  c: ["a"],
  g: [],
}); // -> true
```
