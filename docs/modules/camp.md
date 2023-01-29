# Camp

<p align="center"><img src="../images/camp.png" alt="Camp"/></p>

[Camp (aka Tents and trees, Tents)](https://www.chiark.greenend.org.uk/~sgtatham/puzzles/js/tents.html) is a logic puzzle with simple rules and challenging solutions. The rules of Camp are simple:

1. Pair each tree with a tent adjacent horizontally or vertically. This should be a 1 to 1 relation.
2. Tents never touch each other even diagonally
3. The clues outside the grid indicate the number of tents on that row/column.

## Generating a random task

```nginx
GET https://shadify.dev/api/camp/generator
```

| Parameter  | Description                                                                                                                                           |
| ---------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| `width`    | _Optional_ <br><br> A number from 5 to 15 which corresponds to the width of the generated field. <br><br> The default value is 7.                     |
| `height`   | _Optional_ <br><br> A number from 5 to 15 which corresponds to the height of the generated field. <br><br> The default value is 7.                    |
| `solution` | _Optional_ <br><br> The true/false value that specifies whether the solution should be sent with the task or not. <br><br> The default value is true. |

Returned response:

```json
{
    "width": 7,
    "height": 7,
    "trees": 9,
    "rowTents": [1, 2, 2, 1, 0, 2, 1],
    "columnTents": [1, 1, 0, 2, 0, 2, 3],
    "task": [
        [0, 0, 0, 1, 0, 0, 1],
        [0, 0, 0, 0, 0, 1, 0],
        [0, 0, 0, 0, 0, 0, 0],
        [0, 1, 0, 0, 0, 1, 0],
        [0, 0, 0, 0, 0, 0, 1],
        [0, 0, 0, 0, 0, 0, 1],
        [1, 0, 0, 1, 0, 0, 0]
    ],
    "solution": [
        [0, 0, 0, 1, 0, 2, 1],
        [0, 0, 0, 2, 0, 1, 2],
        [0, 2, 0, 0, 0, 2, 0],
        [0, 1, 0, 0, 0, 1, 2],
        [0, 0, 0, 0, 0, 0, 1],
        [2, 0, 0, 2, 0, 0, 1],
        [1, 0, 0, 1, 0, 0, 2]
    ]
}
```

> -   `width` – the width of the generated field. <br>
> -   `heigth` – the height of the generated field.
> -   `trees` – the total number of trees on the field (and therefore the total number of tents that need to be set up). <br>
> -   `rowTents` – an array of values that corresponds to the number of tents that should be placed in the corresponding row from 1 to the `height` value. <br>
> -   `columnTents` – an array of values that corresponds to the number of tents that should be placed in the corresponding column from 1 to the `width` value. <br>
> -   `task` – the task that needs to be solved. Values of ones are trees around which you must place values of twos, which correspond to tents. <br>
> -   `solution` – a completely solved task, where the values of ones are trees, and the values of twos are tents.

## Verifying a completed task

```nginx
POST https://shadify.dev/api/camp/verifier
```

Example of JSON body:

```json
{
    "rowTents": [1, 2, 2, 1, 0, 2, 1],
    "columnTents": [1, 1, 0, 2, 0, 2, 3],
    "solution": [
        [0, 0, 2, 1, 0, 2, 1],
        [0, 0, 0, 0, 0, 1, 2],
        [0, 2, 0, 0, 0, 2, 0],
        [0, 1, 0, 0, 0, 1, 2],
        [0, 0, 0, 0, 0, 0, 1],
        [2, 0, 0, 2, 0, 0, 1],
        [1, 0, 0, 1, 0, 0, 2]
    ]
}
```

Returned response:

```json
{
    "isValid": false,
    "position": "row-1",
    "message": "the number of tents in each row must match the assignment"
}
```
