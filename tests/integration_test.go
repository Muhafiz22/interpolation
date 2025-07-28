package tests_test

import (
	"fmt"
	"numanalysis/functions"
	"testing"
)

func TestImpsonsOneThirdMethod(t *testing.T) {

	var a float32 = 4.0
	var b float32 = 5.2
	var n float32 = 6.0

	var y []float32 = []float32{0.60206, 0.62325, 0.64345, 0.66276, 0.68124, 0.69897, 0.71600}

	var reality float32 = functions.SimpsonsOneThirdMethod(a, b, n, y)
	var expected float32 = 0.74989
	fmt.Println(expected)
	fmt.Println(reality)
	if expected != reality {
		t.Fatal()
	}
}
