package ratings

import (
    "encoding/json"
    "github.com/sirupsen/logrus"
    "middleware/example/internal/models"
    "middleware/example/internal/services/ratings"
    "net/http"
)

// GetRatings
// @Tags         ratings
// @Summary      Get all ratings.
// @Description  Get all ratings.
// @Success      200            {array}  models.Rating
// @Failure      500            "Something went wrong"
// @Router       /ratings [get]
func GetRatings(w http.ResponseWriter, _ *http.Request) {
    // Appeler le service pour récupérer toutes les notes
    ratingsList, err := ratings.GetAllRatings()
    if err != nil {
        logrus.Errorf("error: %s", err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    body, _ := json.Marshal(ratingsList)
    _, _ = w.Write(body)
}
