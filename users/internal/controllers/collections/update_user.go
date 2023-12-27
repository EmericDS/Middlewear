package collections

import (
	"encoding/json"
	"middleware/example/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// UpdateUser
// @Tags         users
// @Summary      Update an existing user.
// @Description  Update an existing user.
// @Param        id            path      string  true  "User UUID formatted ID"
// @Param        user          body      models.User  true  "User object"
// @Success      200            "User updated successfully"
// @Failure      400            "Invalid request payload"
// @Failure      404            "User not found"
// @Failure      500            "Something went wrong"
// @Router       /users/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, _ := uuid.FromString(chi.URLParam(r, "id"))
	var updatedUser models.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		logrus.Errorf("error decoding request payload: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: Implémenter la logique pour mettre à jour l'utilisateur dans la base de données avec l'ID userID

	w.WriteHeader(http.StatusOK)
}
