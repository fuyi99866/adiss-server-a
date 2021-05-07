package request

import "adiss-server-a/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}