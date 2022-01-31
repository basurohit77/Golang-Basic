package utils

import "testing"

func Test_97_IsPrime(t *testing.T) {
	//arrange
	var no = 97
	var expectedResult = true

	//act
	result := IsPrime(no)

	//assert
	if result != expectedResult {
		t.Errorf("Expected %v but got %v", expectedResult, result)
	}
}

func Test_96_IsPrime(t *testing.T) {
	//arrange
	var no = 96
	var expectedResult = false

	//act
	result := IsPrime(no)

	//assert
	if result != expectedResult {
		t.Errorf("Expected %v but got %v", expectedResult, result)
	}
}

func Test_95_IsPrime(t *testing.T) {
	//arrange
	var no = 95
	var expectedResult = false

	//act
	result := IsPrime(no)

	//assert
	if result != expectedResult {
		t.Errorf("Expected %v but got %v", expectedResult, result)
	}
}
