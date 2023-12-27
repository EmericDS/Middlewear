package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleweare/internal/controllers/songs"
	"middleweare/internal/helpers"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Route("/songs", func(r chi.Router) {
		r.Get("/", songs.GetSongs)
		r.Post("/", songs.CreateSong)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(songs.Ctx)
			r.Get("/", songs.GetSong)
			r.Put("/", songs.UpdateSong)
			r.Delete("/", songs.DeleteSong)
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
	defer helpers.CloseDB(db)

	schemes := []string{
		`CREATE TABLE IF NOT EXISTS songs (
			id UUID PRIMARY KEY NOT NULL,
			title VARCHAR(255) NOT NULL,
			artist VARCHAR(255) NOT NULL,
			album VARCHAR(255) NOT NULL,
			release_year INT NOT NULL
		);`,
	}

	for _, scheme := range schemes {
		_, err := db.Exec(scheme)
		if err != nil {
			logrus.Fatalf("Could not generate table! Error was : %s", err.Error())
		}
	}
}