package convertion

import (
	"errors"
	"strconv"
)

func StringToInt(numericString string) (int, error) {
	number, err := strconv.Atoi(numericString)

	if err != nil {
		return 0, errors.New("convertion.string.to.int")
	}
	return number, nil

}
