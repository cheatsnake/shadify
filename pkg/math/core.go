package math

import (
	"math"
	"strconv"

	"github.com/cheatsnake/shadify/pkg/assists"
)

// Get a random addition expression with a given range of numbers
func GetAddition(minFirst, maxFirst, minSecond, maxSecond int) BasicExpression {
	operation := "+"
	first := assists.GetRandomInteger(minFirst, maxFirst)
	second := assists.GetRandomInteger(minSecond, maxSecond)
	expression := strconv.Itoa(first) + " " + operation + " " + strconv.Itoa(second)
	answer := first + second

	return BasicExpression{
		first, second, operation, expression, answer,
	}
}

// Get a random subtraction expression with a given range of numbers
func GetSubtraction(minFirst, maxFirst, minSecond, maxSecond int, allowNegative bool) BasicExpression {
	operation := "-"
	first := assists.GetRandomInteger(minFirst, maxFirst)
	second := assists.GetRandomInteger(minSecond, maxSecond)

	if second > first && !allowNegative {
		second, first = first, second
	}

	expression := strconv.Itoa(first) + " " + operation + " " + strconv.Itoa(second)
	answer := first - second

	return BasicExpression{
		first, second, operation, expression, answer,
	}
}

// Get a random multiplication expression with a given range of numbers
func GetMultiplication(minFirst, maxFirst, minSecond, maxSecond int) BasicExpression {
	operation := "*"
	first := assists.GetRandomInteger(minFirst, maxFirst)
	second := assists.GetRandomInteger(minSecond, maxSecond)
	expression := strconv.Itoa(first) + " " + operation + " " + strconv.Itoa(second)
	answer := first * second

	return BasicExpression{
		first, second, operation, expression, answer,
	}
}

// Get a random division expression with a given range of numbers
func GetDivision(minFirst, maxFirst int) BasicExpression {
	operation := "/"
	first := assists.GetRandomInteger(minFirst, maxFirst)

	if first == 0 {
		first = 1
	}

	divisors := findDivisors(first)

	if len(divisors) > 2 {
		divisors = divisors[1 : len(divisors)-1]
	}

	index := assists.GetRandomInteger(0, len(divisors)-1)
	second := divisors[index]
	expression := strconv.Itoa(first) + " " + operation + " " + strconv.Itoa(second)
	answer := first / second

	return BasicExpression{
		first, second, operation, expression, answer,
	}
}

// Get a random quadratic expression with a given range of coefficients
func GetQuadratic(minA, maxA, minB, maxB, minC, maxC int) QuadraticEquation {
	a := assists.GetRandomInteger(minA, maxA)
	b := assists.GetRandomInteger(minB, maxB)
	c := assists.GetRandomInteger(minC, maxC)

	equation := strconv.Itoa(a) + "x^2 + " + strconv.Itoa(b) + "x + " + strconv.Itoa(c) + " = 0"

	discriminant := (b * b) - 4*a*c
	if discriminant < 0 {
		return QuadraticEquation{equation, a, b, c, discriminant, "none", "none"}
	}

	x1 := (-float64(b) - (math.Sqrt(float64(discriminant)))) / (2 * float64(a))
	x2 := (-float64(b) + (math.Sqrt(float64(discriminant)))) / (2 * float64(a))

	return QuadraticEquation{
		equation, a, b, c, discriminant,
		strconv.FormatFloat(x1, 'f', 3, 64),
		strconv.FormatFloat(x2, 'f', 3, 64),
	}
}
