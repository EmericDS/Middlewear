# ratings.py
import json
import requests
from marshmallow import EXCLUDE

from src.schemas.rating import RatingSchema
from src.models.rating import Rating as RatingModel
from src.models.http_exceptions import *
import src.repositories.ratings as ratings_repository

ratings_url = "http://localhost:4000/ratings/"  # URL de l'API ratings (golang)


def get_rating(id):
    response = requests.request(method="GET", url=ratings_url+id)
    return response.json(), response.status_code


def create_rating(rating_data):
    rating_model = RatingModel.from_dict(rating_data)
    rating_schema = RatingSchema().loads(json.dumps(rating_data), unknown=EXCLUDE)

    response = requests.request(method="POST", url=ratings_url, json=rating_schema)
    if response.status_code != 201:
        return response.json(), response.status_code

    try:
        rating_model.id = response.json()["id"]
        ratings_repository.add_rating(rating_model)
    except Exception:
        raise SomethingWentWrong

    return response.json(), response.status_code


def modify_rating(id, rating_data):
    rating_schema = RatingSchema().loads(json.dumps(rating_data), unknown=EXCLUDE)
    response = None
    if not RatingSchema.is_empty(rating_schema):
        response = requests.request(method="PUT", url=ratings_url+id, json=rating_schema)
        if response.status_code != 200:
            return response.json(), response.status_code

    rating_model = RatingModel.from_dict(rating_data)
    if not rating_model.is_empty():
        rating_model.id = id
        found_rating = ratings_repository.get_rating_from_id(id)
        if not rating_model.user_id:
            rating_model.user_id = found_rating.user_id
        if not rating_model.song_id:
            rating_model.song_id = found_rating.song_id
        if not rating_model.rating:
            rating_model.rating = found_rating.rating
        try:
            ratings_repository.update_rating(rating_model)
        except exc.IntegrityError as e:
            if "NOT NULL" in e.orig.args[0]:
                raise UnprocessableEntity
            raise Conflict

    return (response.json(), response.status_code) if response else get_rating(id)
