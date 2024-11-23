# Math

<p align="center"><img src="../images/math-expressions.png" alt="Math-expressions"/></p>

A module for generating random mathematical expressions. Great for creating various simulators, where you'll have to wiggle your brain. You can see [my application](https://cheatsnake.github.io/MindMath/) as an example.

## Generating an expression with addition

```nginx
GET https://shadify.yurace.pro/api/math/add
```

-   Generating an expression with subtraction

```nginx
GET https://shadify.yurace.pro/api/math/sub
```

-   Generating an expression with multiplication

```nginx
GET https://shadify.yurace.pro/api/math/mul
```

-   Generating an expression with division

```nginx
GET https://shadify.yurace.pro/api/math/div
```

The following parameters are available for the above endpoints:

| Parameter                  | Description                                                                                                                                                                                                                                                                                                                                                                 |
| -------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `min-first`, `max-first`   | _Optional_ <br><br> Parameters that change the range of generated numbers only for the first number present in the expression. <br><br> The default value for `min-first` is 1, for `max-first` it is 99.                                                                                                                                                                   |
| `min-second`, `max-second` | _Optional_ <br><br> Parameters that change the range of generated numbers only for the second number present in the expression. These parameters do not affect expressions with division, because the second number is chosen randomly from the list of possible divisors of the first number. <br><br> The default value for `min-second` is 1, for `max-second` it is 99. |
| `negative`                 | _Optional_ <br><br> The number 0/1 which enables/disables the possibility of generating a result with a negative value for expressions with subtraction. Initially, when generating an expression with subtraction, the first number is always greater than the second number. To change this behavior, use this parameter. <br><br> Default value is 0.                    |

Returned response:

```json
{
    "first": 17,
    "second": 56,
    "operation": "+",
    "expression": "17 + 56",
    "answer": 73
}
```

## Generating a quadratic equation

<p align="center"><img src="../images/quadratic-equation.png" alt="Quadratic-equation"/></p>

[Quadratic equation](https://en.wikipedia.org/wiki/Quadratic_equation) is an equation of the form ax2 + bx + c = 0, where the coefficients a, b and c are arbitrary numbers.

```nginx
GET https://shadify.yurace.pro/api/math/quad
```

| Parameter        | Description                                                                                                                             |
| ---------------- | --------------------------------------------------------------------------------------------------------------------------------------- |
| `min-a`, `max-a` | _Optional_ <br><br> Minimum and maximum values of the coefficient A. <br><br> The default value for `min-a` is 1, for `max-a` it is 20. |
| `min-b`, `max-b` | _Optional_ <br><br> Minimum and maximum values of the coefficient B. <br><br> The default value for `min-b` is 1, for `max-b` it is 40. |
| `min-c`, `max-c` | _Optional_ <br><br> Minimum and maximum values of the coefficient C. <br><br> The default value for `min-c` is 1, for `max-c` it is 20. |

Returned response:

```json
{
    "equation": "1x^2 + 14x + 24 = 0",
    "a": 1,
    "b": 14,
    "c": 24,
    "discriminant": 100,
    "x1": "-12.000",
    "x2": "-2.000"
}
```
