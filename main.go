package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"webscraper/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct{
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT Not Found in the Environment variables")
	}
	dbURL:=os.Getenv("DB_URL")
	 if dbURL==""{
		log.Fatal("DB_URL Not Found in the Environment")
	 }
	conn,err:=sql.Open("postgres",dbURL)
	if err !=nil{
		log.Fatal("Cant Connect to database : ",err)
	}

	
	apiConfig:=apiConfig{
		DB:	database.New(conn)

	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/error", handlerError)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server Starting on PORT : %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error: Listening To Server : ", err)
	}
}

// go build -o rsagg.exe; ./rsagg.exe
