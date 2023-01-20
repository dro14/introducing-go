package math

// Average function finds arithmetic mean of the slice values
func Average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// Min function finds the minimum value in the slice
func Min(xs []float64) float64 {
	min := xs[0]
	for _, v := range xs {
		if min > v {
			min = v
		}
	}
	return min
}

// Max function finds the maximum value in the slice
func Max(xs []float64) float64 {
	max := xs[0]
	for _, v := range xs {
		if max < v {
			max = v
		}
	}
	return max
}
