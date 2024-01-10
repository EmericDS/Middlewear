from marshmallow import Schema, fields, validates_schema, ValidationError

# Schéma de chanson de sortie (renvoyé au front)
class SongSchema(Schema):
    id = fields.String(description="UUID")
    title = fields.String(description="Title")
    artist = fields.String(description="Artist")
    release_date = fields.Date(description="Release date")
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("title") or obj.get("title") == "") and \
               (not obj.get("artist") or obj.get("artist") == "") and \
               (not obj.get("release_date") or obj.get("release_date") == "")


class BaseSongSchema(Schema):
    title = fields.String(description="Title")
    artist = fields.String(description="Artist")
    release_date = fields.Date(description="Release date")


# Schéma de chanson de modification (title, artist, release_date)
class SongUpdateSchema(BaseSongSchema):
    # permet de définir dans quelles conditions le schéma est validé ou non
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("title" in data and data["title"] != "") or
                ("artist" in data and data["artist"] != "") or
                ("release_date" in data and data["release_date"] != "")):
            raise ValidationError("at least one of ['title','artist','release_date'] must be specified")
