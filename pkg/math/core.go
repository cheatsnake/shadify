package math

import (
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
