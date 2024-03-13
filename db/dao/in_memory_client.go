package dao

import (
	"fmt"
	"paymentGateway/api/client"
	commons2 "paymentGateway/api/commons"
	"paymentGateway/commons"
)

type InMemoryClientDB struct {
	db commons.Set
}

func NewInMemoryClientDB() *InMemoryClientDB {
	return &InMemoryClientDB{db: make(commons.Set)}
}

var _ ClientDao = &InMemoryClientDB{}

func (i *InMemoryClientDB) AddClient(client *client.Client) (string, error) {
	clientCode := getClientCode(client.Name)
	if _, ok := i.db[clientCode]; !ok {
		i.db[clientCode] = commons.EmptyValue
	}
	return clientCode, nil
}

func (i *InMemoryClientDB) DeleteClientByCode(clientCode string) error {
	if _, ok := i.db[clientCode]; !ok {
		return fmt.Errorf("no such client in the database")
	}
	delete(i.db, clientCode)
	return nil
}

func (i *InMemoryClientDB) HasClient(client *client.Client) (bool, error) {
	clientCode := getClientCode(client.Name)
	_, ok := i.db[clientCode]
	return ok, nil
}

func getClientCode(name string) string {
	clientCode, ok := ClientCodes[name]
	if !ok {
		clientCode = generateClientCode(name)
	}
	return clientCode
}

func generateClientCode(name string) string {
	newClientCode := commons.GenerateID(commons2.IdCategory_CLIENT)
	ClientCodes[name] = newClientCode
	return newClientCode
}
