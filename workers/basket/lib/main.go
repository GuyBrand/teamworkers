package basket

import "teamworkers/workers/crud"

func Start(lastService bool) {
	crud.InitRouter("basket", createH, readH, updateH, deleteH, lastService, GetByIdInterface) //not really crudable
}
