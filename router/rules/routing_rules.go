package rules

import (
	"paymentGateway/api/router"
)

type RoutingRule interface {
}

var _ RoutingRule = &router.TrafficBased{}
var _ RoutingRule = &router.PayModeBased{}