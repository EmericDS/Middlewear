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

