# connected components count

Write a function, connectedComponentsCount, that takes in the adjacency list of an undirected graph. The function should return the number of connected components within the graph.

test_00:
```js
connectedComponentsCount({
0: [8, 1, 5],
1: [0],
5: [0, 8],
8: [0, 5],
2: [3, 4],
3: [2, 4],
4: [3, 2]
}); // -> 2
```

test_01:
```js
connectedComponentsCount({
1: [2],
2: [1,8],
6: [7],
9: [8],
7: [6, 8],
8: [9, 7, 2]
}); // -> 1
```

test_02:
```js
connectedComponentsCount({
3: [],
4: [6],
6: [4, 5, 7, 8],
8: [6],
7: [6],
5: [6],
1: [2],
2: [1]
}); // -> 3
```

test_03:
```js
connectedComponentsCount({}); // -> 0
```
`
test_04:
```js
connectedComponentsCount({
0: [4,7],
1: [],
2: [],
3: [6],
4: [0],
6: [3],
7: [0],
8: []
}); // -> 5
```
