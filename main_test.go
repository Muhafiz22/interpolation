package main

import (
	"fmt"
	"testing"
)

func TestCheckForward(t *testing.T) {

	var input_x []float32 = []float32{1.0, 2.0, 3.0, 4.0}
	var input_y []float32 = []float32{2.0, 4.0, 8.0, 16.0}

	var input_xu float32 = 2.5

	/*var input_x []float32 = []float32{1891, 1901, 1911, 1921, 1931}
	var input_y []float32 = []float32{46, 66, 81, 93, 101}

	var input_xu float32 = 1925.0*/

	c, result, deltaY := checkForward(input_x, input_y, input_xu)

	if c {

		fmt.Println("Using Forward Difference Interpolation")
	} else {

		fmt.Println("Using Backward Difference Interpolation")
	}

	fmt.Println(result)
	fmt.Println(deltaY)
}

/*func TestCalculateForwardDiffInterpolation(t *testing.T) {

	var input_x []float32 = []float32{1.0, 2.0, 3.0, 4.0}
	var input_y []float32 = []float32{2.0, 4.0, 8.0, 16.0}

	var input_xu float32 = 2.5

	c := checkForward(input_x, input_xu)
	if c {
		fmt.Println("use forward")
	} else {
		fmt.Println("use backward")
	}

	table, reality := calculateForwardDiffInterpolation(input_x, input_y, input_xu)
	var expected float32 = 5.625

	fmt.Println(table)
	fmt.Println("Expected:", expected)
	fmt.Println("Reality:", reality)
	if expected != reality {
		t.Fatal()

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

	table, reality := calculateBackwardDiffInterpolation(input_x, input_y, input_xu)
	var expected float32 = 96.8368

	fmt.Println(table)
	fmt.Println("Expected:", expected)
	fmt.Println("Reality:", reality)
	if expected != reality {
		t.Fatal()

	}
}*/
