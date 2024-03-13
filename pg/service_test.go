package pg

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/currency"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/genproto/googleapis/type/money"
	"google.golang.org/grpc/codes"
	"paymentGateway/api/client"
	"paymentGateway/api/pay"
	"paymentGateway/api/pg"
	"paymentGateway/api/router"
	mocks2 "paymentGateway/mocks/api/router"
	mocks "paymentGateway/mocks/db/dao"
	"testing"
)

func TestPgService_AddClient(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()

	mockClientDao := mocks.NewMockClientDao(ctr)
	mockPayDao := mocks.NewMockPayModeDao(ctr)
	mockRouterSvc := mocks2.NewMockRouterClient(ctr)

	svc := NewService(mockClientDao, mockPayDao, mockRouterSvc)
	ctx := context.Background()

	tests := []struct {
		name      string
		req       *pg.AddClientRequest
		mockCalls func()
		res       *pg.AddClientResponse
		err       error
	}{
		{
			name: "successful add client",
			req: &pg.AddClientRequest{Client: &client.Client{
				Name: "AMAZON",
			}},
			mockCalls: func() {
				mockClientDao.EXPECT().AddClient(&client.Client{Name: "AMAZON"}).Return(client.PopularClientCodes_AMAZON.String(), nil)
			},
			res: &pg.AddClientResponse{
				Id:     client.PopularClientCodes_AMAZON.String(),
				Status: &status.Status{Code: int32(codes.OK)},
			},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockCalls()

			res, err := svc.AddClient(ctx, test.req)

			assert.Equal(t, test.res, res)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestService_MakePayment(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()

	mockClientDao := mocks.NewMockClientDao(ctr)
	mockPayDao := mocks.NewMockPayModeDao(ctr)
	mockRouterSvc := mocks2.NewMockRouterClient(ctr)

	svc := NewService(mockClientDao, mockPayDao, mockRouterSvc)
	ctx := context.Background()

	money := &money.Money{
		CurrencyCode: currency.INR.String(),
		Units:        5000,
	}
	paymode := &pay.PayMode{
		ModeType:    pay.PayModes_UPI_VPA,
		PaymentInfo: &pay.PayMode_Upi{Upi: &pay.UPI{Vpa: "bhrigu@hdfcbank"}},
	}

	tests := []struct {
		name      string
		req       *pg.MakePaymentRequest
		mockCalls func()
		res       *pg.MakePaymentResponse
		err       error
	}{
		{
			name: "successful payment",
			req: &pg.MakePaymentRequest{PaymentDetails: &pay.PG_Transaction{
				Money:   money,
				Paymode: paymode,
			}, ClientCode: client.PopularClientCodes_AMAZON.String()},
			mockCalls: func() {
				mockRouterSvc.EXPECT().Route(ctx, &router.RouteRequest{
					Money:      money,
					Paymode:    paymode,
					ClientCode: client.PopularClientCodes_AMAZON.String(),
				}).Return(&router.RouteResponse{Status: &status.Status{Code: int32(codes.OK)}}, nil).Times(1)
			},
			res: &pg.MakePaymentResponse{Status: &status.Status{Code: int32(codes.OK)}},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockCalls()

			res, err := svc.MakePayment(ctx, test.req)

			assert.Equal(t, test.res, res)
			assert.Equal(t, test.err, err)
		})
	}
}
