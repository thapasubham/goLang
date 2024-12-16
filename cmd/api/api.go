package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thapasubham/go-learn/cmd/service"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {

	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/", index)
	subRouter := router.PathPrefix("/api/").Subrouter()

	storeHander := service.NewStore(s.db)
	userHander := service.NewHandler(storeHander)
	userHander.RegisterRoutes(subRouter)

	fmt.Println("Listening on: ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello What is up fella")
}
