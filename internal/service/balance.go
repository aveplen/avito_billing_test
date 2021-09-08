package service

import (
	"github.com/aveplen/avito_test/internal/models"
	"github.com/aveplen/avito_test/internal/repository"
)

type BalanceService struct {
	balanceRepo repository.Balance
}

func NewBalanceService(repo repository.Balance) *BalanceService {
	return &BalanceService{
		balanceRepo: repo,
	}
}

func (b *BalanceService) Deposit(deposit models.Deposit) error {
	exists, err := b.balanceRepo.CheckExists(deposit.UserID)
	if err != nil {
		return err
	}
	if exists {
		return b.balanceRepo.DepositUpdate(deposit)
	}
	return b.balanceRepo.DepositInsert(deposit)
}

func (b *BalanceService) Withdraw(withdrawal models.Withdrawal) error {
	return b.balanceRepo.Withdraw(withdrawal)
}

func (b *BalanceService) Transact(transaction models.Transaction) error {
	exists, err := b.balanceRepo.CheckExists(transaction.ToID)
	if err != nil {
		return err
	}
	if exists {
		return b.balanceRepo.Transact(transaction)
	}
	if err := b.balanceRepo.DepositInsert(models.Deposit{
		UserID: transaction.ToID,
		Money:  "0",
	}); err != nil {
		return err
	}
	return b.balanceRepo.Transact(transaction)
}

func (b *BalanceService) Balance(balanceRequest models.BalanceRequest) (models.BalanceResponse, error) {
	return b.balanceRepo.Balance(balanceRequest)
}
