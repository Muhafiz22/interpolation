package main

import (
	"errors"
)

func main() {

}
func validateInput(x []float32, y []float32) (error, []float32, []float32) {

	if len(x) != len(y) {
		return errors.New("Insufficient Input"), nil, nil
	}
	return nil, x, y

}
