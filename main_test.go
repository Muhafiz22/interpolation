package main

import (
	"fmt"
	"testing"
)

func TestCheckForward(t *testing.T) {

	var input_x []float32 = []float32{40, 50, 60, 70, 80}

	var input_xu float32 = 45.0
	c := checkForward(input_x, input_xu)

	if c {
		fmt.Println("Use forward")
	} else {
		fmt.Println("use backward")
	}
}

func TestCalculateBackwardDiffInterpolation(t *testing.T) {

	var input_x []float32 = []float32{1891, 1901, 1911, 1921, 1931}
	var input_y []float32 = []float32{46, 66, 81, 93, 101}

	var input_xu float32 = 1925.0

	c := checkForward(input_x, input_xu)
	if c {
		fmt.Println("use forward")
	} else {
		fmt.Println("use backward")
	}

	reality, _ := calculateBackwardDiffInterpolation(input_x, input_y, input_xu)
	var expected float32 = 96.8368
	fmt.Println("Reality: ", reality)
	fmt.Println("Expected: ", expected)

	fmt.Println("Expected:", expected)
	fmt.Println("Reality:", reality)
	if expected != reality {
		t.Fatal()

	}

}

func TestRoundingTillFive(t *testing.T) {
	input := 2.334345873498
	output := roundTilFive(input)
	expected := 2.334346

	if output != float32(expected) {
		fmt.Print(output)
		t.Fatal("Not same")
	}
}

func TestVProduct(t *testing.T) {
	output := calcVProduct(-0.6, 1)
	fmt.Println("output: ", output)
}
