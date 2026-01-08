package core

type (
	OperationType   uint8
	OperationStatus uint8
)

const (
	OperationTypePayment OperationType = iota + 1
	OperationTypeWithdraw
)

const (
	OperationStatusPending OperationStatus = iota + 1
	OperationStatusFailed
	OperationStatusSuccess
)
