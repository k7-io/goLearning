package main

import "testing"

func TestAdd(t *testing.T) {
	inputs := [][]int{{1, 2}, {2, 3}, {4, 7}}
	expected := []int{3, 5, 11}
	for i := 0; i < len(inputs); i++ {
		ret := Add(inputs[i][0], inputs[i][1])
		if ret != expected[i] {
			t.Errorf("input is %d, the expected is %d, the actual is %d", inputs[i], expected[i], ret)
		}
	}
}

func TestArea(t *testing.T) {
	inputs := []float64{1, 2, 4}
	expected := []float64{3.14, 12.56, 50.24}
	for i := 0; i < len(inputs); i++ {
		ret := Area(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %f, the expected is %f, the actual is %f", inputs[i], expected[i], ret)
		}
	}
}
