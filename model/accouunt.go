package model

import "time"

type Account struct {
	AccountID  string                 `json:"account_id"`
	Name       string                 `json:"name" form:"name"`
	Number     string                 `json:"number" form:"number"`
	BankName   string                 `json:"bank_name"`
	Currency   string                 `json:"currency"`
	CreatedAt  time.Time              `json:"created_at"`
	BalanceID  string                 `json:"balance_id" `
	IdentityID string                 `json:"identity_id" form:"identity_id"`
	LedgerID   string                 `json:"ledger_id"`
	MetaData   map[string]interface{} `json:"meta_data"`
	Ledger     *Ledger                `json:"ledger"`
	Balance    *Balance               `json:"balance"`
	Identity   *Identity              `json:"identity"`
}
