package response

import "adiss-server-a/model"

type SysAuthorityResponse struct {
	Authority model.SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      model.SysAuthority `json:"authority"`
	OldAuthorityId string             `json:"oldAuthorityId"`
}
