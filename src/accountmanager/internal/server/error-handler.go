package server

import (
	"errors"
	"github/com/CargoMan0/GoPay/src/accountmanager/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func handleError(err error) error {
	switch {
	case errors.Is(err, service.ErrWrongPassword):
		return status.Errorf(codes.FailedPrecondition, "Wrong password")
	case errors.Is(err, service.ErrAccountNotFound):
		return status.Errorf(codes.NotFound, "Account not found")
	case errors.Is(err, service.ErrAccountAlreadyExists):
		return status.Errorf(codes.AlreadyExists, "Account already exists")
	default:
		return status.Errorf(codes.Internal, "Internal server error: %v", err)
	}
}
