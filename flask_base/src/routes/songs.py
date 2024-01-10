import json
from flask import Blueprint, request
from src.models.http_exceptions import *
from src.schemas.song import SongUpdateSchema, SongCreateSchema
from src.schemas.errors import *
import src.services.songs as songs_service

songs = Blueprint(name="songs", import_name=__name__)

@songs.route('/', methods=['POST'])
def CreateSong():
    """
    ---
    post:
      description: Creating a new song
      requestBody:
        required: true
        content:
            application/json:
                schema: SongCreate
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
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
          - songs
    """
    try:
        song_create = SongCreateSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = BadRequestSchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    try:
        return songs_service.CreateSong(song_create)
    except InternalServerError:
        error = InternalServerErrorSchema().loads(json.dumps({"message": "Internal Server Error"}))
        return error, error.get("code")


@songs.route('/<id>', methods=['GET'])
def GetSong(id):
    """
    ---
    get:
      description: Getting a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - songs
    """
    return songs_service.GetSong(id)


@songs.route('/<id>', methods=['PUT'])
def UpdateSong(id):
    """
    ---
    put:
      description: Updating a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      requestBody:
        required: true
        content:
            application/json:
                schema: SongUpdate
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
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
          - songs
    """
    try:
        song_update = SongUpdateSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    try:
        return songs_service.UpdateSong(id, song_update)
    except NotFound:
        error = NotFoundSchema().loads(json.dumps({"message": "Song not found"}))
        return error, error.get("code")
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
        return error, error.get("code")
    except Exception:
        error = InternalServerErrorSchema().loads(json.dumps({"message": "Internal Server Error"}))
        return error, error.get("code")


@songs.route('/<id>', methods=['DELETE'])
def DeleteSong(id):
    """
    ---
    delete:
      description: Deleting a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
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
          - songs
    """
    try:
        return songs_service.DeleteSong(id)
    except NotFound:
        error = NotFoundSchema().loads(json.dumps({"message": "Song not found"}))
        return error, error.get("code")
    except Exception:
        error = InternalServerErrorSchema().loads(json.dumps({"message": "Internal Server Error"}))
        return error, error.get("code")


@songs.route('/', methods=['GET'])
def GetSongs():
    """
    ---
    get:
      description: Getting all songs
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: List[Song]
            application/yaml:
              schema: List[Song]
      tags:
          - songs
    """
    try:
        return songs_service.GetSongs()
    except Exception:
        error = InternalServerErrorSchema().loads(json.dumps({"message": "Internal Server Error"}))
        return error, error.get("code")
