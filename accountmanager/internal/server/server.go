package server

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github/com/CargoMan0/GoPay/accountmanager/internal/config"
	"github/com/CargoMan0/GoPay/accountmanager/internal/entity"
	"github/com/CargoMan0/GoPay/pkg/accountmanager"
	"google.golang.org/grpc"
	"net"
)

type AccountService interface {
	NewAccount(ctx context.Context, data *entity.NewAccountData) (*entity.NewAccountResult, error)
	GetAccount(ctx context.Context, id uuid.UUID) (*entity.Account, error)
}

type Server struct {
	accountmanager.UnimplementedAccountManagerServer
	grpcServer     *grpc.Server
	accountService AccountService
	port           int
}

func NewGRPC(cfg config.GRPCServer, accountService AccountService) *Server {
	grpcSrv := grpc.NewServer()

	server := &Server{
		grpcServer:     grpcSrv,
		port:           cfg.Port,
		accountService: accountService,
	}
	accountmanager.RegisterAccountManagerServer(grpcSrv, server)

	return server
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	err = s.grpcServer.Serve(listener)
	if err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}
