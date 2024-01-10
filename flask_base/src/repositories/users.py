from src.helpers import db
from src.models.user import User

def create_user(user):
    db.session.add(user)
    db.session.commit()

def delete_user(user_id):
    db.session.query(User).filter(User.id == user_id).delete()
    db.session.commit()

def get_user(user_id):
    return db.session.query(User).get(user_id)

def get_users():
    return db.session.query(User).all()

def update_user(user_id, updated_user):
    db.session.query(User).filter(User.id == user_id).update(updated_user)
    db.session.commit()
