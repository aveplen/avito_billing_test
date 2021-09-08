package handler

import (
	"net/http"
	"regexp"

	"github.com/aveplen/avito_test/internal/consts"
	"github.com/aveplen/avito_test/internal/models"
	"github.com/gin-gonic/gin"
)

var (
	compiledMoneyRegex    *regexp.Regexp
	compiledCurrencyRegex *regexp.Regexp
)

func CheckMoney(money string) bool {
	if compiledMoneyRegex == nil {
		compiledMoneyRegex = regexp.MustCompile(`^(?:0|(?:[1-9][0-9]*))(?:\.[0-9]{1,2})?$`)
	}
	return compiledMoneyRegex.MatchString(money)
}

func CheckCurrency(currency string) bool {
	if compiledCurrencyRegex == nil {
		compiledCurrencyRegex = regexp.MustCompile(`^[A-Z]{3}$`)
	}
	return compiledCurrencyRegex.MatchString(currency)
}

func CheckID(id int64) bool {
	return id >= 1
}

func (h *Handler) ValidateDeposit(c *gin.Context) {
	var deposit models.Deposit
	if err := c.BindJSON(&deposit); err != nil {
		h.logger.Warn(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"err": "valid deposit request must consist of user_id::number and money::string",
		})
		return
	}
	if !CheckID(deposit.UserID) {
		h.logger.Warn(consts.ErrNonPositiveRequestID.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"err": "user_id must be valid positive int64",
		})
		return
	}
	match := CheckMoney(deposit.Money)
	if !match {
		h.logger.Warn(consts.ErrMoneyDoesntMatchRegex.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"err": "money must not contain leading zeros, trailing delimeter or fractional part longer then 2 digits",
		})
		return
	}
	c.Set("validDeposit", deposit)
	c.Next()
}

func (h *Handler) ValidateWithdrawal(c *gin.Context) {
	var withdrawal models.Withdrawal
	if err := c.BindJSON(&withdrawal); err != nil {
		h.logger.Warn(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"err": "valid withdrawal request must consist of user_id::number and money::string",
		})
		return
	}
	if !CheckID(withdrawal.UserID) {
		h.logger.Warn(consts.ErrNonPositiveRequestID.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"err": "user_id must be valid positive int64",
		})
		return
	}
	match := CheckMoney(withdrawal.Money)
	if !match {
		h.logger.Warn(consts.ErrMoneyDoesntMatchRegex.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"err": "money must not contain leading zeros, trailing delimeter or fractional part longer then 2 digits",
		})
		return
	}
	c.Set("validWithdrawal", withdrawal)
	c.Next()
}

func (h *Handler) ValidateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.BindJSON(&transaction); err != nil {
		h.logger.Warn(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"err": "valid transaction request must consist of from_id::number, to_id::number and money::string",
		})
		return
	}
	if !CheckID(transaction.FromID) {
		h.logger.Warn(consts.ErrNonPositiveRequestID.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"err": "from_id must be valid positive int64",
		})
		return
	}
	if !CheckID(transaction.ToID) {
		h.logger.Warn(consts.ErrNonPositiveRequestID.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"err": "to_id must be valid positive int64",
		})
		return
	}
	match := CheckMoney(transaction.Money)
	if !match {
		h.logger.Warn(consts.ErrMoneyDoesntMatchRegex.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"err": "money must not contain leading zeros, trailing delimeter or fractional part longer then 2 digits",
		})
		return
	}
	c.Set("validTransaction", transaction)
	c.Next()
}

func (h *Handler) ValidateBalanceRequest(c *gin.Context) {
	var balance models.BalanceRequest
	if err := c.BindJSON(&balance); err != nil {
		h.logger.Warn(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"err": "valid balance request must consist of user_id::number, to_id::number and money::string",
		})
		return
	}
	if !CheckID(balance.UserID) {
		h.logger.Warn(consts.ErrNonPositiveRequestID.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"err": "user_id must be valid positive int64",
		})
		return
	}
	c.Set("validBalanceRequest", balance)
	c.Next()
}

func (h *Handler) ValidateCurrency(c *gin.Context) {
	cur := c.Query("currency")
	if cur == "" {
		c.Next()
		return
	}
	if !CheckCurrency(cur) {
		h.logger.Warn(consts.ErrInvalidCurrency.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"err": "currency could only be represented by 3 capital letters",
		})
	}
	c.Set("currency", cur)
	c.Next()
}
