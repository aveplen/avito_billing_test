package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/aveplen/avito_test/internal/config"
	"github.com/aveplen/avito_test/internal/consts"
	"github.com/aveplen/avito_test/internal/models"
	"github.com/aveplen/avito_test/internal/utils"
)

var moneyRegex *regexp.Regexp

type CurrencyApiClient struct {
	GenerateUrl func(string) string
}

func NewCurrencyApiClient(cfg config.CurrencyApi) *CurrencyApiClient {
	moneyRegex = regexp.MustCompile(`[0-9]+(?:\.[0-9]*)?`)
	return &CurrencyApiClient{
		GenerateUrl: func(targetCurrency string) string {
			return fmt.Sprintf("https://free.currconv.com/api/v7/convert?q=RUB_%s&compact=ultra&apiKey=%s",
				targetCurrency,
				cfg.Key)
		},
	}
}

func (api *CurrencyApiClient) CurrencyConversion(req models.CurrencyConversionReqest) (models.CurrencyConversionResponse, error) {
	res := models.CurrencyConversionResponse{}
	apiResponse, err := http.Get(api.GenerateUrl(req.Currency))
	if err != nil {
		return res, fmt.Errorf("currency conversion: could not get response from api: %w", err)
	}
	status := apiResponse.StatusCode
	if status != 200 {
		return res, fmt.Errorf("currency conversion: api returned %d: %w", status, consts.ErrApiBadStatusCode)
	}
	body, err := ioutil.ReadAll(apiResponse.Body)
	if err != nil {
		return res, fmt.Errorf("currency conversion: could not read response body: %w", err)
	}
	if len(body) == 0 {
		return res, fmt.Errorf("currency conversion: %w", consts.ErrApiEmptyResponseBody)
	}
	if len(body) == 2 && body[0] == '{' && body[1] == '}' {
		return res, fmt.Errorf("currency conversion: %w", consts.ErrApiEmptyResponseBody)
	}
	responseMoney := moneyRegex.Find(body)
	if len(responseMoney) == 0 {
		return res, fmt.Errorf("currencty conversion: %w", consts.ErrApiResponseNoMatch)
	}
	res.ConvertedMoney = string(
		utils.StringFloor(
			utils.StringMultiplication(
				[]byte(req.Money),
				responseMoney),
			2))
	return res, nil
}
