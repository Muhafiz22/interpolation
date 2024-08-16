package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {

	r := []float64{40, 50, 60, 70, 80}

	s := []float64{31.0, 73.0, 124.0, 159.0, 190.0}
	err, _, _ := validateInput(r, s)

	if err != nil {
		fmt.Println(err)
	}

}

func TestCheckForward(t *testing.T) {

	r := []float64{40, 50, 60, 70, 80}

	var p float64 = 45.0
	c := checkForward(r, p)

	if c {
		fmt.Println("use forward")
	} else {
		fmt.Println("use backward")
	}
}

func TestCalculateForwardDiffInterpolation(t *testing.T) {

	x := []float64{40, 50, 60, 70, 80}

	y := []float64{31.0, 73.0, 124.0, 159.0, 190.0}
	xu := 45.0
	d := calculateForwardDiffInterpolation(x, y, xu)
	fmt.Println(d)
	if d == 15.468750000000002 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed")
	}
}
