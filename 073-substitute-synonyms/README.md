# substituting synonyms

Write a function, substitutingSynonyms, that takes in a sentence and an object as arguments. The object contains words as keys whose values are arrays containing synonyms. The function should return an array containing all possible sentences that can be formed by substituting words of the sentence with their synonyms.

You may return the possible sentences in any order, as long as you return all of them.

test_00:
```js
const sentence = "follow the yellow brick road";
const synonyms = {
  follow: ["chase", "pursue"],
  yellow: ["gold", "amber", "lemon"],
};

substituteSynonyms(sentence, synonyms);
// [
//   'chase the gold brick road',
//   'chase the amber brick road',
//   'chase the lemon brick road',
//   'pursue the gold brick road',
//   'pursue the amber brick road',
//   'pursue the lemon brick road'
// ]
```

test_01:
```js
const sentence = "I think it's gonna be a long long time";
const synonyms = {
  think: ["believe", "reckon"],
  long: ["lengthy", "prolonged"],
};

substituteSynonyms(sentence, synonyms);
// [
//   "I believe it's gonna be a lengthy lengthy time",
//   "I believe it's gonna be a lengthy prolonged time",
//   "I believe it's gonna be a prolonged lengthy time",
//   "I believe it's gonna be a prolonged prolonged time",
//   "I reckon it's gonna be a lengthy lengthy time",
//   "I reckon it's gonna be a lengthy prolonged time",
//   "I reckon it's gonna be a prolonged lengthy time",
//   "I reckon it's gonna be a prolonged prolonged time"
// ]
```

test_02:
```js
const sentence = "palms sweaty knees weak arms heavy";
const synonyms = {
  palms: ["hands", "fists"],
  heavy: ["weighty", "hefty", "burdensome"],
  weak: ["fragile", "feeble", "frail", "sickly"],
};

substituteSynonyms(sentence, synonyms);
// [
//   'hands sweaty knees fragile arms weighty',
//   'hands sweaty knees fragile arms hefty',
//   'hands sweaty knees fragile arms burdensome',
//   'hands sweaty knees feeble arms weighty',
//   'hands sweaty knees feeble arms hefty',
//   'hands sweaty knees feeble arms burdensome',
//   'hands sweaty knees frail arms weighty',
//   'hands sweaty knees frail arms hefty',
//   'hands sweaty knees frail arms burdensome',
//   'hands sweaty knees sickly arms weighty',
//   'hands sweaty knees sickly arms hefty',
//   'hands sweaty knees sickly arms burdensome',
//   'fists sweaty knees fragile arms weighty',
//   'fists sweaty knees fragile arms hefty',
//   'fists sweaty knees fragile arms burdensome',
//   'fists sweaty knees feeble arms weighty',
//   'fists sweaty knees feeble arms hefty',
//   'fists sweaty knees feeble arms burdensome',
//   'fists sweaty knees frail arms weighty',
//   'fists sweaty knees frail arms hefty',
//   'fists sweaty knees frail arms burdensome',
//   'fists sweaty knees sickly arms weighty',
//   'fists sweaty knees sickly arms hefty',
//   'fists sweaty knees sickly arms burdensome'
// ]
```
