package util

import (
	"strconv"
)

func IntToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func StringToInt(str string) (*int, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return nil, err
	}

	return &i, nil
}
