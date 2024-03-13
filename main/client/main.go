package main

import (
	"context"
	startup "paymentGateway/main"
)

var ctx context.Context

func main() {
	conn := startup.CreateConnection()
	defer conn.Close()

	startup.CreateClients(conn)

	ctx = context.Background()

	registerClient()
}
