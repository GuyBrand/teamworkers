package crud

import (
	"encoding/json"
	"net/http"
)

// added code - straight forward implementation - must use go-like code, interface with Get method, and UseLocal method !!
type GetterFunc func(id string) interface{}

var Getters map[string]GetterFunc
var localMode bool

func init() {
	Getters = make(map[string]GetterFunc)
}

func GetLocalGetter(svc string) GetterFunc {
	if f, ok := Getters[svc]; ok && localMode {
		return f
	}
	return nil
}

func setLocalMode(w http.ResponseWriter, r *http.Request) {
	type localstat struct {
		Status bool `json:"status"`
	}
	var l localstat
	if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	localMode = l.Status
	w.WriteHeader(http.StatusCreated)
}
