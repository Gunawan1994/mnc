package entity

import "time"

type Payment struct {
	PaymentID     string    `json:"payment_id"`
	Amount        int64     `json:"amount"`
	Remarks       string    `json:"remarks"`
	BalanceBefore int64     `json:"balance_before"`
	BalanceAfter  int64     `json:"balance_after"`
	CreatedDate   time.Time `json:"created_date"`
}

type Transfer struct {
	TransferID    string    `json:"transfer_id"`
	Amount        int64     `json:"amount"`
	Remarks       string    `json:"remarks"`
	BalanceBefore int64     `json:"balance_before"`
	BalanceAfter  int64     `json:"balance_after"`
	CreatedDate   time.Time `json:"created_date"`
}
