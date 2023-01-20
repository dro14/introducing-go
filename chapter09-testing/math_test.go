package math

import "testing"

type test1 struct {
	input  []float64
	output float64
}

func TestAverage(t *testing.T) {
	tests := []test1{
		{[]float64{1, 2}, 1.5},
		{[]float64{1, 1, 1, 1, 1, 1}, 1},
		{[]float64{-1, 1}, 0},
		{[]float64{}, 0},
	}

	for _, pair := range tests {
		v := Average(pair.input)
		if v != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", v,
			)
		}
	}
}
