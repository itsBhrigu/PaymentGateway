package dao

import (
	"fmt"
	"paymentGateway/router/rules"
)

type InMemoryRoutingDB struct {
	db map[string]rules.RoutingRule
}

func NewInMemoryRoutingDB() *InMemoryRoutingDB {
	return &InMemoryRoutingDB{db: make(map[string]rules.RoutingRule)}
}

var _ RoutingRuleDao = &InMemoryRoutingDB{}

func (i *InMemoryRoutingDB) SetRoutingRule(s string, rule rules.RoutingRule) error {
	i.db[s] = rule
	return nil
}

func (i *InMemoryRoutingDB) GetRoutingRule(s string) (rules.RoutingRule, error) {
	if rule, ok := i.db[s]; ok {
		return rule, nil
	}
	return nil, fmt.Errorf("no rule for the client")
}
