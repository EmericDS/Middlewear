package users

import (
	"database/sql"
	"errors"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/collections"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllUsers() ([]models.User, error) {
	var err error
	// calling repository
	users, err := repository.GetAllUsers()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving users: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return users, nil
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	user, err := repository.GetUserById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving users: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return user, err
}

func CreateUser(newUser models.User) error {
    db, err := helpers.OpenDB()
    if err != nil {
        return err
    }
    defer helpers.CloseDB(db)

    _, err = db.Exec("INSERT INTO users (id, username, email) VALUES (?, ?, ?)",
        newUser.ID, newUser.Username, newUser.Email)
    if err != nil {
        return err
    }

    return nil
}
func UpdateUser(updatedUser models.User) error {
    db, err := helpers.OpenDB()
    if err != nil {
        return err
    }
    defer helpers.CloseDB(db)

    result, err := db.Exec("UPDATE users SET username=?, email=? WHERE id=?",
        updatedUser.Username, updatedUser.Email, updatedUser.ID)
    if err != nil {
        return err
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        // Aucune ligne mise à jour, l'élément n'existe probablement pas
        return helpers.ErrNotFound
    }

    return nil
}

func DeleteUser(id uuid.UUID) error {
    db, err := helpers.OpenDB()
    if err != nil {
        return err
    }
    defer helpers.CloseDB(db)

    result, err := db.Exec("DELETE FROM users WHERE id=?", id)
    if err != nil {
        return err
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        // Aucune ligne supprimée, l'élément n'existe probablement pas
        return helpers.ErrNotFound
    }

    return nil
}
