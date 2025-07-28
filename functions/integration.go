package functions

func SimpsonsOneThirdMethod(a float32, b float32, n float32, y []float32) float32 {

	h := (b - a) / n
	x := make([]float32, int(n+1))
	x[0] = a
	for i := 0; i < int(n); i++ {

		x[i+1] = x[i] + h
	}

	y = make([]float32, int(n+1))
	var result, oddsum, evensum float32
	for i := 1; i < int(n); i++ {

		if i%2 == 0 {
			evensum += y[i]
		} else {
			oddsum += y[i]
		}
	}
	sum := ((y[0] + y[int(n)]) + (4 * oddsum) + (2 * evensum))
	result = float32(((h / 3.0) * sum))

	return float32(result)
}
