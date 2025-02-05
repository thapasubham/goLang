package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/thapasubham/go-learn/cmd/service/expense"
	service "github.com/thapasubham/go-learn/cmd/service/user"
	"github.com/thapasubham/go-learn/cmd/utils"
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Adjust for your frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	router.HandleFunc("/", index)

	subRouter := router.PathPrefix("/api/").Subrouter()

	storeHander := service.NewStore(s.db)
	userHander := service.NewHandler(storeHander)

	expenseStore := expense.NewStore(s.db)
	expenseHandler := *expense.NewHandler(expenseStore)

	expenseHandler.RegisterRoutes(subRouter)
	userHander.RegisterRoutes(subRouter)

	fmt.Println("Listening on: ", s.addr)
	handler := c.Handler(router)
	return http.ListenAndServe(s.addr, handler)
}
func index(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w)
	utils.WriteJson(w, 200, map[string]string{"hi": "welcome"})
}
