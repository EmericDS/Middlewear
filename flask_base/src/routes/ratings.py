import json
from flask import Blueprint, request
from src.models.http_exceptions import *
from src.schemas.rating import RatingUpdateSchema, RatingCreateSchema
from src.schemas.errors import *
import src.services.ratings as ratings_service

ratings = Blueprint(name="ratings", import_name=__name__)

@ratings.route('/', methods=['POST'])
def CreateRating():
    """
    ---
    post:
      description: Creating a new rating
      requestBody:
        required: true
        content:
            application/json:
                schema: RatingCreate
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
        '400':
          description: Bad request
          content:
            application/json:
              schema: BadRequest
            application/yaml:
              schema: BadRequest
        '500':
          description: Internal server error
          content:
            application/json:
              schema: InternalServerError
            application/yaml:
              schema: InternalServerError
      tags:
          - ratings
    """
    try:
        rating_create = RatingCreateSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = BadRequestSchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    try:
        return ratings_service.CreateRating(rating_create)
    except InternalServerError:
        error = InternalServerErrorSchema().loads(json.dumps({"message": "Internal Server Error"}))
        return error, error.get("code")


@ratings.route('/<id>', methods=['GET'])
def GetRating(id):
    """
    ---
    get:
      description: Getting a rating
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - ratings
    """
    return ratings_service.GetRating(id)


@ratings.route('/<id>', methods=['PUT'])
def UpdateRating(id):
    """
    ---
    put:
      description: Updating a rating
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      requestBody:
        required: true
        content:
            application/json:
                schema: RatingUpdate
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - ratings
    """
    try:
        rating_update = RatingUpdateSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    try:
        return ratings_service.UpdateRating(id, rating_update)
    except NotFound:
        error = NotFoundSchema().loads(json.dumps({"message": "Rating not found"}))
        return error, error.get("code")
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
        return error, error.get("code")
    except Exception:
        error = InternalServerErrorSchema().loads(json.dumps({"message": "Internal Server Error"}))
        return error, error.get("code")


@ratings.route('/<id>', methods=['DELETE'])
def DeleteRating(id):
    """
    ---
    delete:
      description: Deleting a rating
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      responses:
        '204':
          description: No content
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - ratings
    """
    try:
        return ratings_service.DeleteRating(id)
    except NotFound:
        error = NotFoundSchema().loads(json.dumps({"message": "Rating not found"}))
        return error, error.get("code")
    except Exception:
        error = InternalServerErrorSchema().loads(json.dumps({"message": "Internal Server Error"}))
        return error, error.get("code")


@ratings.route('/', methods=['GET'])
def GetRatings():
    """
    ---
    get:
      description: Getting all ratings
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: List[Rating]
            application/yaml:
              schema: List[Rating]
        '500':
          description: Internal server error
          content:
            application/json:
              schema: InternalServerError
            application/yaml:
              schema: InternalServerError
      tags:
          - ratings
    """
    try:
        return ratings_service.GetRatings()
    except InternalServerError:
        error = InternalServerErrorSchema().loads(json.dumps({"message": "Internal Server Error"}))
        return error, error.get("code")
