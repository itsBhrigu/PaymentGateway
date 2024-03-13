package bank

import (
	"context"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"paymentGateway/api/bank"
)

type Service struct {
	bank.UnimplementedBankServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ProcessPayment(ctx context.Context, req *bank.ProcessPaymentRequest) (*bank.ProcessPaymentResponse, error) {
	return &bank.ProcessPaymentResponse{Status: &status.Status{Code: int32(codes.OK)}}, nil
}
