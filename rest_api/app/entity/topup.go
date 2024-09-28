package entity

import "time"

type TopUp struct {
	TopUpID       string    `json:"top_up_id"`
	AmountTopUp   int64     `json:"amount_top_up"`
	BalanceBefore int64     `json:"balance_before"`
	BalanceAfter  int64     `json:"balance_after"`
	CreatedDate   time.Time `json:"created_date"`
}
