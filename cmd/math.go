package cmd

import (
	"fmt"
	"math"
	"strconv"
)

func convertArgs(first string, second string) (float64, float64, error) {
	num1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("first argument is invalid")
	}
	num2, err := strconv.ParseFloat(second, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("second argument is invalid")
	}
	return num1, num2, nil
}

func Add(first string, second string) (string, error) {
	num1, num2, err := convertArgs(first, second)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%f", num1+num2), nil
}

func Subtract(from string, subtract string) (string, error) {
	num1, num2, err := convertArgs(from, subtract)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%f", num1-num2), nil
}

func Multiply(first string, second string, shouldRoundUp bool) (string, error) {
	num1, num2, err := convertArgs(first, second)
	if err != nil {
		return "", err
	}
	if shouldRoundUp {
		return fmt.Sprintf("%.2f", num1*num2), nil
	}
	return fmt.Sprintf("%f", num1*num2), nil
}

func Divide(divide string, by string, shouldRoundUp bool) (result string, e error) {
	num1, num2, err := convertArgs(divide, by)
	if num2 == 0 {
		return "", fmt.Errorf("zero-division-error: can not divide by zero")
	}
	if err != nil {
		return "", err
	}
	if shouldRoundUp {
		return fmt.Sprintf("%.2f", num1/num2), nil
	}
	return fmt.Sprintf("%f", num1/num2), nil
}

func Square(num string) (string, error) {
	number, _, err := convertArgs(num, "0.0")
	if err != nil {
		return "", fmt.Errorf("not a number: you need to provide a number to be squared. The provided argument was not a number")
	}
	return fmt.Sprintf("%f", number*number), nil
}

func Sqrt(num string) (string, error) {
	number, _, err := convertArgs(num, "0.0")
	if err != nil {
		return "", fmt.Errorf("not a number: The provided argument was not a number therefore we can not calculate it's square root")
	}
	return fmt.Sprintf("%f", math.Sqrt(number)), nil

}

func Percentage(whole string, percent string) (string, error) {
	num1, num2, err := convertArgs(whole, percent)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%f", num1*num2/100), nil
}
