# Kuromasu

<p align="center"><img src="../images/kuromasu.png" alt="Kuromasu"/></p>

[**Kuromasu**](https://en.wikipedia.org/wiki/Kuromasu) is a puzzle played on a rectangular grid. Some of these cells have numbers in them. Each cell may be either black or white. The object is to determine what type each cell is. The following rules determine which cells are which:

1. Each number on the board represents the number of white cells that can be seen from that cell, including itself. A cell can be seen from another cell if both cells are within the same row or column, and there are no black cells between them in that row or column.
2. Numbered cells must not be black.
3. No two black cells must be horizontally or vertically adjacent.
4. All the white cells must be connected horizontally or vertically.

## Generating a random task

```nginx
GET https://shadify.dev/api/kuromasu/generator
```

| Parameter | Description                                                                                                                                         |
| --------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| `width`   | _Optional_ <br><br> A number from 4 to 15 which corresponds to the width of the generated field. <br><br> The default value is 5.                   |
| `height`  | _Optional_ <br><br> A number from 4 to 15 which corresponds to the height of the generated field. <br><br> The default value is 5.                  |
| `fill`    | _Optional_ <br><br> A number from 10 to 50 which corresponds to the % level of filling the task with ready cells. <br><br> The default value is 30. |

Returned response:

```json
{
    "width": 5,
    "height": 5,
    "task": [
        ["o", "o", "o", "o", "6"],
        ["o", "o", "8", "o", "9"],
        ["o", "o", "o", "7", "o"],
        ["5", "o", "o", "o", "o"],
        ["9", "o", "8", "5", "o"]
    ],
    "solution": [
        ["6", "4", "x", "4", "6"],
        ["9", "7", "8", "7", "9"],
        ["9", "7", "8", "7", "9"],
        ["5", "x", "4", "x", "5"],
        ["9", "5", "8", "5", "9"]
    ]
}
```

> -   `width` – the width of the generated grid. <br>
> -   `heigth` – the height of the generated grid. <br>
> -   `task` – the task that needs to be solved by arranging the boxes (values of `x`) in the right cells. <br>
> -   `solution` – a completely solved task, where cells with `x` values are boxes and numbers is just view counts of each particular cell.

## Verifying a completed task

```nginx
POST https://shadify.dev/api/kuromasu/verifier
```

Example of JSON body:

```json
{
    "solution": [
        ["6", "4", "x", "4", "x"],
        ["9", "7", "8", "7", "9"],
        ["9", "7", "8", "7", "9"],
        ["5", "x", "4", "x", "5"],
        ["9", "5", "8", "5", "9"]
    ]
}
```

Returned response:

```json
{
    "isValid": false,
    "position": {
        "row": 1,
        "column": 4
    }
}
```
