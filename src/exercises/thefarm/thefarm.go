package thefarm

import (
	"errors"
	"fmt"
)

type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}

func DivideFood(f FodderCalculator, cows int) (float64, error) {
	amount, err := f.FodderAmount(cows)
	if err != nil {
		return 0.0, err
	}

	factor, err := f.FatteningFactor()
	if err != nil {
		return 0.0, err
	}

	return amount / float64(cows) * factor, nil
}

func ValidateInputAndDivideFood(f FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0.0, errors.New("invalid number of cows")
	}

	return DivideFood(f, cows)
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

func ValidateNumberOfCows(cows int) error {
	if cows == 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "no cows don't need food",
		}
	}

	if cows < 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "there are no negative cows",
		}
	}

	return nil
}
