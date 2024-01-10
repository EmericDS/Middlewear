import json
from flask import Blueprint, request, jsonify
from marshmallow import ValidationError
from flask_login import login_required
from src.models.http_exceptions import *
from src.schemas.user import UserUpdateSchema
from src.schemas.errors import *
from src.services.collections import (
    create_user,
    delete_user,
    get_user,
    get_users,
    update_user,
)

# from routes import users
users = Blueprint(name="users", import_name=__name__)


@users.route('/<id>', methods=['GET'])
@login_required
def get_user(id):
    """
    ---
    get:
      description: Getting a user
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of user id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: User
            application/yaml:
              schema: User
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - users
    """
    # Appeler la fonction correspondante de l'API Users pour obtenir un utilisateur par ID
    user = get_user(id)
    if user:
        return jsonify(user), 200
    else:
        return jsonify({"message": "User not found"}), 404


@users.route('/<id>', methods=['PUT'])
@login_required
def update_user(id):
    """
    ---
    put:
      description: Updating a user
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of user id
      requestBody:
        required: true
        content:
            application/json:
                schema: UserUpdate
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: User
            application/yaml:
              schema: User
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
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
          - users
    """
    # Parser le corps de la requête
    try:
        user_update = UserUpdateSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    # Modification de l'utilisateur (username, nom, mot de passe, etc.)
    try:
        return update_user(id, user_update)
    except Conflict:
        error = ConflictSchema().loads(json.dumps({"message": "User already exists"}))
        return error, error.get("code")
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
        return error, error.get("code")
    except Forbidden:
        error = ForbiddenSchema().loads(json.dumps({"message": "Can't manage other users"}))
        return error, error.get("code")
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")


@users.route('/', methods=['POST'])
def create_user():
    """
    ---
    post:
      description: Creating a new user
      requestBody:
        required: true
        content:
          application/json:
            schema: User
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: User
            application/yaml:
              schema: User
        '400':
          description: Bad request
          content:
            application/json:
              schema: BadRequest
            application/yaml:
              schema: BadRequest
        '409':
          description: Conflict
          content:
            application/json:
              schema: Conflict
            application/yaml:
              schema: Conflict
      tags:
          - users
    """
    try:
        return users_service.create_user(request)
    except BadRequest:
        error = BadRequestSchema().loads(json.dumps({"message": "Bad Request"}))
        return error, error.get("code")
    except Conflict:
        error = ConflictSchema().loads(json.dumps({"message": "User already exists"}))
        return error, error.get("code")
    except Exception:
        error = InternalServerErrorSchema().loads(json.dumps({"message": "Internal Server Error"}))
        return error, error.get("code")


@users.route('/<id>', methods=['DELETE'])
def delete_user(id):
    """
    ---
    delete:
      description: Deleting a user
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of user id
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
        '500':
          description: Internal server error
          content:
            application/json:
              schema: InternalServerError
            application/yaml:
              schema: InternalServerError
      tags:
          - users
    """
    try:
        return users_service.delete_user(id)
    except NotFound:
        error = NotFoundSchema().loads(json.dumps({"message": "User not found"}))
        return error, error.get("code")
    except Exception:
        error = InternalServerErrorSchema().loads(json.dumps({"message": "Internal Server Error"}))
        return error, error.get("code")


@users.route('/', methods=['GET'])
def get_users():
    """
    ---
    get:
      description: Getting all users
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: List[User]
            application/yaml:
              schema: List[User]
        '500':
          description: Internal server error
          content:
            application/json:
              schema: InternalServerError
            application/yaml:
              schema: InternalServerError
      tags:
          - users
    """
    try:
        return users_service.get_users()
    except Exception:
        error = InternalServerErrorSchema().loads(json.dumps({"message": "Internal Server Error"}))
        return error, error.get("code")
