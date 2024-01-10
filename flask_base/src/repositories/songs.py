from src.helpers import db
from src.models.song import Song

def create_song(song):
    db.session.add(song)
    db.session.commit()

def delete_song(song_id):
    db.session.query(Song).filter(Song.id == song_id).delete()
    db.session.commit()

def get_song(song_id):
    return db.session.query(Song).get(song_id)

def get_songs():
    return db.session.query(Song).all()

def update_song(song_id, updated_song):
    db.session.query(Song).filter(Song.id == song_id).update(updated_song)
    db.session.commit()
