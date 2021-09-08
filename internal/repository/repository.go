package repository

import (
	"github.com/aveplen/avito_test/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	Balance interface {
		CheckExists(int64) (bool, error)
		DepositInsert(models.Deposit) error
		DepositUpdate(models.Deposit) error
		Withdraw(models.Withdrawal) error
		Transact(models.Transaction) error
		Balance(models.BalanceRequest) (models.BalanceResponse, error)
	}

	Repository struct {
		Balance
	}
)

func NewRepository(dbpool *pgxpool.Pool) *Repository {
	return &Repository{
		Balance: NewUserBalance(dbpool),
	}
}
