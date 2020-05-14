package iseven

import "testing"

func TestIsEven(t *testing.T) {
	cases := []struct {
		number int
	}{
		{number: 0},
		{number: 1},
		{number: 2},
		{number: 3},
		{number: 4},
		{number: 5},
	}
	for _, c := range cases {
		if IsEven(c.number) == false {
			t.Error("The number is odd")
		}
	}
}
