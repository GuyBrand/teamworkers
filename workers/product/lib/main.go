package product

import (
	"teamworkers/workers/crud"
)

func Start(lastService bool) {
	crud.InitRouter("product", createH, readH, updateH, deleteH, lastService, GetByIdInterface)
}
