from flask_login import UserMixin
from werkzeug.security import generate_password_hash, check_password_hash
from src.helpers import db

class User(UserMixin, db.Model):
    __tablename__ = 'users'

    id = db.Column(db.String(255), primary_key=True)
    username = db.Column(db.String(255), unique=True, nullable=False)
    encrypted_password = db.Column(db.String(255), nullable=False)

    def __init__(self, uuid, username, password):
        self.id = uuid
        self.username = username
        self.set_password(password)

    def is_empty(self):
        return (not self.id or self.id == "") and \
               (not self.username or self.username == "") and \
               (not self.encrypted_password or self.encrypted_password == "")

    def set_password(self, password):
        self.encrypted_password = generate_password_hash(password)

    def check_password(self, password):
        return check_password_hash(self.encrypted_password, password)

    @staticmethod
    def from_dict_with_clear_password(obj):
        username = obj.get("username") if obj.get("username") != "" else None
        password = generate_password_hash(obj.get("password")) if obj.get("password") != "" else None
        return User(None, username, password)
