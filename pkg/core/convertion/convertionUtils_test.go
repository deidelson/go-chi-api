package convertion

import "testing"

func TestStringToIntHappyPath(t *testing.T) {
	number, err := StringToInt("1")

	if err != nil {
		t.Error("unexpected error")

	}

	if number != 1 {
		t.Error("number must be 1 and its ", number)
	}
}

func TestStringToIntHappyError(t *testing.T) {
	_, err := StringToInt("a")

	if err == nil {
		t.Error("there is not an error")
	}

}
