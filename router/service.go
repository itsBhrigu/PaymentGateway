package router

import (
	"context"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"math/rand"
	"paymentGateway/api/bank"
	"paymentGateway/api/pay"
	"paymentGateway/api/router"
	"paymentGateway/db/dao"
	"paymentGateway/router/rules"
)

type Service struct {
	router.UnimplementedRouterServer
	dao     dao.RoutingRuleDao
	bankSvc bank.BankClient
}

func NewService(dao dao.RoutingRuleDao, svc bank.BankClient) *Service {
	return &Service{dao: dao, bankSvc: svc}
}

func (s *Service) Route(ctx context.Context, req *router.RouteRequest) (*router.RouteResponse, error) {
	routingRule, err := s.dao.GetRoutingRule(req.ClientCode)
	if err != nil {
		return &router.RouteResponse{Status: &status.Status{Code: int32(codes.NotFound)}}, nil
	}

	bankCode := getDestinationBank(routingRule, req.Paymode.ModeType)
	_, _ = s.bankSvc.ProcessPayment(ctx, &bank.ProcessPaymentRequest{Payment: &pay.Bank_Transaction{Id: "RandomID", Money: req.Money}, BankCode: bankCode})

	return &router.RouteResponse{Status: &status.Status{Code: int32(codes.OK)}}, nil
}

func (s *Service) SetRoutingRule(ctx context.Context, req *router.SetRoutingRuleRequest) (*router.SetRoutingRuleResponse, error) {
	_ = s.dao.SetRoutingRule(req.Client, req.GetRule())
	return &router.SetRoutingRuleResponse{Status: &status.Status{Code: int32(codes.OK)}}, nil
}

func getDestinationBank(routingRule rules.RoutingRule, paymode pay.PayModes) bank.BankCode {
	var bankCode bank.BankCode

	switch routingRule.(type) {
	case router.TrafficBased:
		randNum := 100 * rand.Float32()
		bankCode = getBankFromTrafficSplit(randNum, serialiseBankSplit(routingRule.(router.TrafficBased).Split))
	case router.PayModeBased:
		bankCode = bank.BankCode(bank.BankCode_value[routingRule.(router.PayModeBased).Split[paymode.String()]])
	}

	return bankCode
}

func getBankFromTrafficSplit(num float32, split map[bank.BankCode]float32) bank.BankCode {
	for bank, percentage := range split {
		num = num - percentage
		if num <= 0 {
			return bank
		}
	}
	return bank.BankCode_BankCode_UNDEFINED
}

func serialiseBankSplit(split map[string]float32) map[bank.BankCode]float32 {
	m := make(map[bank.BankCode]float32)
	for key, val := range split {
		newKey := bank.BankCode_value[key]
		m[bank.BankCode(newKey)] = val
	}
	return m
}
