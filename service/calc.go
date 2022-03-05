package services

import "errors"

type Calculator struct {
	Num1 int
	Num2 int
}

func (c *Calculator) Add() (int, error) {
	if c.Num1 < 0 || c.Num2 < 0 {
		return -1, errors.New("tidak boleh negatif")
	}
	return c.Num1 + c.Num2, nil
}

func (c *Calculator) Sub() int {
	return c.Num1 - c.Num2
}

// func (c *Calculator) Div() (float64, error) {
// 	if c.Num1 < 1 || c.Num2 < 1 {
// 		return 0, errors.New("number should be greater than 0")
// 	}
// 	return float64(c.Num1 / c.Num2), nil
// }

// func (c *Calculator) Mul() (int, error) {
// 	return c.Num1 * c.Num2, nil
// }
