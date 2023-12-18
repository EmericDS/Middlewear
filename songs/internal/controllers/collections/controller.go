package collections

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleweare/internal/models"
	"middleweare/internal/repositories/songs"
	"net/http"
)

// GetCollection
// @Tags         collections
// @Summary      Get a collection.
// @Description  Get a collection.
// @Param        id           	path      string  true  "Collection UUID formatted ID"
// @Success      200            {object}  models.Collection
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /collections/{id} [get]
// GetSongs handles GET /songs
func GetSongs(w http.ResponseWriter, r *http.Request) {
	songs, err := songs.GetAllSongs()
	if err != nil {
		logrus.Errorf("error retrieving songs: %s", err.Error())
		handleError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, songs)
}

// GetSong handles GET /songs/{id}
func GetSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songID, _ := ctx.Value("songID").(uuid.UUID)

	song, err := songs.GetSongById(songID)
	if err != nil {
		logrus.Errorf("error retrieving song: %s", err.Error())
		handleError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, song)
}

// CreateSong handles POST /songs
func CreateSong(w http.ResponseWriter, r *http.Request) {
	var song models.Song
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		logrus.Errorf("error decoding request body: %s", err.Error())
		handleError(w, &models.CustomError{
			Message: "Invalid request body",
			Code:    http.StatusBadRequest,
		})
		return
	}

	newSong, err := songs.CreateSong(song)
	if err != nil {
		logrus.Errorf("error creating song: %s", err.Error())
		handleError(w, err)
		return
	}

	respondWithJSON(w, http.StatusCreated, newSong)
}

// UpdateSong handles PUT /songs/{id}
func UpdateSong(w http.ResponseWriter, r *http.Request) {
	var updatedSong models.Song
	if err := json.NewDecoder(r.Body).Decode(&updatedSong); err != nil {
		logrus.Errorf("error decoding request body: %s", err.Error())
		handleError(w, &models.CustomError{
			Message: "Invalid request body",
			Code:    http.StatusBadRequest,
		})
		return
	}

	ctx := r.Context()
	songID, _ := ctx.Value("songID").(uuid.UUID)

	_, err := songs.UpdateSong(songID, updatedSong)
	if err != nil {
		logrus.Errorf("error updating song: %s", err.Error())
		handleError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedSong)
}

// DeleteSong handles DELETE /songs/{id}
func DeleteSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songID, _ := ctx.Value("songID").(uuid.UUID)

	err := songs.DeleteSong(songID)
	if err != nil {
		logrus.Errorf("error deleting song: %s", err.Error())
		handleError(w, err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

func handleError(w http.ResponseWriter, err error) {
	customError, isCustom := err.(*models.CustomError)
	if isCustom {
		w.WriteHeader(customError.Code)
		respondWithJSON(w, customError.Code, customError)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		respondWithJSON(w, http.StatusInternalServerError, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		})
	}
}
