package base

// TransactionNoGenerator 生成经济交易流水号。
type TransactionNoGenerator interface {
	NewTransactionNo() (string, error)
}
