package main

import (
	"net/http"
	"database/sql"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"// I used sqlite in my local to test. It is much more easy to build environment.

	handlers "./handlers"
)

func main() {
	router := mux.NewRouter()
	var err error
	db, err := initialDB("user", "password", "retail_ai_sample_db")
	
	handlers.InitializeRecipeHandlers(router, db)
	http.ListenAndServe(":8080", router)
}

func initialDB(userName, password, dbName string) (*sql.DB, error){
	var err error
	db, err := sql.Open("") // the database configuration
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}