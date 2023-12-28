package users

import (
    "github.com/gofrs/uuid"
    "middleware/example/internal/helpers"
    "middleware/example/internal/models"
)

func GetAllUsers() ([]models.User, error) {
    db, err := helpers.OpenDB()
    if err != nil {
        return nil, err
    }
    rows, err := db.Query("SELECT * FROM users")
    helpers.CloseDB(db)
    if err != nil {
        return nil, err
    }

    usersList := []models.User{}
    for rows.Next() {
        var user models.User
        err = rows.Scan(&user.ID, &user.Username, &user.Email)
        if err != nil {
            return nil, err
        }
        usersList = append(usersList, user)
    }
    _ = rows.Close()

    return usersList, nil
}

func GetUserByID(id uuid.UUID) (*models.User, error) {
    db, err := helpers.OpenDB()
    if err != nil {
        return nil, err
    }
    row := db.QueryRow("SELECT * FROM users WHERE id=?", id.String())
    helpers.CloseDB(db)

    var user models.User
    err = row.Scan(&user.ID, &user.Username, &user.Email)
    if err != nil {
        return nil, err
    }
    return &user, nil
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
