package handler

import (
	"net/http"
	"regexp"

	"github.com/aveplen/avito_test/internal/consts"
	"github.com/aveplen/avito_test/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

/*
	парсит код ошибки из базы

	нереальный костыль, но как
	сделать адекватно я чет не нашёл : )
*/
func ErrorCode(err error) string {
	r := regexp.MustCompile(`SQLSTATE (.{5})`)
	return r.FindStringSubmatch(err.Error())[1]
}

func (h *Handler) Deposit(c *gin.Context) {
	validDeposit, ok := c.Get("validDeposit")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.logger.Warn(consts.ErrValidDepositDoesntExistInContext.Error())
		return
	}
	deposit, ok := validDeposit.(models.Deposit)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.logger.Warn(consts.ErrDepositTypeAssertFailed.Error())
		return
	}
	if err := h.service.BalanceService.Deposit(deposit); err != nil {
		h.logger.Warn(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) Withdraw(c *gin.Context) {
	validWithdrawal, ok := c.Get("validWithdrawal")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.logger.Warn(consts.ErrValidWithdrawalDoesntExistInContext.Error())
		return
	}
	withdrawal, ok := validWithdrawal.(models.Withdrawal)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.logger.Warn(consts.ErrWithdrawalTypeAssertFailed.Error())
		return
	}
	if err := h.service.BalanceService.Withdraw(withdrawal); err != nil {
		h.logger.Warn(err.Error())
		pgerr := ErrorCode(err)
		switch pgerr {
		case
			pgerrcode.CheckViolation,
			pgerrcode.NotNullViolation:
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"err": "requested withdrawal could not be completed because user's balance is too low",
			})
		default:
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) Transact(c *gin.Context) {
	validTransaction, ok := c.Get("validTransaction")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.logger.Warn(consts.ErrValidTransactionDoesntExistInContext.Error())
		return
	}
	transaction, ok := validTransaction.(models.Transaction)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.logger.Warn(consts.ErrTransactionTypeAssertFailed.Error())
		return
	}
	if err := h.service.BalanceService.Transact(transaction); err != nil {
		h.logger.Warn(err.Error())
		pgerr := ErrorCode(err)
		switch pgerr {
		case
			pgerrcode.CheckViolation,
			pgerrcode.NotNullViolation:
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
				"err": "requested transaction could not be completed because user's balance is too low",
			})
		default:
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) Balance(c *gin.Context) {
	validBalanceRequest, ok := c.Get("validBalanceRequest")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.logger.Warn(consts.ErrValidBalanceDoesntExistInContext.Error())
		return
	}
	balanceRequest, ok := validBalanceRequest.(models.BalanceRequest)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.logger.Warn(consts.ErrBalanceTypeAssertFailed.Error())
		return
	}
	balanceResponse, err := h.service.BalanceService.Balance(balanceRequest)
	if err != nil {
		h.logger.Warn(err.Error())
		if errors.Is(err, pgx.ErrNoRows) {
			c.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
				"err": "user with given id doesn't exist",
			})
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	validCurrency, ok := c.Get("currency")
	if !ok {
		c.JSON(http.StatusOK, balanceResponse)
		return
	}
	cur, ok := validCurrency.(string)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		h.logger.Warn(consts.ErrBalanceTypeAssertFailed.Error())
		return
	}
	conversionRequest := models.CurrencyConversionReqest{
		Money:    balanceResponse.Money,
		Currency: cur,
	}
	conversionResponse, err := h.service.CurrencyApiService.CurrencyConversion(conversionRequest)
	if err != nil {
		h.logger.Warn(err.Error())
		switch {
		case errors.Is(err, consts.ErrApiBadStatusCode):
			{
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
					"err": "remote server returned error",
				})
			}
		case errors.Is(err, consts.ErrApiEmptyResponseBody):
			{
				c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
					"err": "remote server dindn't return anything, check your currency request",
				})
			}
		case errors.Is(err, consts.ErrApiResponseNoMatch):
			{
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
					"err": "remote server returned weird response",
				})
			}
		}
		return
	}
	balanceResponse.Money = conversionResponse.ConvertedMoney
	c.JSON(http.StatusOK, balanceResponse)
}
