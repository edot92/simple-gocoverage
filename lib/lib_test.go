package lib

import (
	"log"
	"testing"
)

type Test struct {
	inParamDeret []int
	outCalc      ResultLib
}

var tests = []Test{
	{
		[]int{1, 2, 3},
		ResultLib{3,
			[]int{1, 2, 3, 8, 9, 3, 2, 1},
		},
	},
	{
		[]int{1, 2}, ResultLib{
			2, []int{1, 2, 1, 2, 2, 1},
		},
	},
	{
		[]int{7, 1, 2}, ResultLib{
			2, []int{7, 1, 2, 9, 7, 2, 1},
		},
	},
}

func TestCariDeretMax(t *testing.T) {
	for _, test := range tests {
		switch len(test.inParamDeret) {
		case 3:
			{
				if test.inParamDeret[0] == 7 {
					t.Skip("Run Case Three")
				}
				t.Skip("Run Case One")
			}
		case 2:
			{
				t.Skip("Run Case Two")
			}
		}
	}
}

func Test_caseOne(t *testing.T) {
	const indexArrTest = 0
	returnCase, err := CariDeretMax(tests[indexArrTest].inParamDeret)
	if err != nil {
		log.Fatal(err)
	}
	if returnCase.ValBigger != tests[indexArrTest].outCalc.ValBigger {
		t.Errorf(": returnCase(%d)=%s; want %s", tests[indexArrTest].inParamDeret, returnCase, tests[indexArrTest].outCalc.ValBigger)
	}
}

func Test_caseTwo(t *testing.T) {
	const indexArrTest = 1
	returnCase, _ := CariDeretMax(tests[indexArrTest].inParamDeret)
	if returnCase.ValBigger != tests[indexArrTest].outCalc.ValBigger {
		t.Errorf(": returnCase(%d)=%s; want %s", tests[indexArrTest].inParamDeret, returnCase, tests[indexArrTest].outCalc.ValBigger)
	}
}
func Test_caseThree(t *testing.T) {
	const indexArrTest = 2
	returnCase, _ := CariDeretMax(tests[indexArrTest].inParamDeret)
	if returnCase.ValBigger != tests[indexArrTest].outCalc.ValBigger {
		t.Errorf(": returnCase(%d)=%s; want %s", tests[indexArrTest].inParamDeret, returnCase, tests[indexArrTest].outCalc.ValBigger)
	}
}
