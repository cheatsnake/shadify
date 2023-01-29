# Anagram

<p align="center"><img src="../images/anagram.png" alt="Anagram"/></p>

[**Anagrams**](https://en.wikipedia.org/wiki/Anagrams_(game)) is a whole type of puzzles associated with composing all possible words from a given set of letters. This module implements the simplest variation of anagrams: a random word is given, the letters of which should be used to compose as many other words as possible. 

> Only English nouns are used for composition, a list of which can be found in this repository at the [`nouns.txt`](https://github.com/cheatsnake/shadify/blob/master/data/nouns.txt).

## Generate random anagram

```nginx
GET https://shadify.dev/api/anagram/generator
```

Returned response:

```json
{
    "task":"possibility",
    "words":[
        "bit","bolt","boy","lip","list","loss","oil",
        "pilot","plot","toy","pot","slip","soil","soy",
        "spot","spy","stop","tip","top"
    ]
}
```

> - `task` - a word from which you need to make other words. <br> 
> - `words` - an array of all possible words that are compiled from `task`. 
