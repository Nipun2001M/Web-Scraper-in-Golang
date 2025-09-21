package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	
	godotenv.Load(".env")
	portString:=os.Getenv("PORT")
	if portString==""{
		log.Fatal("PORT Not Found in the Environment variables")
	}


	router :=chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, 
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, 
	}))


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