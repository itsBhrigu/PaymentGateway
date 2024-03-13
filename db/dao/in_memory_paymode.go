package dao

import (
	"fmt"
	"paymentGateway/api/client"
	"paymentGateway/api/pay"
	"paymentGateway/commons"
)

type InMemoryPayModeDB struct {
	dbPG     map[*pay.PayMode]commons.EmptyType
	dbClient map[string][]pay.PayModes
}

func NewInMemoryPayModeDB() *InMemoryPayModeDB {
	return &InMemoryPayModeDB{
		dbPG:     make(map[*pay.PayMode]commons.EmptyType),
		dbClient: make(map[string][]pay.PayModes),
	}
}

var _ PayModeDao = &InMemoryPayModeDB{}

func (i *InMemoryPayModeDB) AddPayMode(mode *pay.PayMode) error {
	i.dbPG[mode] = commons.EmptyValue
	return nil
}

func (i InMemoryPayModeDB) DeletePayMode(mode *pay.PayMode) error {
	if _, ok := i.dbPG[mode]; ok {
		delete(i.dbPG, mode)
		return nil
	}
	return fmt.Errorf("mode hasn't been added to Payment Gateway")
}

func (i InMemoryPayModeDB) ListPayModes() ([]pay.PayModes, error) {
	var payModes []pay.PayModes
	for paymode, _ := range i.dbPG {
		payModes = append(payModes, paymode.ModeType)
	}
	return payModes, nil
}

func (i *InMemoryPayModeDB) AddPayModeForClient(mode *pay.PayMode, client *client.Client) error {
	currentPayModes := i.dbClient[client.ClientCode]
	for _, paymode := range currentPayModes {
		if paymode == mode.ModeType {
			return fmt.Errorf("paymode already exists")
		}
	}
	i.dbClient[client.ClientCode] = append(i.dbClient[client.ClientCode], mode.ModeType)
	return nil
}

func (i *InMemoryPayModeDB) DeletePayModeForClient(mode *pay.PayMode, client *client.Client) error {
	return fmt.Errorf("not implemented")
}

func (i *InMemoryPayModeDB) ListPayModesOfClient() ([]pay.PayModes, error) {
	return nil, fmt.Errorf("not implemented")
}
