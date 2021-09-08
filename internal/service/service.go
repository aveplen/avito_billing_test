package service

import (
	"github.com/aveplen/avito_test/internal/config"
	"github.com/aveplen/avito_test/internal/repository"
)

type Service struct {
	BalanceService     *BalanceService
	CurrencyApiService *CurrencyApiClient
}

func NewService(cfg *config.Config, repo *repository.Repository) *Service {
	return &Service{
		BalanceService:     NewBalanceService(repo.Balance),
		CurrencyApiService: NewCurrencyApiClient(cfg.CurrencyApi),
	}
}
