package main

import (
	"fmt"
	"math"
)

func checkForward(x []float32, y []float32, xu float32) (bool, float32, [][]float32) {
	c := xu <= x[int(math.Abs(float64(len(x)/2)))]

	var deltaY [][]float32
	var result float32
	if c {
		deltaY, result = calculateForwardDiffInterpolation(x, y, xu)
	} else {

		deltaY, result = calculateBackwardDiffInterpolation(x, y, xu)
	}
	return c, result, deltaY
}

func calculateForwardDiffInterpolation(x []float32, y []float32, xu float32) ([][]float32, float32) {

	n := len(x)
	delta_Y := make([][]float32, n)

	for i := range delta_Y {

		delta_Y[i] = make([]float32, n-i) //Creating rows corresponding to different coloumn sizes
		delta_Y[i][0] = y[i]
	}

	for j := 1; j < n; j++ { //iterating for diffn degree frwrd differences
		for i := 0; i < n-j; i++ { // iterating no.of row values to be calculated for diff frwrd diff

			delta_Y[i][j] = delta_Y[i+1][j-1] - delta_Y[i][j-1] //Calculating higher frwd diffn
		}
	}

	result := y[0]
	u := (xu - x[0]) / (x[1] - x[0])
	fmt.Println("U:", u)
	var uproduct float32 = 1.0

	for i := 1; i < n; i++ {

		uproduct *= (u - float32(i-1))
		perterm := (uproduct * delta_Y[0][i]) / factorial(i)
		result += perterm
	}
	return delta_Y, result
}

func calculateBackwardDiffInterpolation(x []float32, y []float32, xu float32) ([][]float32, float32) {
	n := len(x)
	delta_Y := make([][]float32, n)
	for i := range delta_Y {

		delta_Y[i] = make([]float32, n)
		delta_Y[i][0] = y[i]
	}
	for j := 1; j < n; j++ {
		for i := n - 1; i >= j; i-- {
			delta_Y[i][j] = delta_Y[i][j-1] - delta_Y[i-1][j-1]
		}
	}
	result := y[n-1]
	v := (xu - x[n-1]) / (x[1] - x[0])
	var vproduct float32 = 1.0
	for i := 0; i < n-1; i++ {

		vproduct *= (v + float32(i))
		perterm := (vproduct * delta_Y[n-1][i+1]) / factorial(i+1)
		result += perterm
	}
	return delta_Y, result
}

func factorial(n int) float32 {
	var fact float32 = 1.0
	if n == 0 || n == 1 {
		return 1
	}
	for i := 2; i <= n; i++ {

		fact *= float32(i)
	}
	return fact
}
