package bowling

import (
	"fmt"
	"testing"
)

func scoreChecker(input []Frame, expectedScore int, expectedError error) error {
	score, err := GetScore(input)

	if err != expectedError && !(err != nil && expectedError != nil && err.Error() == expectedError.Error()) {
		return fmt.Errorf("Score error : %+v, expected %+v", err, expectedError)
	}
	if score != expectedScore {
		return fmt.Errorf("Score : %d, expected %d", score, expectedScore)
	}
	return nil
}

func TestNullScore(t *testing.T) {
	input := []Frame{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	expected := 0
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestScore(t *testing.T) {
	input := []Frame{{7, 2}, {5, 1}, {9, 0}, {9, 0}, {9, 0}, {9, 0}, {9, 0}, {9, 0}, {9, 0}, {9, 0}}
	expected := 87
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestErreurLancers(t *testing.T) {
	input := []Frame{{7, 5}, {13, 2}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}}
	expected := 0
	expectedError := fmt.Errorf("A throw give a score > 10")
	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestSpare(t *testing.T) {
	input := []Frame{{7,3}, {2,0}, {9,0}, {9, 0}, {9, 0}, {9, 0}, {9, 0}, {9, 0}, {9, 0}, {9, 0}}
	expected := 86
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestStrike(t *testing.T) {
	input := []Frame{{7,3}, {2,0}, {10, 0}, {0, 9}, {9, 0}, {9, 0}, {9, 0}, {9, 0}, {9, 0}, {9, 0}}
	expected := 96
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestThrowNotNegative(t *testing.T) {
	input := []Frame{{7, 2}, {10, -2}, {9, 0}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}}
	expected := 0
	expectedError := fmt.Errorf("A throw is negative")
	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestLength(t *testing.T) {
	input := []Frame{{7, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}}
	expected := 0
	expectedError := fmt.Errorf("There is more or less than 10 throws")
	if err := scoreChecker(input, expected, expectedError); err != nil {
		t.Fatalf("%+v\n", err)
	}
}
