package consts

import (
	"errors"
)

var (
	// для всех роутов
	ErrNonPositiveRequestID = errors.New("non positive request id")
	// (ну почти)
	ErrMoneyDoesntMatchRegex = errors.New("money doesn't match regex")

	// толкьо для зачисления
	ErrValidDepositDoesntExistInContext = errors.New("valid deposit doesn't exist in context")
	ErrDepositTypeAssertFailed          = errors.New("deposit type assertion failed")

	// для снятия
	ErrValidWithdrawalDoesntExistInContext = errors.New("valid withdrawal doesn't exist in context")
	ErrWithdrawalTypeAssertFailed          = errors.New("withdrawal type assertion failed")

	// для переводов
	ErrValidTransactionDoesntExistInContext = errors.New("valid transaction doesn't exist in context")
	ErrTransactionTypeAssertFailed          = errors.New("transaction type assertion failed")

	// для получения баланса
	ErrValidBalanceDoesntExistInContext = errors.New("valid balance doesn't exist in context")
	ErrBalanceTypeAssertFailed          = errors.New("balance type assertion failed")

	// для получения баланса с конвертацией
	ErrInvalidCurrency                   = errors.New("invalid currency")
	ErrValidCurrencyDoesntExistInContext = errors.New("valid currency doesn't exist in context")
	ErrCurencyTypeAssertFailed           = errors.New("currency type assertion failed")

	// для получения ошибки от API
	ErrApiBadStatusCode     = errors.New("api returned bad status code")
	ErrApiResponseNoMatch   = errors.New("no match found in api response")
	ErrApiEmptyResponseBody = errors.New("api returned empty response body")
)
