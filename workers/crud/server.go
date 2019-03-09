package crud

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/twinj/uuid"
	"net/http"
)

type handler func(http.ResponseWriter, *http.Request)

func InitRouter(entity string, c, r, u, d handler, block bool, gtF GetterFunc) error {

	if addr, ok := services[entity]; !ok {
		err := fmt.Errorf("service %s is not registered", entity)
		log(err.Error())
		return err
	} else {

		Getters[entity] = gtF

		router := mux.NewRouter()
		entity = "/" + entity
		withId := entity + "/{id}"
		router.Handle(entity, logMW(http.HandlerFunc(r))).Methods("GET")
		router.Handle(withId, logMW(http.HandlerFunc(r))).Methods("GET")

		router.Handle(entity, logMW(http.HandlerFunc(c))).Methods("POST")
		router.Handle(entity, logMW(http.HandlerFunc(u))).Methods("PUT")
		router.Handle(withId, logMW(http.HandlerFunc(d))).Methods("DELETE")

		router.Handle("/local", logMW(http.HandlerFunc(setLocalMode))).Methods("POST")

		listen := func() error {
			log(entity + " listening on " + addr)
			if err := http.ListenAndServe(addr, router); err != nil {
				log(fmt.Sprintf("error: %s trying to serve %s on %s", err.Error(), entity, addr))
				return err
			}
			return nil
		}
		if block {
			return listen()
		} else {
			go listen()
		}

		return nil
	}
}

func logMW(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Host, r.Method, " : ", r.URL)
		h.ServeHTTP(w, r)
	})
}

func GetReqId(r *http.Request) string {
	p := mux.Vars(r)
	return p["id"]
}

func GetUUID() string {
	return uuid.NewV4().String()

}
