package domain

import "time"

type Account struct {
	ID        int64
	OwnerName string
	Balance   float64
	Currency  string
	Activated bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AccountValidationErrMsg string

const (
	owner_name_empty_err             = "owner name cannot be empty"
	balance_less_than_zero_value_err = "balance cannot be less than zero"
	currency_empty_err               = "currency cannot be empty"
)

type OwnerNameEmptyError struct {
	msg string
}

func (e *OwnerNameEmptyError) Error() string {
	return e.msg
}

type BalanceLessThanZeroValueError struct {
	msg string
}

func (e *BalanceLessThanZeroValueError) Error() string {
	return e.msg
}

type CurrencyEmptyError struct {
	msg string
}

func (e *CurrencyEmptyError) Error() string {
	return e.msg
}

type AccountValidationError []error

/*
	@ Functionality :
		=> validate that owner name is not empty
		=> validate that currency type is not empty
		=> validate that balance is not less than zero
	@ Return values :
		=> returns ([]error, true) if there is any error
		=> returns (nil, false) if there is no error
*/
func (acc *Account) Validate() (AccountValidationError, bool) {
	var validationErrs AccountValidationError
	if len(acc.OwnerName) == 0 {
		validationErrs = append(validationErrs, &OwnerNameEmptyError{msg: owner_name_empty_err})
	}
	if len(acc.Currency) == 0 {
		validationErrs = append(validationErrs, &CurrencyEmptyError{msg: currency_empty_err})
	}
	if acc.Balance == 0 {
		validationErrs = append(validationErrs, &BalanceLessThanZeroValueError{msg: balance_less_than_zero_value_err})
	}

	if len(validationErrs) == 0 {
		return nil, true
	} else {
		return validationErrs, false
	}
}
