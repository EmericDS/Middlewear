package users

import (
	"middleware/example/internal/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// DeleteUser
// @Tags         users
// @Summary      Delete an existing user.
// @Description  Delete an existing user.
// @Param        id            path      string  true  "User UUID formatted ID"
// @Success      204            "User deleted successfully"
// @Failure      404            "User not found"
// @Failure      500            "Something went wrong"
// @Router       /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, _ := uuid.FromString(chi.URLParam(r, "id"))

	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Errorf("error while opening database: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer helpers.CloseDB(db)

	result, err := db.Exec("DELETE FROM users WHERE id=?", userID)
	if err != nil {
		logrus.Errorf("error deleting user: %s", err.Error())
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
