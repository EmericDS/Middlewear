package ratings

import (
    "github.com/gofrs/uuid"
    "middleware/example/internal/helpers"
    "middleware/example/internal/models"
)

func GetAllRatings() ([]models.Rating, error) {
    db, err := helpers.OpenDB()
    if err != nil {
        return nil, err
    }
    rows, err := db.Query("SELECT * FROM ratings")
    helpers.CloseDB(db)
    if err != nil {
        return nil, err
    }

    ratingsList := []models.Rating{}
    for rows.Next() {
        var rating models.Rating
        err = rows.Scan(&rating.ID, &rating.UserID, &rating.SongID, &rating.Rating, &rating.Comments)
        if err != nil {
            return nil, err
        }
        ratingsList = append(ratingsList, rating)
    }
    _ = rows.Close()

    return ratingsList, nil
}

func GetRatingByID(id uuid.UUID) (*models.Rating, error) {
    db, err := helpers.OpenDB()
    if err != nil {
        return nil, err
    }
    row := db.QueryRow("SELECT * FROM ratings WHERE id=?", id.String())
    helpers.CloseDB(db)

    var rating models.Rating
    err = row.Scan(&rating.ID, &rating.UserID, &rating.SongID, &rating.Rating, &rating.Comments)
    if err != nil {
        return nil, err
    }
    return &rating, nil
}

func CreateRating(newRating models.Rating) error {
    db, err := helpers.OpenDB()
    if err != nil {
        return err
    }
    defer helpers.CloseDB(db)

    _, err = db.Exec("INSERT INTO ratings (id, user_id, song_id, rating, comments) VALUES (?, ?, ?, ?, ?)",
        newRating.ID, newRating.UserID, newRating.SongID, newRating.Rating, newRating.Comments)
    if err != nil {
        return err
    }

    return nil
}
func UpdateRating(updatedRating models.Rating) error {
    db, err := helpers.OpenDB()
    if err != nil {
        return err
    }
    defer helpers.CloseDB(db)

    result, err := db.Exec("UPDATE ratings SET user_id=?, song_id=?, rating=?, comments=? WHERE id=?",
        updatedRating.UserID, updatedRating.SongID, updatedRating.Rating, updatedRating.Comments, updatedRating.ID)
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

func DeleteRating(id uuid.UUID) error {
    db, err := helpers.OpenDB()
    if err != nil {
        return err
    }
    defer helpers.CloseDB(db)

    result, err := db.Exec("DELETE FROM ratings WHERE id=?", id)
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
