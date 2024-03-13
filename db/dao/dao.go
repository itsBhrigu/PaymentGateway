package dao

import (
	"paymentGateway/api/client"
	"paymentGateway/api/pay"
	"paymentGateway/router/rules"
)

type ClientDao interface {
	AddClient(*client.Client) (string, error)
	DeleteClientByCode(string) error
	HasClient(*client.Client) (bool, error)
}

type PayModeDao interface {
	AddPayMode(*pay.PayMode) error
	AddPayModeForClient(*pay.PayMode, *client.Client) error
	DeletePayMode(*pay.PayMode) error
	DeletePayModeForClient(*pay.PayMode, *client.Client) error
	ListPayModes() ([]pay.PayModes, error)
	ListPayModesOfClient() ([]pay.PayModes, error)
}

type RoutingRuleDao interface {
	SetRoutingRule(string, rules.RoutingRule) error
	GetRoutingRule(string) (rules.RoutingRule, error)
}