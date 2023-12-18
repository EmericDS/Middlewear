package collections

import (
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/collections"
	"net/http"
)

func GetAllSongs() ([]models.Song, error) {
	songs, err := repository.GetAllSongs()
	if err != nil {
		logrus.Errorf("error retrieving songs : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return songs, nil
}

func GetSongById(id uuid.UUID) (*models.Song, error) {
	song, err := repository.GetSongById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "songs not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving songs : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return song, err
}

func CreateSong(song models.Song) (*models.Song, error) {
	createdSong, err := repository.CreateSong(song)
	if err != nil {
		logrus.Errorf("error creating song: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return createdSong, nil
}

func UpdateSong(id uuid.UUID, updatedSong models.Song) (int64, error) {
	rowsAffected, err := repository.UpdateSong(id, updatedSong)
	if err != nil {
		logrus.Errorf("error updating song: %s", err.Error())
		return 0, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return rowsAffected, nil
}

func DeleteSong(id uuid.UUID) error {
	err := repository.DeleteSong(id)
	if err != nil {
		logrus.Errorf("error deleting song: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return nil
}