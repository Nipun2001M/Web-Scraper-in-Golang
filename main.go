package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	
	godotenv.Load(".env")
	portString:=os.Getenv("PORT")
	if portString==""{
		log.Fatal("PORT Not Found in the Environment variables")
	}
	router :=chi.NewRouter()
	srv:=&http.Server{
		Handler: router,
		Addr: ":"+portString,

	}
	log.Printf("Server Starting on PORT  : %v",portString)
	err:=srv.ListenAndServe()
	if err!=nil{
		log.Fatal("Error : Listning To Server")
	}

}