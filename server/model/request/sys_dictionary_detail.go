package request

import "adiss-server-a/model"

type SysDictionaryDetailSearch struct{
    model.SysDictionaryDetail
    PageInfo
}