package ratings

import (
    "encoding/json"
    "github.com/gofrs/uuid"
    "github.com/sirupsen/logrus"
    "middleware/example/internal/models"
    "middleware/example/internal/services/ratings"
    "net/http"
)

// GetRating
// @Tags         ratings
// @Summary      Get a specific rating.
// @Description  Get a specific rating.
// @Param        id            path      string  true  "Rating UUID formatted ID"
// @Success      200            {object}  models.Rating
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /ratings/{id} [get]
func GetRating(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    ratingID, _ := ctx.Value("ratingID").(uuid.UUID)

    rating, err := ratings.GetRatingByID(ratingID)
    if err != nil {
        logrus.Errorf("error: %s", err.Error())
        customError, isCustom := err.(*models.CustomError)
        if isCustom {
            w.WriteHeader(customError.Code)
            body, _ := json.Marshal(customError)
            _, _ = w.Write(body)
        } else {
            w.WriteHeader(http.StatusInternalServerError)
        }
        return
    }

    w.WriteHeader(http.StatusOK)
    body, _ := json.Marshal(rating)
    _, _ = w.Write(body)
}
