// Package mathutils provides basic mathematical utilities.
//
// This package includes simple arithmetic operations like addition.
package Add

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes two integers and returns their sum.
//
// For more information about addition, visit:
// https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a1, a2 T) T {
	return a1 + a2

}
