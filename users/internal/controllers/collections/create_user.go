package users

import (
	"encoding/json"
	"middleware/example/internal/helpers"
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

	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("error while opening database: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("INSERT INTO users (id, username, email) VALUES (?, ?, ?)",
		newUser.ID, newUser.Username, newUser.Email)
	if err != nil {
		logrus.Errorf("error creating user: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
