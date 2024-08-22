package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {

	r := []float64{1891, 1901, 1911, 1921, 1931}

	s := []float64{46, 66, 81, 93, 101}
	err, _, _ := validateInput(r, s)

	if err != nil {
		fmt.Println(err)
	}

}

func TestCheckForward(t *testing.T) {

	r := []float64{1891, 1901, 1911, 1921, 1931}

	var p float64 = 1925.0
	c := checkForward(r, p)

	if c {
		fmt.Println("use forward")
	} else {
		fmt.Println("use backward")
	}
}

func TestCalculateForwardDiffInterpolation(t *testing.T) {

	x := []float64{45, 50, 55, 60}
	y := []float64{0.7071, 0.7660, 0.8192, 0.8660}

	xu := 52.0

	d := calculateForwardDiffInterpolation(x, y, xu)
	fmt.Println(d)

}

func TestcalculateBackwardDiffInterpolation(t *testing.T) {

	x := []float64{1891, 1901, 1911, 1921, 1931}
	y := []float64{46, 66, 81, 93, 101}

	xu := 1925.0

	e := calculateBackwardDiffInterpolation(x, y, xu)
	fmt.Println(e)

}
