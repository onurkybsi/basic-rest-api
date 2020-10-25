package models

// TransactionResponse struct
type TransactionResponse struct {
	IsSuccess bool   `json:"isSuccess"`
	Message   string `json:"message"`
}
