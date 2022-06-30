package math

import (
	"math"
	"strconv"

	"github.com/cheatsnake/shadify/pkg/assists"
)

func NewCore() *Core {
	return &Core{}
}

func (mc *Core) Addition(minFirst, maxFirst, minSecond, maxSecond int) BasicExpression {
	operation := "+"
	first := assists.GetRandomInteger(minFirst, maxFirst)
	second := assists.GetRandomInteger(minSecond, maxSecond)
	expression := strconv.Itoa(first) + " " + operation + " " + strconv.Itoa(second)
	answer := first + second

	return BasicExpression{
		first, second, operation, expression, answer,
	}
}

func (mc *Core) Subtraction(minFirst, maxFirst, minSecond, maxSecond int, allowNegative bool) BasicExpression {
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

func (mc *Core) Multiplication(minFirst, maxFirst, minSecond, maxSecond int) BasicExpression {
	operation := "*"
	first := assists.GetRandomInteger(minFirst, maxFirst)
	second := assists.GetRandomInteger(minSecond, maxSecond)
	expression := strconv.Itoa(first) + " " + operation + " " + strconv.Itoa(second)
	answer := first * second

	return BasicExpression{
		first, second, operation, expression, answer,
	}
}

func (mc *Core) Division(minFirst, maxFirst int) BasicExpression {
	operation := "/"
	first := assists.GetRandomInteger(minFirst, maxFirst)
	divisors := FindDivisors(first);

	if len(divisors) > 2 {
		divisors = divisors[1:len(divisors) - 1]
	}

	index := assists.GetRandomInteger(0, len(divisors) - 1)
	second := divisors[index]
	expression := strconv.Itoa(first) + " " + operation + " " + strconv.Itoa(second)
	answer := first / second

	return BasicExpression{
		first, second, operation, expression, answer,
	}
}

func (mc* Core) Quadratic(minA, maxA, minB, maxB, minC, maxC int) QuadraticEquation {
	a := assists.GetRandomInteger(minA, maxA)
	b := assists.GetRandomInteger(minB, maxB)
	c := assists.GetRandomInteger(minC, maxC)

	equation := strconv.Itoa(a) + "x^2 + " + strconv.Itoa(b) + "x + " + strconv.Itoa(c) + " = 0"

	discriminant := (b * b) - 4 * a * c
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