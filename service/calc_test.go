package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CalculatorTestSuite struct {
	suite.Suite
	calcService Calculator
}

func (suite *CalculatorTestSuite) SetupTest() {
	suite.calcService = Calculator{}
}

// func TestCalculator_Add(t *testing.T) {
// 	t.Run("Calculator Add operator testing", func(t *testing.T) {
// 		numSample1 := 10
// 		numSample2 := 20

// 		calc := Calculator{
// 			Num1: numSample1,
// 			Num2: numSample2,
// 		}

// 		if calc.Add() != 30 {
// 			t.Error("It should return 30")
// 		}
// 	})
// }

// func TestCalculator_Sub(t *testing.T) {
// 	t.Run("Calculator Sub operator testing", func(t *testing.T) {
// 		numSample1 := 3
// 		numSample2 := 2

// 		calc := Calculator{
// 			Num1: numSample1,
// 			Num2: numSample2,
// 		}

// 		if calc.Sub() != 1 {
// 			t.Error("It should return 1")
// 		}
// 	})
// }

// func TestCalculator_Div(t *testing.T) {
// 	t.Run("Calculator Div operator testing", func(t *testing.T) {
// 		numSample1 := 5
// 		numSample2 := 2

// 		calc := Calculator{
// 			Num1: numSample1,
// 			Num2: numSample2,
// 		}

// 		result, err := calc.Div()

// 		if result != 2 {
// 			t.Error("It should return 2")
// 		}

// 		if err != nil {
// 			t.Error(err)
// 		}

// 	})

// t.Run("Calculator Div operator testing", func(t *testing.T) {
// 	numSample1 := -1
// 	numSample2 := -1

// 	calc := Calculator{
// 		Num1: numSample1,
// 		Num2: numSample2,
// 	}

// 	_, err := calc.Div()

// 	if err != nil && err.Error() != "value should be greater than 0" {
// 		t.Error("value should be greater than 0")
// 	}

// })
// }

// func TestCalculator_Mul(t *testing.T) {
// 	t.Run("Calculator Mul operator testing", func(t *testing.T) {
// 		numSample1 := 5
// 		numSample2 := 2

// 		calc := Calculator{
// 			Num1: numSample1,
// 			Num2: numSample2,
// 		}

// 		result, err := calc.Mul()

// 		if result != 10 {
// 			t.Error("It should return 10")
// 		}

// 		if err != nil {
// 			t.Error(err)
// 		}

// 	})
// }

func (suite *CalculatorTestSuite) TestCalculator_Sub() {
	suite.calcService.Num1 = 2
	suite.calcService.Num2 = 1

	result := suite.calcService.Sub()
	assert.Equal(suite.T(), 1, result)
}

func (suite *CalculatorTestSuite) TestCalculator_Add() {
	// suite.calcService.Num1 = 2
	// suite.calcService.Num2 = 1

	test := []struct {
		num1 int
		num2 int
		res  int
	}{
		{1, 1, 2},
		{-1, 2, -1},
	}

	for _, test1 := range test {
		suite.calcService.Num1 = test1.num1
		suite.calcService.Num2 = test1.num2
		result, err := suite.calcService.Add()
		if err != nil {
			assert.NotNil(suite.T(), err)
			assert.Equal(suite.T(), test1.res, result)
		} else {
			assert.Nil(suite.T(), err)
			assert.Equal(suite.T(), test1.res, result)
		}
	}
}

func TestCaclculatorTestSuite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}
