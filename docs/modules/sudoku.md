# Sudoku

<p align="center"><img src="../images/sudoku.png" alt="Sudoku"/></p>

[**Sudoku**](https://en.wikipedia.org/wiki/Sudoku) is a popular number puzzle game. The essence of the game is to fill free cells with numbers from 1 to 9 so that in each row, each column and each small 3×3 square each digit occurs only once.

## Generating random Sudoku

```nginx
GET https://shadify.dev/api/sudoku/generator
```

| Parameter | Description                                                                                                                              |
| --------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `fill`    | _Optional_ <br><br> A number from 0 to 100 that corresponds to the field fill level (as a percentage). <br><br> The default value is 30. |

Returned response:

```json
{
    "grid": [
        [9, 8, 7, 3, 2, 1, 5, 6, 4],
        [6, 5, 4, 9, 8, 7, 2, 3, 1],
        [3, 2, 1, 6, 5, 4, 8, 9, 7],
        [7, 6, 5, 1, 9, 8, 3, 4, 2],
        [1, 9, 8, 4, 3, 2, 6, 7, 5],
        [4, 3, 2, 7, 6, 5, 9, 1, 8],
        [8, 7, 6, 2, 1, 9, 4, 5, 3],
        [5, 4, 3, 8, 7, 6, 1, 2, 9],
        [2, 1, 9, 5, 4, 3, 7, 8, 6]
    ],
    "task": [
        [0, 8, 7, 0, 0, 0, 0, 0, 4],
        [0, 0, 0, 9, 0, 0, 2, 0, 0],
        [0, 2, 1, 6, 5, 4, 0, 0, 7],
        [0, 0, 0, 1, 0, 8, 0, 4, 2],
        [0, 9, 8, 4, 3, 0, 0, 0, 0],
        [4, 0, 2, 0, 0, 0, 0, 0, 0],
        [8, 0, 6, 2, 0, 0, 0, 0, 0],
        [5, 0, 3, 0, 0, 6, 0, 2, 9],
        [0, 0, 0, 0, 4, 0, 0, 0, 0]
    ]
}
```

> -   `grid` - fully solved Sudoku. <br>
> -   `task` - the task, which must be solved by filling in the zeros with the correct numbers.

## Checking Sudoku

To check Sudoku you can use GET-request with required query-parameter

```nginx
GET https://shadify.dev/api/sudoku/verifier
```

| Parameter | Description                                                                                                                                  |
| --------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `task`    | _Required_ <br><br> A string of the form `123-123-123`, which corresponds to the Sudoku field, where each line (row) is separated by a dash. |

An example of a GET-request:

```nginx
https://shadify.dev/api/sudoku/verifier?task=123564897-456897231-789231564-
234675918-567918342-891342675-345786129-678129453-912453786
```

Also, for convenience, you can use POST-request with body in JSON format

```nginx
POST https://shadify.dev/api/sudoku/verifier
```

Example body for a POST request:

```json
{
    "task": [
        [9, 6, 3, 2, 8, 5, 7, 4, 1],
        [8, 5, 2, 1, 7, 4, 6, 3, 9],
        [1, 7, 4, 3, 9, 6, 8, 5, 2],
        [6, 3, 9, 8, 5, 2, 3, 1, 7],
        [5, 2, 8, 7, 4, 1, 3, 9, 6],
        [7, 4, 1, 9, 6, 3, 5, 2, 8],
        [4, 1, 7, 6, 3, 9, 2, 8, 5],
        [3, 9, 6, 5, 2, 8, 1, 7, 4],
        [2, 8, 5, 4, 1, 7, 9, 6, 3]
    ]
}
```

Returned response:

```json
{
    "isValid": false,
    "position": "row-4"
}
```
