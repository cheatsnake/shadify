# Schulte

<p align="center"><img src="../images/schulte.gif" alt="Schulte"/></p>

[**Schulte Tables**](https://youtu.be/4R1kycHqsqM) – tables with randomly arranged objects (usually numbers or letters) used to test and develop quickness of finding these objects in a certain order (usually in ascending order for numbers and alphabetical order for letters). Exercises with tables can improve peripheral visual perception, which will have a positive impact on the skill of speed reading.

## Generating a Random Table

```nginx
GET https://shadify.yurace.pro/api/schulte/generator
```

| Parameter | Description                                                                                                                                                                                                                       |
| --------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `size`    | _Optional_ <br><br> The size of the generated table. The available range is from 1 to 15. <br><br> The default value is 5.                                                                                                        |
| `mode`    | _Optional_ <br><br> The _number_ / _alphabet_ string, which defines the generated values for the table: numbers or letters. If _alphabet_ is selected, the `size` parameter will always be 5. <br><br> Default value is _number_. |

Returned response:

```json
{
    "grid": [
        [9, 18, 6, 1, 5],
        [11, 24, 25, 15, 19],
        [13, 7, 10, 21, 23],
        [20, 22, 12, 17, 16],
        [2, 8, 4, 14, 3]
    ]
}
```
