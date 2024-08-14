package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {

	r := make([]float32, 5)
	s := make([]float32, 5)
	s = append(s, 20, 30, 40, 50, 60)
	r = append(r, 2, 3, 4, 5, 6)
	err, _, _ := validateInput(r, s)

	if err != nil {
		t.Fatal(err.Error())
	}

}

func TestCheckForward(t *testing.T) {

	r := make([]float32, 5)
	r = append(r, 2, 3, 4, 5, 6)
	var p float32 = 1.5
	c := checkForward(r, p)

	if c {
		fmt.Print("use forward")
	} else {
		fmt.Print("use backward")
	}
}
