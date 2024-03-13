package pg

import (
	"context"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"paymentGateway/api/pay"
	"paymentGateway/api/router"

	"paymentGateway/api/pg"
	"paymentGateway/db/dao"
)

type Service struct {
	pg.UnimplementedPaymentGatewayServer
	clientDao    dao.ClientDao
	payModeDao   dao.PayModeDao
	routerClient router.RouterClient
}

func NewService(clientDao dao.ClientDao, payModeDao dao.PayModeDao, client router.RouterClient) *Service {
	return &Service{clientDao: clientDao, payModeDao: payModeDao, routerClient: client}
}

func (s *Service) AddClient(ctx context.Context, req *pg.AddClientRequest) (*pg.AddClientResponse, error) {
	id, _ := s.clientDao.AddClient(req.Client)
	return &pg.AddClientResponse{Id: id, Status: &status.Status{Code: int32(codes.OK)}}, nil
}

func (s *Service) DeleteClient(ctx context.Context, req *pg.DeleteClientRequest) (*pg.DeleteClientResponse, error) {
	_ = s.clientDao.DeleteClientByCode(req.ClientCode)
	return &pg.DeleteClientResponse{Status: &status.Status{Code: int32(codes.OK)}}, nil
}

func (s *Service) HasClient(ctx context.Context, req *pg.HasClientRequest) (*pg.HasClientResponse, error) {
	ok, _ := s.clientDao.HasClient(req.Client)
	return &pg.HasClientResponse{
		HasClient: ok,
		Status:    &status.Status{Code: int32(codes.OK)},
	}, nil
}

func (s *Service) ListPayModes(ctx context.Context, req *pg.ListPayModesRequest) (*pg.ListPayModesResponse, error) {
	var paymodes []pay.PayModes
	if req.ClientCode == "" {
		paymodes, _ = s.payModeDao.ListPayModes()
	} else {
		paymodes, _ = s.payModeDao.ListPayModesOfClient()
	}
	return &pg.ListPayModesResponse{
		Paymodes: paymodes,
		Status:   &status.Status{Code: int32(codes.OK)},
	}, nil
}

func (s *Service) MakePayment(ctx context.Context, req *pg.MakePaymentRequest) (*pg.MakePaymentResponse, error) {
	res, err := s.routerClient.Route(ctx, &router.RouteRequest{
		Money: req.PaymentDetails.Money,
		Paymode: &pay.PayMode{
			ModeType:    pay.PayModes_UPI_VPA,
			PaymentInfo: req.PaymentDetails.Paymode.PaymentInfo,
		},
		ClientCode: req.ClientCode,
	})

	if err != nil {
		return &pg.MakePaymentResponse{Status: &status.Status{Code: int32(codes.Internal)}}, err
	}

	if res.Status.Code != int32(codes.OK) {
		return &pg.MakePaymentResponse{Status: &status.Status{Code: res.Status.Code}}, nil
	}

	return &pg.MakePaymentResponse{Status: &status.Status{Code: int32(codes.OK)}}, nil
}
