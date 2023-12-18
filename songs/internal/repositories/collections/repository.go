package collections

import (
	"github.com/gofrs/uuid"
	"middleweare/internal/helpers"
	"middleweare/internal/models"
)

func GetAllSongs() ([]models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	rows, err := db.Query("SELECT * FROM songs")
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	songs := []models.Song{}
	for rows.Next() {
		var song models.Song
		err = rows.Scan(&song.ID, &song.Title, &song.Artist, &song.Album, &song.ReleaseYear)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}
	// don't forget to close rows
	_ = rows.Close()

	return songs, nil
}

func GetSongById(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	row := db.QueryRow("SELECT * FROM collections WHERE id=?", id.String())
	

	var song models.Song
	err = row.Scan(&song.ID, &song.Title, &song.Artist, &song.Album, &song.ReleaseYear)
	if err != nil {
		return nil, err
	}
	return &song, err
}

func CreateSong(song models.Song) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	result, err := db.Exec("INSERT INTO songs (id, title, artist, album, release_year) VALUES (?, ?, ?, ?, ?)",
		song.ID, song.Title, song.Artist, song.Album, song.ReleaseYear)
	if err != nil {
		return nil, err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &song, nil
}

func UpdateSong(id uuid.UUID, updatedSong models.Song) (int64, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return 0, err
	}
	defer helpers.CloseDB(db)

	result, err := db.Exec("UPDATE songs SET title=?, artist=?, album=?, release_year=? WHERE id=?",
		updatedSong.Title, updatedSong.Artist, updatedSong.Album, updatedSong.ReleaseYear, id.String())
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func DeleteSong(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM songs WHERE id=?", id.String())
	if err != nil {
		return err
	}

	return nil
}
