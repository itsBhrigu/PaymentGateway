package main

import (
	"fmt"
	"paymentGateway/api/client"
	"paymentGateway/api/pg"
	startup "paymentGateway/main"
)

func registerClient()  {
	res, _ := startup.PGClient.AddClient(ctx, &pg.AddClientRequest{Client: &client.Client{ClientCode: client.PopularClientCodes_AMAZON.String()}})
	fmt.Println(res)
}