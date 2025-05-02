// Package mathutils provides basic mathematical utilities.
//
// This package includes simple arithmetic operations like addition.
package Add

import (
	"golang.org/x/exp/constraints"
)

// Number is a type constraint that permits any integer or floating-point type.
type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes two numbers of type Number and returns their sum.
//
// Supported types include all integers and floating-point numbers.
// For more information about addition, visit:
// https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
