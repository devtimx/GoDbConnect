package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()
	/* add routers for endpoints
	example
	router.HandleFunc("/signup", middleware.CheckDB(routers.SignUp)).Methods("POST")
	*/
	PORT := os.Getenv("PORT")
	/* set port for server */
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
