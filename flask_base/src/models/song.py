from src.helpers import db

class Song(db.Model):
    __tablename__ = 'songs'

    id = db.Column(db.String(255), primary_key=True)
    title = db.Column(db.String(255), nullable=False)
    artist = db.Column(db.String(255), nullable=False)
    genre = db.Column(db.String(255), nullable=False)
    duration = db.Column(db.Integer, nullable=False)

    def __init__(self, uuid, title, artist, genre, duration):
        self.id = uuid
        self.title = title
        self.artist = artist
        self.genre = genre
        self.duration = duration

    def is_empty(self):
        return (not self.id or self.id == "") and \
               (not self.title or self.title == "") and \
               (not self.artist or self.artist == "") and \
               (not self.genre or self.genre == "") and \
               (self.duration is None or self.duration == 0)
