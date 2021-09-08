package models

type CurrencyConversionReqest struct {
	Money    string
	Currency string
}

type CurrencyConversionResponse struct {
	ConvertedMoney string
}
