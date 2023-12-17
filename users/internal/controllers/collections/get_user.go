package collections

import (
	"database/sql"
	"encoding/json"
	"errors"
	"middleware/example/internal/repositories/collections"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// GetUser
// @Tags         users
// @Summary      Get a user.
// @Description  Get a user.
// @Param        id           	path      string  true  "User UUID formatted ID"
// @Success      200            {object}  models.User
// @Failure      404            "User not found"
// @Failure      500            "Something went wrong"
// @Router       /users/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, _ := ctx.Value("userID").(uuid.UUID)

	user, err := collections.GetUserById(userID) // Assure-toi que tu fais référence au bon package ici
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		logrus.Errorf("error : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(user)
	_, _ = w.Write(body)
	return
}
