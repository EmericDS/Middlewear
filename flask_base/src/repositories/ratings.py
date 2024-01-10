from src.helpers import db
from src.models.rating import Rating

def create_rating(rating):
    db.session.add(rating)
    db.session.commit()

def delete_rating(rating_id):
    db.session.query(Rating).filter(Rating.id == rating_id).delete()
    db.session.commit()

def get_rating(rating_id):
    return db.session.query(Rating).get(rating_id)

def get_ratings():
    return db.session.query(Rating).all()

def update_rating(rating_id, updated_rating):
    db.session.query(Rating).filter(Rating.id == rating_id).update(updated_rating)
    db.session.commit()
