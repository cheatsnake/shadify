# Countries

Countries module allows you to generate quizes such as guess the capital or guess the country from an image of the flag. It is a simple and useful module for creating applications to test and practice knowledge of all countries.

## Generating capital quiz

<p align="center"><img src="../images/guess-capital.png" alt="Guess capital"/></p>

```nginx
GET https://shadify.yurace.pro/api/countries/capital-quiz
```

| Parameter  | Description                                                                                                                                                                                                                |
| ---------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `variants` | _Optional_ <br><br> A number from 2 to 6 corresponding to the number of different options from which you have to choose the correct capital of the given country. <br><br> The default value is 4.                         |
| `amount`   | _Optional_ <br><br> A number from 1 to 20 which is responsible for the number of quizzes returned. The use of this parameter ensures that among all quizzes received, all will be unique. <br><br> The default value is 1. |

Returned response:

```json
{
    "country": "Cyprus",
    "flag": "https://flagcdn.com/w320/cy.png",
    "variants": ["Nicosia", "Juba", "Oslo", "Jamestown"],
    "answer": "Nicosia"
}
```

> -   `country` - the country for which you must guess the capital. <br>
> -   `flag` - country's flag image (it powered by [Flagpedia API](https://flagpedia.net/download/api), so you can customize it). <br>
> -   `variants` - possible options for answering. <br>
> -   `answer` - correct answer.

## Generating country quiz

<p align="center"><img src="../images/guess-country.png" alt="Guess country"/></p>

```nginx
GET https://shadify.yurace.pro/api/countries/country-quiz
```

| Parameter  | Description                                                                                                                                                                                                                |
| ---------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `variants` | _Optional_ <br><br> A number from 2 to 6 corresponding to the number of different options from which you have to choose the correct country of the given flag image. <br><br> The default value is 4.                      |
| `amount`   | _Optional_ <br><br> A number from 1 to 20 which is responsible for the number of quizzes returned. The use of this parameter ensures that among all quizzes received, all will be unique. <br><br> The default value is 1. |

Returned response:

```json
{
    "flag": "https://flagcdn.com/w320/dk.png",
    "variants": ["Yemen", "Denmark", "Norfolk Island", "Vietnam"],
    "answer": "Denmark"
}
```

> `flag` - the flag of the country you have to guess. <br> `variants` - possible options for answering. <br> `answer` - correct answer.
