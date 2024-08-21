package Transaction

import "github.com/mohammadrahimi/PayMent_Service_Go/src/Domain/PayMent/Errors"

type Transaction struct {
	TransactionCode int32
	TansactionState string
}

func New(transactionCode int32, tansactionState string) (Transaction, error) {
	if transactionCode <= 0 && tansactionState == "" {
		return Transaction{}, Errors.ErrorTransaction
	}

	return Transaction{
		TransactionCode: transactionCode,
		TansactionState: tansactionState,
	}, nil

}

func (m *Transaction) GetransactionCode() int32 {
	return m.TransactionCode
}
func (m *Transaction) GetTansactionState() string {
	return m.TansactionState
}