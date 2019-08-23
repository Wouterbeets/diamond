package main

import "testing"

func Test_diamond_A(t *testing.T) {
	in := 'A'
	expected := []string{"A"}
	out := diamond(in)
	if len(out) == 0 {
		t.Errorf("out is empty")
	}
	for i := range out {
		if expected[i] != out[i] {
			t.Errorf("out != expected \nout: %+v, expected: %+v", out, expected)
		}
	}
}
