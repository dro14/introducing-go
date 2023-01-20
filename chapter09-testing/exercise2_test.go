// Series of tests for the Min and Max functions written in the previous chapter:

package math

import "testing"

type test2 struct {
	xs       []float64
	max, min float64
}

func TestMath(t *testing.T) {
	tests := []test2{
		{[]float64{1, 4, -9, -4.6, 0, 123423.134}, 123423.134, -9},
		{[]float64{5, 5, 5, 5, 5, 5, 5, 5}, 5, 5},
		{[]float64{1, 1, 2, 3, 4, 5, 6, 7}, 7, 1},
		{[]float64{}, 0, 0},
	}

	for _, unit := range tests {
		max := Max(unit.xs)
		if max != unit.max {
			t.Errorf("expected %f got %f", unit.max, max)
		}
		min := Min(unit.xs)
		if min != unit.min {
			t.Errorf("expected %f got %f", unit.min, min)
		}
	}
}
