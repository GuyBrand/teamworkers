package crud

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"net/http"
	"time"
)

//Not Really a Service Discovery in case you wondered...

var services = map[string]string{
	"product":   ":9000",
	"basket":    ":9001",
	"inventory": ":9002",
}

func fakeDiscovery(svc string) (string, bool) {
	if port, ok := services[svc]; !ok {
		return "", false
	} else {
		return "http://localhost" + port, true
	}
}

func GetEntityFromService(svc, id string, entity interface{}) error {
	if f := GetLocalGetter(svc); f != nil {
		fmt.Printf("getting %s/%s using in memory !\n", svc, id)
		tempEnt := f(id)
		return copier.Copy(entity, tempEnt)
	} else if url, ok := fakeDiscovery(svc); !ok {
		return fmt.Errorf("service %s not found", svc)
	} else {
		var cl = &http.Client{Timeout: 5 * time.Second}
		url += "/" + svc + "/" + id
		if r, err := cl.Get(url); err != nil {
			return fmt.Errorf("cant connect %s, err:%s", url, err.Error())
		} else if r.StatusCode != http.StatusOK {
			return fmt.Errorf("error getting entity %s/%s status %d", svc, id, r.StatusCode)
		} else {
			defer r.Body.Close()
			return json.NewDecoder(r.Body).Decode(entity)
		}
	}
}
