package repository

import (
	"context"
	"fmt"

	"github.com/aveplen/avito_test/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserBalance struct {
	dbpool *pgxpool.Pool
}

func NewUserBalance(dbpool *pgxpool.Pool) *UserBalance {
	return &UserBalance{
		dbpool: dbpool,
	}
}

func (ub *UserBalance) CheckExists(userID int64) (bool, error) {
	var exists bool
	if err := ub.dbpool.QueryRow(
		context.Background(),
		"SELECT EXISTS(SELECT 1 FROM user_balance WHERE user_id = $1)",
		userID).Scan(&exists); err != nil {
		return false, fmt.Errorf("user_balance deposit_get: %w", err)
	}
	return exists, nil
}

func (ub *UserBalance) DepositInsert(deposit models.Deposit) error {
	/*
		метод вызывается только после проверки на существование записи
		с данным user_id, поэтому никаких ошибок быть не должно
	*/
	if _, err := ub.dbpool.Exec(
		context.Background(),
		"INSERT INTO user_balance (user_id, balance) VALUES ($1, $2)",
		deposit.UserID, deposit.Money); err != nil {
		return fmt.Errorf("user_balance deposit_insert: %w", err)
	}
	return nil
}

func (ub *UserBalance) DepositUpdate(deposit models.Deposit) error {
	/*
		метод вызывается только после проверки на существование записи
		с данным user_id, поэтому никаких ошибок быть не должно
	*/
	if _, err := ub.dbpool.Exec(
		context.Background(),
		"UPDATE user_balance SET balance = balance + $2 WHERE user_id = $1",
		deposit.UserID, deposit.Money); err != nil {
		return fmt.Errorf("user_balance deposit_update: %w", err)
	}
	return nil
}

func (ub *UserBalance) Withdraw(withdrawal models.Withdrawal) error {
	/*
		на балансе висит ограничение CHECK (balance::numeric >= 0)
		поэтому в случае чего вернётся ошибка
	*/
	if _, err := ub.dbpool.Exec(
		context.Background(),
		"UPDATE user_balance SET balance = balance - $2 WHERE user_id = $1",
		withdrawal.UserID, withdrawal.Money); err != nil {
		return fmt.Errorf("user_balance withdraw: %w", err)
	}
	return nil
}

func (ub *UserBalance) Transact(transaction models.Transaction) error {
	/*
		на балансе висит ограничение CHECK (balance::numeric >= 0)
		поэтому в случае чего вернётся ошибка
	*/
	if _, err := ub.dbpool.Exec(
		context.Background(),
		`
		UPDATE user_balance SET balance = CASE
			WHEN user_id = $1 THEN balance - $3
			WHEN user_id = $2 THEN balance + $3
		END
		WHERE user_id IN ($1, $2)`,
		transaction.FromID, transaction.ToID, transaction.Money); err != nil {
		return fmt.Errorf("user_balance transact: %w", err)
	}
	return nil
}

func (ub *UserBalance) Balance(balanceRequest models.BalanceRequest) (models.BalanceResponse, error) {
	/*
		база всегда возвращает значение из поля типа money в формате $1,234,567.89
		поэтому нужно немного скорректировать результирующую строку:

		удалить все запятые и знак доллара
	*/
	var balanceResponse models.BalanceResponse
	if err := ub.dbpool.QueryRow(
		context.Background(),
		`
		SELECT
			user_id,
			REGEXP_REPLACE(balance::text, '[$,]', '', 'g')
		FROM
			user_balance
		WHERE
			user_id = $1`,
		balanceRequest.UserID).Scan(&balanceResponse.UserID, &balanceResponse.Money); err != nil {
		return balanceResponse, fmt.Errorf("user_balance balance: %w", err)
	}
	return balanceResponse, nil
}
