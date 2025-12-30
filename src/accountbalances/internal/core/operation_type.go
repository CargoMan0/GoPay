package core

type OperationType uint8

const (
	OperationTypePayment OperationType = iota + 1
	OperationTypeWithdraw
)
