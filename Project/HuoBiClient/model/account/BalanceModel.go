package account

import (

)

type BalanceModel struct {
	Currency string `json:"Currency"`
	Type     string `json:"Type"`
	Amount  float64 `json:"Amount"`
}