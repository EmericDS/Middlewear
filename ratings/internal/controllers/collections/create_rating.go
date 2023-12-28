package ratings

import (
    "encoding/json"
    "github.com/sirupsen/logrus"
    "middleware/example/internal/helpers"
    "middleware/example/internal/models"
    "net/http"
)

// CreateRating
// @Tags         ratings
// @Summary      Create a new rating.
// @Description  Create a new rating.
// @Param        rating        body      models.Rating  true  "Rating object"
// @Success      201            "Rating created successfully"
// @Failure      400            "Invalid request payload"
// @Failure      500            "Something went wrong"
// @Router       /ratings [post]
func CreateRating(w http.ResponseWriter, r *http.Request) {
    var newRating models.Rating
    err := json.NewDecoder(r.Body).Decode(&newRating)
    if err != nil {
        logrus.Errorf("error decoding request payload: %s", err.Error())
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    db, err := helpers.OpenDB()
    if err != nil {
        logrus.Errorf("error while opening database: %s", err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    defer helpers.CloseDB(db)

    _, err = db.Exec("INSERT INTO ratings (id, user_id, song_id, score, comment) VALUES (?, ?, ?, ?, ?)",
        newRating.ID, newRating.UserID, newRating.SongID, newRating.Score, newRating.Comment)
    if err != nil {
        logrus.Errorf("error creating rating: %s", err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}
