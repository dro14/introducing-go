package math

// Finds arithmetic mean of the slice values
//func Average(xs []float64) float64 {
//	total := 0.0
//	for _, x := range xs {
//		total += x
//	}
//	return total / float64(len(xs))
//}

// Min function finds the minimum value in the slice
func Min(xs []float64) float64 {
	min := 0.0
	for i, v := range xs {
		if i == 0 || min > v {
			min = v
		}
	}
	return min
}

// Max function finds the maximum value in the slice
func Max(xs []float64) float64 {
	max := 0.0
	for i, v := range xs {
		if i == 0 || max < v {
			max = v
		}
	}
	return max
}
