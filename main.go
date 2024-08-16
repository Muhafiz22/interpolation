package main

import (
	"errors"
	"math"
)

func main() {

}
func validateInput(x []float64, y []float64) ([]float64, []float64, error) {

	if len(x) != len(y) {
		return nil, nil, errors.New("insufficient Input")
	}
	return x, y, nil

}

func checkForward(x []float64, xu float64) bool {

	return xu <= x[int(math.Abs(float64(len(x)/2)))]

}

func calculateForwardDiffInterpolation(x []float64, y []float64, xu float64) float64 {

	n := len(x)
	delta_Y := make([][]float64, n)

	for i := range delta_Y {

		delta_Y[i] = make([]float64, n-i) //Creating rows corresponding to different coloumn sizes
		delta_Y[i][0] = y[i]
	}

	for j := 1; j < n; j++ { //iterating for diffn degree frwrd differences
		for i := 0; i < n-j; i++ { // iterating no.of row values to be calculated for diff frwrd diff

			delta_Y[i][j] = delta_Y[i+1][j-1] - delta_Y[i][j-1] //Calculating higher frwd diffn
		}
	}

	result := y[0]
	u := (xu - x[0]) / (x[1] - x[0])
	uproduct := 1.0

	for i := 1; i < n; i++ {

		uproduct *= (u - float64(i-1))
		result += (uproduct * delta_Y[0][i]) / factorial(i)
	}

	return result
}

func factorial(n int) float64 {

	fact := 1.0
	if n == 0 || n == 1 {
		return 1
	}
	for i := 2; i <= n; i++ {

		fact *= float64(i)
	}
	return fact
}
