package server

import (
	"context"
	"github.com/google/uuid"
	"github/com/CargoMan0/GoPay/accountmanager/internal/entity"
	"github/com/CargoMan0/GoPay/pkg/accountmanager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) NewAccount(ctx context.Context, req *accountmanager.NewAccountRequest) (*accountmanager.NewAccountResponse, error) {
	data := &entity.NewAccountData{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	res, err := s.accountService.NewAccount(ctx, data)
	if err != nil {
		return nil, handleError(err)
	}

	resp := &accountmanager.NewAccountResponse{
		AccountAddress:   res.AccessToken,
		RefreshToken:     res.RefreshToken,
		RegistrationDate: timestamppb.New(res.RegistrationDate),
	}

	return resp, nil
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

	resp := &accountmanager.GetAccountResponse{
		Username: account.Username,
		Email:    account.Email,
	}

	return resp, nil
}

func (s *Server) ChangePassword(ctx context.Context, req *accountmanager.ChangePasswordRequest) (*emptypb.Empty, error) {
	var data = &entity.ChangePasswordData{
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}

	err := s.accountService.ChangePassword(ctx, data)
	if err != nil {
		return nil, handleError(err)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) LoginAccount(ctx context.Context, req *accountmanager.LoginAccountRequest) (*accountmanager.LoginAccountResponse, error) {
	// TODO: validate
	email := req.Email
	password := req.Password

	res, err := s.accountService.LoginAccount(ctx, email, password)
	if err != nil {
		return nil, handleError(err)
	}

	resp := &accountmanager.LoginAccountResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}

	return resp, nil
}
