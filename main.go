package main

import (
	"errors"
	"math"
)

func main() {

}
func validateInput(x []float32, y []float32) (error, []float32, []float32) {

	if len(x) != len(y) {
		return errors.New("Insufficient Input"), nil, nil
	}
	return nil, x, y

}

func checkForward(x []float32, xu float32) bool {

	return xu <= x[int(math.Abs(float64(len(x)/2)))]

}
