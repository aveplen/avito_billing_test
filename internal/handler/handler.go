package handler

import (
	"github.com/aveplen/avito_test/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	service *service.Service
	logger  *zap.Logger
}

func NewHandler(service *service.Service, logger *zap.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) RouteBalance(route *gin.Engine) {
	billing := route.Group("/billing")
	{
		billing.POST("/balance",
			h.ValidateBalanceRequest,
			h.ValidateCurrency,
			h.Balance)

		billing.POST("/deposit",
			h.ValidateDeposit,
			h.Deposit)

		billing.POST("/withdraw",
			h.ValidateWithdrawal,
			h.Withdraw)

		billing.POST("/transact",
			h.ValidateTransaction,
			h.Transact)
	}
}
