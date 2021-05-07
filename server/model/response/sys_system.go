package response

import "adiss-server-a/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
