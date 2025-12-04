package day1

import (
	"fmt"
	"testing"
)

func TestPass1Sample(t *testing.T) {
	want := 3

	var pass = getPass1("./data/example")
	if pass != want {
		t.Errorf("Expected %d but got %d", want, pass)
	}
	fmt.Println("Pass1 sample: ", pass)
}

func TestPass1Input(t *testing.T) {

	var pass = getPass1("./data/input")
	fmt.Println("Pass1 input:", pass)
}

func TestPass2Sample(t *testing.T) {
	want := 6

	var pass = getPass2("./data/example")
	if pass != want {
		t.Errorf("Expected %d but got %d", want, pass)
	}
	fmt.Println("Pass2 sample: ", pass)
}

func TestPass2Input(t *testing.T) {

	var pass = getPass2("./data/input")
	fmt.Println("Pass2:", pass)
}
