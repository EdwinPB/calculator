package models

import (
	"errors"
	"strconv"
	"strings"
)

type Calculator struct {
	Numbers string `json:"s_numbers" db:""`
}

// Calculators is not required by pop and may be deleted
type Calculators []Calculator

func Calculate(sNumbers string) (int, error) {
	if sNumbers == "" {
		return 0, nil
	}

	delimiter := ","

	if strings.HasPrefix(sNumbers, "//") {
		delimiter = sNumbers[2:3]
	}

	if strings.Contains(sNumbers, "\n") {
		sNumbers = strings.ReplaceAll(sNumbers, "\n", delimiter)
	}

	if strings.Contains(sNumbers, delimiter) {
		sNums := strings.Split(sNumbers, delimiter)
		result := 0
		for _, sNum := range sNums {
			num, _ := strconv.Atoi(sNum)
			if num < 0 {
				return 0, errors.New("Invalid Input")
			}
			result += num
		}
		return result, nil
	}

	num, _ := strconv.Atoi(sNumbers)
	if num < 0 {
		return 0, errors.New("Invalid Input")
	}

	return num, nil
}
