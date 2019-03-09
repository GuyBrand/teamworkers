package basket

import (
	"encoding/json"
	"net/http"
	"teamworkers/workers/basket/lib/proto"
	"teamworkers/workers/crud"
)

func createH(w http.ResponseWriter, r *http.Request) {
	var b basketitem.AddBasketItemRequest
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else if !Add(b) {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func readH(w http.ResponseWriter, r *http.Request) {
	if id := crud.GetReqId(r); id == "" {
		json.NewEncoder(w).Encode(basket)
	} else if p := GetById(id); p != nil {
		json.NewEncoder(w).Encode(p)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&item{})
	}
}

func updateH(w http.ResponseWriter, r *http.Request) {
	var b item
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else if Update(b) {
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
