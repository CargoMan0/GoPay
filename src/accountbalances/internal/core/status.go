package core

type Status uint8

const (
	StatusPending = iota + 1
	StatusFailed
	StatusSuccess
)
