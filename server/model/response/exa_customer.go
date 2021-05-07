package response

import "adiss-server-a/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
