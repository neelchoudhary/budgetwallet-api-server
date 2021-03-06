package models

import (
	"database/sql"

	"github.com/plaid/plaid-go/plaid"
)

// FinancialTransaction ...
type FinancialTransaction struct {
	ID                       int64   `json:"id"`
	UserID                   int64   `json:"user_id"`
	ItemID                   int64   `json:"item_id"`
	AccountID                int64   `json:"account_id"`
	CategoryID               int64   `json:"category_id"`
	DailyAccountSnapshotID   int64   `json:"daily_account_snapshot_id"`
	MonthlyAccountSnapshotID int64   `json:"monthly_account_snapshot_id"`
	PlaidAccountID           string  `json:"plaid_account_id"`
	PlaidCategoryID          string  `json:"plaid_category_id"`
	PlaidTransactionID       string  `json:"plaid_transaction_id"`
	Name                     string  `json:"transaction_name"`
	Amount                   float64 `json:"amount"`
	Date                     string  `json:"date"`
	Pending                  bool    `json:"pending"`
}

// NewFinancialTransactionFromPlaid creates a new financial transaction from a plaid transaction
func NewFinancialTransactionFromPlaid(userID int64, itemID int64, plaidTransaction plaid.Transaction) FinancialTransaction {
	transaction := FinancialTransaction{
		UserID:             userID,
		ItemID:             itemID,
		PlaidAccountID:     plaidTransaction.AccountID,
		PlaidCategoryID:    plaidTransaction.CategoryID,
		PlaidTransactionID: plaidTransaction.ID,
		Name:               plaidTransaction.Name,
		Amount:             plaidTransaction.Amount,
		Date:               plaidTransaction.Date,
		Pending:            plaidTransaction.Pending,
	}
	return transaction
}

// FilterTransactions filter transactions
func FilterTransactions(ls []FinancialTransaction, f func(FinancialTransaction) bool) []FinancialTransaction {
	fls := make([]FinancialTransaction, 0)
	for _, v := range ls {
		if f(v) {
			fls = append(fls, v)
		}
	}
	return fls
}

// FinancialTransactionRepository interface
type FinancialTransactionRepository interface {
	AddTransaction(tx *sql.Tx, transaction *FinancialTransaction) error
	UpdateTransaction(tx *sql.Tx, userID int64, transactionID int64, transaction *FinancialTransaction) error
	DoesTransactionExist(tx *sql.Tx, userID int64, plaidTransactionID string) (bool, error)
	GetTransactionByID(tx *sql.Tx, userID int64, transactionID int64) (*FinancialTransaction, error)
	GetTransactionByPlaidID(tx *sql.Tx, userID int64, plaidTransactionID string) (*FinancialTransaction, error)
	GetAccountTransactions(tx *sql.Tx, userID int64, accountID int64) ([]FinancialTransaction, error)
	GetItemTransactions(tx *sql.Tx, userID int64, itemID int64) ([]FinancialTransaction, error)
	GetUserTransactions(tx *sql.Tx, userID int64) ([]FinancialTransaction, error)
	RemoveItemTransactions(tx *sql.Tx, userID int64, itemID int64) error
	RemoveUserTransactions(tx *sql.Tx, userID int64) error
	RemoveTransactionByID(tx *sql.Tx, userID int64, transactionID int64) error
}
