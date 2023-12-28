package ratings

import (
    "github.com/gofrs/uuid"
    "github.com/sirupsen/logrus"
    "middleware/example/internal/helpers"
    "net/http"
)

// DeleteRating
// @Tags         ratings
// @Summary      Delete an existing rating.
// @Description  Delete an existing rating.
// @Param        id            path      string  true  "Rating UUID formatted ID"
// @Success      204            "Rating deleted successfully"
// @Failure      404            "Rating not found"
// @Failure      500            "Something went wrong"
// @Router       /ratings/{id} [delete]
func DeleteRating(w http.ResponseWriter, r *http.Request) {
    ratingID, _ := uuid.FromString(chi.URLParam(r, "id"))

    db, err := helpers.OpenDB()
    if err != nil {
        logrus.Errorf("error while opening database: %s", err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    defer helpers.CloseDB(db)

    result, err := db.Exec("DELETE FROM ratings WHERE id=?", ratingID)
    if err != nil {
        logrus.Errorf("error deleting rating: %s", err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
