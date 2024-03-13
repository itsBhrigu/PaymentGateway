package startup

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"paymentGateway/api/bank"
	"paymentGateway/api/pg"
	"paymentGateway/api/router"
)

const ListenerPort = ":8080"

var RouterClient router.RouterClient
var BankClient bank.BankClient
var PGClient pg.PaymentGatewayClient

func CreateConnection() *grpc.ClientConn {
	conn, err := grpc.Dial(ListenerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("client connection failure: %v", err)
	}
	return conn
}

func CreateClients(conn *grpc.ClientConn) {
	PGClient = pg.NewPaymentGatewayClient(conn)
	RouterClient = router.NewRouterClient(conn)
	BankClient = bank.NewBankClient(conn)
}
