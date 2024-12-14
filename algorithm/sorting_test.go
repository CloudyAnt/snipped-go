package algorithm

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

type Case struct {
	input  []int
	output []int
}

func createCases() []Case {
	case1 := Case{
		[]int{-2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]int{-2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	case2 := Case{
		[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2},
		[]int{-2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	case3 := Case{
		[]int{21, 39, -72, 18, 32, -10, 98, 12, 56, 28, -1, 10},
		[]int{-72, -10, -1, 10, 12, 18, 21, 28, 32, 39, 56, 98},
	}
	cases := []Case{case1, case2, case3}
	return cases
}

func compareArray(a1, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func TestSorting(t *testing.T) {
	//testSorting(t, selection)
	//testSorting(t, bubble)
	//testSorting(t, insertion)
	//testSorting(t, qs)
	testSorting(t, heap)
}

func testSorting(t *testing.T, f func([]int)) {
	// GIVEN
	cases := createCases()

	funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
	for _, c := range cases {
		t.Run(fmt.Sprintf("%s:%v\n", funcName, c.input), func(t *testing.T) {
			// WHEN
			f(c.input)

			// THEN
			if !compareArray(c.output, c.input) {
				t.Errorf("expected:%v , actual:%v", c.output, c.input)
			}
		})
	}
}
