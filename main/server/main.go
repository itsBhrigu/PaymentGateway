package main

import (
	"log"
	"net"
	startup "paymentGateway/main"

	pg2 "paymentGateway/api/pg"
	bank2 "paymentGateway/bank"

	"google.golang.org/grpc"
	"paymentGateway/api/bank"
	"paymentGateway/api/router"
	"paymentGateway/db/dao"
	"paymentGateway/pg"
	router2 "paymentGateway/router"
)

var server *grpc.Server

var pgSvc *pg.Service
var routerSvc *router2.Service
var bankSvc *bank2.Service

func main() {
	listener := startListening()

	conn := startup.CreateConnection()
	defer conn.Close()

	server = grpc.NewServer()

	startup.CreateClients(conn)
	createServices()

	registerServices()

	_ = server.Serve(listener)
}

func startListening() net.Listener {
	listener, err := net.Listen("tcp", startup.ListenerPort)
	if err != nil {
		log.Fatalln("failed to listen %v", err)
	}
	return listener
}


func createServices() {
	pgSvc = pg.NewService(dao.NewInMemoryClientDB(), dao.NewInMemoryPayModeDB(), startup.RouterClient)
	routerSvc = router2.NewService(dao.NewInMemoryRoutingDB(), startup.BankClient)
	bankSvc = bank2.NewService()
}

func registerServices() {
	pg2.RegisterPaymentGatewayServer(server, pgSvc)
	router.RegisterRouterServer(server, routerSvc)
	bank.RegisterBankServer(server, bankSvc)
}
