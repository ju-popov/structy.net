# compress

Write a function, compress, that takes in a string as an argument. The function should return a compressed version of the string where consecutive occurences of the same characters are compressed into the number of occurences followed by the character. Single character occurences should not be changed.

```
'aaa' compresses to '3a'
'cc' compresses to '2c'
't' should remain as 't'
```

You can assume that the input only contains alphabetic characters.

test_00:
```js
compress('ccaaatsss'); // -> '2c3at3s'
```

test_01:
```js
compress('ssssbbz'); // -> '4s2bz'
```

test_02:
```js
compress('ppoppppp'); // -> '2po5p'
```

test_03:
```js
compress('nnneeeeeeeeeeeezz'); // -> '3n12e2z'
```

test_04:
```js
compress('yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy');
// -> '127y'
```
