from src.helpers import db

class Rating(db.Model):
    __tablename__ = 'ratings'

    id = db.Column(db.String(255), primary_key=True)
    userId = db.Column(db.String(255), nullable=False)
    songId = db.Column(db.String(255), nullable=False)
    rating = db.Column(db.Integer, nullable=False)
    comments = db.Column(db.String(255), nullable=False)

    def __init__(self, uuid, userId, songId, rating, comments):
        self.id = uuid
        self.userId = userId
        self.songId = songId
        self.rating = rating
        self.comments = comments

    def is_empty(self):
        return (not self.id or self.id == "") and \
               (not self.userId or self.userId == "") and \
               (not self.songId or self.songId == "") and \
               (self.rating is None or self.rating == 0) and \
               (not self.comments or self.comments == "")
