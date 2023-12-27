package collections

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
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

	// TODO: Implémenter la logique pour supprimer l'utilisateur dans la base de données avec l'ID userID

	w.WriteHeader(http.StatusNoContent)
}
