package request

import "adiss-server-a/model"

type SysDictionarySearch struct{
    model.SysDictionary
    PageInfo
}