package response

import "adiss-server-a/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
