package aoc01

import (
	"os"
	"fmt"
	"testing"
	"strings"
)

func TestTotalFuel(t *testing.T) {
	type testcase struct {
		fuel      int
		totalFuel int
	}

	testcases := []testcase{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, tt := range testcases {
		tt := tt
		t.Run(fmt.Sprintf("fuel=%d", tt.fuel), func(t *testing.T) {
			t.Parallel()
			actual := totalFuel(tt.fuel)
			if tt.totalFuel != actual {
				t.Errorf("totalFuel(%d): expected=%d, actual=%d", tt.fuel, tt.totalFuel, actual)
			}
		})
	}
}

func TestTotalFuelForFile(t *testing.T) {
	input := "14\n1969\n100756\n"

	r := strings.NewReader(input)
	actual, err := totalFuelForFile(r)
	if err != nil {
		t.Error(err)
	}

	if actual != 2 + 966 + 50346 {
		t.Fail()
	}
}

func TestTotalFuelForReal(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Skip()
	}

	out, _ := totalFuelForFile(f)
	t.Log("ANSWER:", out)
}