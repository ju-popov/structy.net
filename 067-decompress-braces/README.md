# decompress braces

Write a function, decompressBraces, that takes in a compressed string as an argument. The function should return the string decompressed.

The compression format of the input string is 'n{subString}', where the subString within braces should be repeated n times.

You may assume that every number n is guaranteed to be an integer between 1 through 9.

You may assume that the input is valid and the decompressed string will only contain alphabetic characters.

test_00:
```js
decompressBraces("2{q}3{tu}v"); 
// -> qqtututuv 
```

test_01:
```js
decompressBraces("ch3{ao}"); 
// -> chaoaoao
```

test_02:
```js
decompressBraces("2{y3{o}}s"); 
// -> yoooyooos
```

test_03:
```js
decompressBraces("z3{a2{xy}b}"); 
// -> zaxyxybaxyxybaxyxyb 
```

test_04:
```js
decompressBraces("2{3{r4{e}r}io}"); 
// -> reeeerreeeerreeeerioreeeerreeeerreeeerio 
```

test_05:
```js
decompressBraces("go3{spinn2{ing}s}"); 
// -> gospinningingsspinningingsspinningings 
```

test_06:
```js
decompressBraces("2{l2{if}azu}l"); 
// -> lififazulififazul 
```

test_07:
```js
decompressBraces("3{al4{ec}2{icia}}"); 
// -> alececececiciaiciaalececececiciaiciaalececececiciaicia 
```
