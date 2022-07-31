package math

import "testing"

type TestBasicExpressionValues struct {
	minFirst, maxFirst, minSecond, maxSecond int
	allowNegative                            bool
	expected                                 BasicExpression
}

type TestQuadraticEquationValues struct {
	minA, maxA, minB, maxB, minC, maxC int
	expected                           QuadraticEquation
}

func TestGetAddition(t *testing.T) {
	testValues := []TestBasicExpressionValues{
		{2, 2, 3, 3, false, BasicExpression{2, 3, "+", "2 + 3", 5}},
		{-5, -5, -2, -2, false, BasicExpression{-5, -2, "+", "-5 + -2", -7}},
		{0, 0, 0, 0, false, BasicExpression{0, 0, "+", "0 + 0", 0}},
	}

	for _, tv := range testValues {
		result := GetAddition(tv.minFirst, tv.maxFirst, tv.minSecond, tv.maxSecond)
		if result != tv.expected {
			t.Errorf(
				"GetAddition(%d, %d, %d, %d) FAILED. Expected %+v, but got %+v\n",
				tv.minFirst, tv.maxFirst, tv.minSecond, tv.maxSecond, tv.expected, result)
		} else {
			t.Logf(
				"GetAddition(%d, %d, %d, %d) PASSED\n",
				tv.minFirst, tv.maxFirst, tv.minSecond, tv.maxSecond)
		}
	}
}

func TestGetSubtraction(t *testing.T) {
	testValues := []TestBasicExpressionValues{
		{7, 7, 5, 5, false, BasicExpression{7, 5, "-", "7 - 5", 2}},
		{3, 3, 9, 9, false, BasicExpression{9, 3, "-", "9 - 3", 6}},
		{3, 3, 9, 9, true, BasicExpression{3, 9, "-", "3 - 9", -6}},
		{-5, -5, -4, -4, true, BasicExpression{-5, -4, "-", "-5 - -4", -1}},
		{-5, -5, -4, -4, false, BasicExpression{-4, -5, "-", "-4 - -5", 1}},
		{0, 0, 0, 0, false, BasicExpression{0, 0, "-", "0 - 0", 0}},
	}

	for _, tv := range testValues {
		result := GetSubtraction(tv.minFirst, tv.maxFirst, tv.minSecond, tv.maxSecond, tv.allowNegative)
		if result != tv.expected {
			t.Errorf(
				"GetSubtraction(%d, %d, %d, %d, %t) FAILED. Expected %+v, but got %+v\n",
				tv.minFirst, tv.maxFirst, tv.minSecond, tv.maxSecond, tv.allowNegative, tv.expected, result)
		} else {
			t.Logf(
				"GetSubtraction(%d, %d, %d, %d, %t) PASSED\n",
				tv.minFirst, tv.maxFirst, tv.minSecond, tv.maxSecond, tv.allowNegative)
		}
	}
}

func TestGetMultiplication(t *testing.T) {
	testValues := []TestBasicExpressionValues{
		{2, 2, 3, 3, false, BasicExpression{2, 3, "*", "2 * 3", 6}},
		{-5, -5, -2, -2, false, BasicExpression{-5, -2, "*", "-5 * -2", 10}},
		{-7, -7, 1, 1, false, BasicExpression{-7, 1, "*", "-7 * 1", -7}},
		{99, 99, 0, 0, false, BasicExpression{99, 0, "*", "99 * 0", 0}},
		{0, 0, 0, 0, false, BasicExpression{0, 0, "*", "0 * 0", 0}},
	}

	for _, tv := range testValues {
		result := GetMultiplication(tv.minFirst, tv.maxFirst, tv.minSecond, tv.maxSecond)
		if result != tv.expected {
			t.Errorf(
				"GetMultiplication(%d, %d, %d, %d) FAILED. Expected %+v, but got %+v\n",
				tv.minFirst, tv.maxFirst, tv.minSecond, tv.maxSecond, tv.expected, result)
		} else {
			t.Logf(
				"GetMultiplication(%d, %d, %d, %d) PASSED\n",
				tv.minFirst, tv.maxFirst, tv.minSecond, tv.maxSecond)
		}
	}
}

func TestFindDivisors(t *testing.T) {
	result := findDivisors(0)
	if len(result) != 0 {
		t.Errorf("findDivisors(0) FAILED. Expected %+v, but got %+v\n", 0, result)
	} else {
		t.Logf("findDivisors(0) PASSED\n")
	}

	result = findDivisors(1)
	if len(result) != 1 {
		t.Errorf("findDivisors(1) FAILED. Expected %+v, but got %+v\n", 1, result)
	} else {
		t.Logf("findDivisors(1) PASSED\n")
	}

	result = findDivisors(7)
	if len(result) != 2 {
		t.Errorf("findDivisors(7) FAILED. Expected %+v, but got %+v\n", 2, result)
	} else {
		t.Logf("findDivisors(7) PASSED\n")
	}

	result = findDivisors(-7)
	if len(result) != 4 {
		t.Errorf("findDivisors(-7) FAILED. Expected %+v, but got %+v\n", 4, result)
	} else {
		t.Logf("findDivisors(-7) PASSED\n")
	}
}

func TestGetDivision(t *testing.T) {
	testValues := []TestBasicExpressionValues{
		{0, 0, 0, 0, false, BasicExpression{1, 1, "/", "1 / 1", 1}},
		{1, 1, 0, 0, false, BasicExpression{1, 1, "/", "1 / 1", 1}},
	}

	for _, tv := range testValues {
		result := GetDivision(tv.minFirst, tv.maxFirst)
		if result != tv.expected {
			t.Errorf(
				"GetDivision(%d, %d) FAILED. Expected %+v, but got %+v\n",
				tv.minFirst, tv.maxFirst, tv.expected, result)
		} else {
			t.Logf(
				"GetDivision(%d, %d) PASSED\n",
				tv.minFirst, tv.maxFirst)
		}
	}
}

func TestGetQuadratic(t *testing.T) {
	testValues := []TestQuadraticEquationValues{
		{2, 2, -5, -5, 3, 3, QuadraticEquation{"2x^2 + -5x + 3 = 0", 2, -5, 3, 1, "1.000", "1.500"}},
		{2, 2, 8, 8, 8, 8, QuadraticEquation{"2x^2 + 8x + 8 = 0", 2, 8, 8, 0, "-2.000", "-2.000"}},
		{1, 1, -5, -5, 6, 6, QuadraticEquation{"1x^2 + -5x + 6 = 0", 1, -5, 6, 1, "2.000", "3.000"}},
		{3, 3, 10, 10, 9, 9, QuadraticEquation{"3x^2 + 10x + 9 = 0", 3, 10, 9, -8, "none", "none"}},
		{0, 0, 0, 0, 0, 0, QuadraticEquation{"0x^2 + 0x + 0 = 0", 0, 0, 0, 0, "R", "R"}},
	}

	for _, tv := range testValues {
		result := GetQuadratic(tv.minA, tv.maxA, tv.minB, tv.maxB, tv.minC, tv.maxC)
		if result != tv.expected {
			t.Errorf(
				"GetQuadratic(%d, %d, %d, %d, %d, %d) FAILED. Expected %+v, but got %+v\n",
				tv.minA, tv.maxA, tv.minB, tv.maxB, tv.minC, tv.maxC, tv.expected, result)
		} else {
			t.Logf(
				"GetQuadratic(%d, %d, %d, %d, %d, %d) PASSED\n",
				tv.minA, tv.maxA, tv.minB, tv.maxB, tv.minC, tv.maxC)
		}
	}
}
