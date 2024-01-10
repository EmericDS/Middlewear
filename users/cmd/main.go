package main

import (
	"middleware/example/internal/controllers/collections" // Assure-toi que le chemin d'importation est correct
	"middleware/example/internal/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) { // Utilise "users" à la place de "collections"
		r.Get("/", collections.GetUsers) // Utilise GetUsers à la place de GetCollections
		r.Route("/{id}", func(r chi.Router) {
			r.Use(collections.Ctx)
			r.Get("/", collections.GetUser) // Utilise GetUser à la place de GetCollection
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:7854")
	logrus.Fatalln(http.ListenAndServe(":7854", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
