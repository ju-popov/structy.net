# semesters required

Write a function, semestersRequired, that takes in a number of courses (n) and a list of prerequisites as arguments. Courses have ids ranging from 0 through n - 1. A single prerequisite of [A, B] means that course A must be taken before course B. Return the minimum number of semesters required to complete all n courses. There is no limit on how many courses you can take in a single semester, as long the prerequisites of a course are satisfied before taking it.

Note that given prerequisite [A, B], you cannot take course A and course B concurrently in the same semester. You must take A in some semester before B.

You can assume that it is possible to eventually complete all courses.

test_00:
```js
const numCourses = 6;
const prereqs = [
  [1, 2],
  [2, 4],
  [3, 5],
  [0, 5],
];
semestersRequired(numCourses, prereqs); // -> 3
```

test_01:
```js
const numCourses = 7;
const prereqs = [
  [4, 3],
  [3, 2],
  [2, 1],
  [1, 0],
  [5, 2],
  [5, 6],
];
semestersRequired(numCourses, prereqs); // -> 5
```

test_02:
```js
const numCourses = 5;
const prereqs = [
  [1, 0],
  [3, 4],
  [1, 2],
  [3, 2],
];
semestersRequired(numCourses, prereqs); // -> 2
```

test_03:
```js
const numCourses = 12;
const prereqs = [];
semestersRequired(numCourses, prereqs); // -> 1
```

test_04:
```js
const numCourses = 3;
const prereqs = [
  [0, 2],
  [0, 1],
  [1, 2],
];
semestersRequired(numCourses, prereqs); // -> 3
```

test_05:
```js
const numCourses = 6;
const prereqs = [
  [3, 4],
  [3, 0],
  [3, 1],
  [3, 2],
  [3, 5],
];
semestersRequired(numCourses, prereqs); // -> 2
```
