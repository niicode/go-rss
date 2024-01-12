package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/niicode/go-rss/internal/db"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *db.Queries
}

func main() {

	godotenv.Load(".env")
	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		log.Fatal("$DB_URL is not found in the evnironment variable")
	}

	//connect to db
	conn, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal(err)
	}

	qry := db.New(conn)

	apiCfg := &apiConfig{
		DB: qry,
	}


	portString := os.Getenv("PORT")
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1 := chi.NewRouter()

	v1.Get("/healthz", handlerReadiness)
	v1.Get("/error", handlerError)
	v1.Post("/users", apiCfg.handlerCreateUser)
	v1.Get("/user", apiCfg.handlerGetUsers)

	router.Mount("/v1", v1)

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Server is running on port %s", portString)

	erro := srv.ListenAndServe()

	if erro != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

