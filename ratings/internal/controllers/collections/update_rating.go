package ratings

import (
    "encoding/json"
    "github.com/gofrs/uuid"
    "github.com/sirupsen/logrus"
    "middleware/example/internal/helpers"
    "middleware/example/internal/models"
    "net/http"
)

// UpdateRating
// @Tags         ratings
// @Summary      Update an existing rating.
// @Description  Update an existing rating.
// @Param        id            path      string  true  "Rating UUID formatted ID"
// @Param        rating        body      models.Rating  true  "Rating object"
// @Success      200            "Rating updated successfully"
// @Failure      400            "Invalid request payload"
// @Failure      404            "Rating not found"
// @Failure      500            "Something went wrong"
// @Router       /ratings/{id} [put]
func UpdateRating(w http.ResponseWriter, r *http.Request) {
    ratingID, _ := uuid.FromString(chi.URLParam(r, "id"))
    var updatedRating models.Rating
    err := json.NewDecoder(r.Body).Decode(&updatedRating)
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

    result, err := db.Exec("UPDATE ratings SET user_id=?, song_id=?, score=?, comment=? WHERE id=?",
        updatedRating.UserID, updatedRating.SongID, updatedRating.Score, updatedRating.Comment, ratingID)
    if err != nil {
        logrus.Errorf("error updating rating: %s", err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
}
