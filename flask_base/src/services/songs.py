import json
import requests
from marshmallow import EXCLUDE

from src.schemas.song import SongSchema
from src.models.song import Song as SongModel
from src.models.http_exceptions import *
import src.repositories.songs as songs_repository

songs_url = "http://localhost:4000/songs/"  # URL de l'API songs (golang)


def get_song(id):
    response = requests.request(method="GET", url=songs_url+id)
    return response.json(), response.status_code


def create_song(song_data):
    song_model = SongModel.from_dict(song_data)
    song_schema = SongSchema().loads(json.dumps(song_data), unknown=EXCLUDE)

    response = requests.request(method="POST", url=songs_url, json=song_schema)
    if response.status_code != 201:
        return response.json(), response.status_code

    try:
        song_model.id = response.json()["id"]
        songs_repository.add_song(song_model)
    except Exception:
        raise SomethingWentWrong

    return response.json(), response.status_code


def modify_song(id, song_data):
    song_schema = SongSchema().loads(json.dumps(song_data), unknown=EXCLUDE)
    response = None
    if not SongSchema.is_empty(song_schema):
        response = requests.request(method="PUT", url=songs_url+id, json=song_schema)
        if response.status_code != 200:
            return response.json(), response.status_code

    song_model = SongModel.from_dict(song_data)
    if not song_model.is_empty():
        song_model.id = id
        found_song = songs_repository.get_song_from_id(id)
        if not song_model.title:
            song_model.title = found_song.title
        if not song_model.artist:
            song_model.artist = found_song.artist
        if not song_model.release_date:
            song_model.release_date = found_song.release_date
        try:
            songs_repository.update_song(song_model)
        except exc.IntegrityError as e:
            if "NOT NULL" in e.orig.args[0]:
                raise UnprocessableEntity
            raise Conflict

    return (response.json(), response.status_code) if response else get_song(id)
