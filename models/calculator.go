package models

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

type Calculator struct {
	ID              uuid.UUID `json:"id" db:"id"`
	EnteredValue    string    `json:"entered_value" db:"entered_value"`
	CalculatedValue string    `json:"calculated_value" db:"calculated_value"`

	UserID uuid.UUID `json:"user_id" db:"user_id"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Calculators is not required by pop and may be deleted.
type Calculators []Calculator

var ErrInvalidValue = errors.New("invalid input")

func Calculate(sNumbers string) (int, error) {
	if sNumbers == "" {
		return 0, nil
	}

	delimiter := ","
	if strings.HasPrefix(sNumbers, "//") {
		delimiter = sNumbers[2:3]
		sNumbers = sNumbers[4:]
	}

	if strings.Contains(sNumbers, "\n") {
		sNumbers = strings.ReplaceAll(sNumbers, "\n", delimiter)
	}

	if strings.Contains(sNumbers, delimiter) {
		sNums := strings.Split(sNumbers, delimiter)
		result := 0
		for _, sNum := range sNums {
			num, err := strconv.Atoi(sNum)
			if num < 0 || err != nil {
				return 0, ErrInvalidValue
			}
			result += num
		}
		return result, nil
	}

	num, err := strconv.Atoi(sNumbers)
	if num < 0 || err != nil {
		return 0, ErrInvalidValue
	}

	return num, nil
}

func (c *Calculator) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: c.UserID, Name: "UserID", Message: "User does not exist"},
	), nil
}

func (cs Calculators) MaxCalculatorsValue() int {
	result := 0
	for _, c := range cs {
		iNum, _ := strconv.Atoi(c.CalculatedValue)
		if iNum > result {
			result = iNum
		}
	}
	return result
}

func (cs Calculators) MinCalculatorsValue() int {
	if len(cs) == 0 {
		return 0
	}
	result, _ := strconv.Atoi(cs[0].CalculatedValue)
	for _, c := range cs {
		iNum, _ := strconv.Atoi(c.CalculatedValue)
		if iNum < result {
			result = iNum
		}
	}
	return result
}

func (cs Calculators) AverageCalculatedValues() float64 {
	if len(cs) == 0 {
		return 0
	}
	sum := 0
	for _, c := range cs {
		num, _ := strconv.Atoi(c.CalculatedValue)
		sum += num
	}
	return float64(sum) / float64(len(cs))
}
