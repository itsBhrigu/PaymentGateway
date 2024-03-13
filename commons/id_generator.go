package commons

import (
	uuid2 "github.com/google/uuid"
	"paymentGateway/api/commons"
)

func GenerateID(category commons.IdCategory) string {
	uuid := uuid2.New()
	id := category.String() + "_" + uuid.String()[:4]
	return id
}
