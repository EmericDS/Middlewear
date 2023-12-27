package collections

import (
	"encoding/json"
	"middleware/example/internal/models"
	"net/http"

	"github.com/sirupsen/logrus"
)

// CreateUser
// @Tags         users
// @Summary      Create a new user.
// @Description  Create a new user.
// @Param        user          body      models.User  true  "User object"
// @Success      201            "User created successfully"
// @Failure      400            "Invalid request payload"
// @Failure      500            "Something went wrong"
// @Router       /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		logrus.Errorf("error decoding request payload: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: Implémenter la logique pour créer un nouvel utilisateur dans la base de données

	w.WriteHeader(http.StatusCreated)
}
