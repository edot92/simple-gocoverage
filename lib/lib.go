package lib

import "errors"

// ResultLib strcut data CariDeretMax
// ValBigger  val bigger from index array
// ValComputeArray arr format
type ResultLib struct {
	ValBigger       int
	ValComputeArray []int
}

// CariDeretMax lenght of array must 2 or 3
// if lenght = 3 n1 of array != 7 then run caseOne
// if lenght = 3 and n1 of array == 7  then run CaseThree
// if lenght = 2 then run caseTwo
func CariDeretMax(paramDeret []int) (ResultLib, error) {
	var resErr ResultLib
	switch len(paramDeret) {
	case 3:
		{
			if paramDeret[0] == 7 {
				return caseThree(paramDeret)
			}
			return caseOne(paramDeret)
		}
	case 2:
		{
			return caseTwo(paramDeret)
		}
	}
	return resErr, errors.New("wrong format, lenght index must 3 and 2")
}

// caseOne contoh : [1, 2, 3, 8, 9, 3, 2, 1] nilai terbesarnya 3 (dari deret [1, 2, 3])
// nilai deret = [1,2,3]
// type array index n = 3 (CaseOne)
// nilai terbesar deret = 3
// size array n+1= nilai terbesar deret +n3 + n2
// size array n+2= (value size array n+1) + n1
func caseOne(paramDeret []int) (ResultLib, error) {
	var result ResultLib
	valBigFromArr := 0
	for _, e := range paramDeret {
		if e > valBigFromArr {
			valBigFromArr = e
		}
	}
	sizeArrN1 := valBigFromArr + paramDeret[2] + paramDeret[1]
	sizeArrN2 := sizeArrN1 + paramDeret[0]
	result.ValBigger = valBigFromArr
	result.ValComputeArray = append(paramDeret, sizeArrN1)
	result.ValComputeArray = append(result.ValComputeArray, sizeArrN2)
	result.ValComputeArray = append(result.ValComputeArray, paramDeret[2], paramDeret[1], paramDeret[0])
	return result, nil
}

// caseTwo contoh :[1, 2, 1, 2, 2, 1] → 2
// nilai deret = [1, 2]
// type array index n = 2 (CaseTwo)
// nilai terbesar deret = 2
// size array n+1= n1
// size array n+2= n2
// result [n1,n2,n1,n2,n2,n1]
func caseTwo(paramDeret []int) (ResultLib, error) {
	var result ResultLib
	valBigFromArr := 0
	for _, e := range paramDeret {
		if e > valBigFromArr {
			valBigFromArr = e
		}
	}
	result.ValBigger = valBigFromArr
	result.ValComputeArray = append(paramDeret, paramDeret[0], paramDeret[1])             // append n1,n2
	result.ValComputeArray = append(result.ValComputeArray, paramDeret[1], paramDeret[0]) //append swap byte
	return result, nil

}

// CaseThree contoh : [7, 1, 2, 9, 7, 2, 1] → 2
// nilai deret = [1, 2]
// type array index n = 3 dan nilai n1 =7 (CaseThree)
// nilai terbesar deret = 2 ,abaikan index n1
// size array n+1= n1+ n3
// swap byte = n1,n3,n4
// result [nilai deret,size array n+1 ,n2,swap byte]
func caseThree(paramDeret []int) (ResultLib, error) {
	var result ResultLib
	valBigFromArr := 0
	for i, e := range paramDeret {
		if e > valBigFromArr {
			if i != 0 { //ignore if fisrt n
				valBigFromArr = e
			}
		}
	}
	result.ValBigger = valBigFromArr
	result.ValComputeArray = append(paramDeret, paramDeret[0]+paramDeret[2])              //n1+n3
	result.ValComputeArray = append(result.ValComputeArray, paramDeret[1], paramDeret[0]) //swap byte
	return result, nil
}
