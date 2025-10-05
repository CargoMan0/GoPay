package server

import (
	"context"
	"github.com/google/uuid"
	"github/com/CargoMan0/GoPay/pkg/accountmanager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TODO: Finish

func (s *Server) NewAccount(ctx context.Context, req *accountmanager.NewAccountRequest) (*accountmanager.NewAccountResponse, error) {

}

func (s *Server) GetAccount(ctx context.Context, req *accountmanager.GetAccountRequest) (*accountmanager.GetAccountResponse, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	account, err := s.accountService.GetAccount(ctx, id)
	if err != nil {
		return nil, handleError(err)
	}

	_ = account
}

func (s *Server) ChangePassword() {
	// s.accountService.
}
