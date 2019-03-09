package product

import (
	"encoding/json"
	"net/http"
	"teamworkers/workers/crud"
)

func createH(w http.ResponseWriter, r *http.Request) {
	var p product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	Add(p)
	w.WriteHeader(http.StatusCreated)
}

func readH(w http.ResponseWriter, r *http.Request) {
	if id := crud.GetReqId(r); id == "" {
		json.NewEncoder(w).Encode(products)
	} else if p := GetById(id); p != nil {
		json.NewEncoder(w).Encode(p)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&product{})
	}
}

func updateH(w http.ResponseWriter, r *http.Request) {
	var p product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else if Update(p) {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func deleteH(w http.ResponseWriter, r *http.Request) {
	if id := crud.GetReqId(r); id == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else if Delete(id) {
		w.WriteHeader(http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
