# overlap subsequence

Write a function, overlapSubsequence, that takes in two strings as arguments. The function should return the length of the longest overlapping subsequence.

A subsequence of a string can be created by deleting any characters of the string, while maintaining the relative order of characters.

test_00:
```js
overlapSubsequence("dogs", "daogt"); // -> 3
```

test_01:
```js
overlapSubsequence("xcyats", "criaotsi"); // -> 4
```

test_02:
```js
overlapSubsequence("xfeqortsver", "feeeuavoeqr"); // -> 5
```

test_03:
```js
overlapSubsequence("kinfolklivemustache", "bespokekinfolksnackwave"); // -> 11
```

test_04:
```js
overlapSubsequence(
  "mumblecorebeardleggingsauthenticunicorn",
  "succulentspughumblemeditationlocavore"
); // -> 15
```
