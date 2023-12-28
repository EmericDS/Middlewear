package main

import (
    "github.com/go-chi/chi/v5"
    "github.com/sirupsen/logrus"
    "middleware/example/internal/controllers/ratings"
    "middleware/example/internal/helpers"
    "net/http"
)

func main() {
    r := chi.NewRouter()

    r.Route("/ratings", func(r chi.Router) {
        r.Get("/", ratings.GetRatings)
        r.Post("/", ratings.CreateRating)
        r.Route("/{id}", func(r chi.Router) {
            r.Use(ratings.Ctx)
            r.Get("/", ratings.GetRating)
            r.Put("/", ratings.UpdateRating)
            r.Delete("/", ratings.DeleteRating)
        })
    })

    logrus.Info("[INFO] Web server started. Now listening on *:7855")
    logrus.Fatalln(http.ListenAndServe(":7855", r))
}

func init() {
    db, err := helpers.OpenDB()
    if err != nil {
        logrus.Fatalf("error while opening database : %s", err.Error())
    }
    schemes := []string{
        `CREATE TABLE IF NOT EXISTS ratings (
            id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
            user_id VARCHAR(255) NOT NULL,
            song_id VARCHAR(255) NOT NULL,
            rating INT NOT NULL,
            comments VARCHAR(255)
        );`,
    }
    for _, scheme := range schemes {
        if _, err := db.Exec(scheme); err != nil {
            logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
        }
    }
    helpers.CloseDB(db)
}
