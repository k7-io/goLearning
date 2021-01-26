package main

import "testing"

func TestLevel(t *testing.T) {
	inputs := []int{20, 65, 75}
	expected := []string{"not pass", "pass", "good"}
	for i := 0; i < len(inputs); i++ {
		ret := Level(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d, the expected is %s, the actual is %s", inputs[i], expected[i], ret)
		}
	}
}
