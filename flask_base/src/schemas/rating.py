from marshmallow import Schema, fields, validates_schema, ValidationError

# Schéma de notation de sortie (renvoyé au front)
class RatingSchema(Schema):
    id = fields.String(description="UUID")
    user_id = fields.String(description="User UUID")
    song_id = fields.String(description="Song UUID")
    rating = fields.Integer(description="Rating")
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("user_id") or obj.get("user_id") == "") and \
               (not obj.get("song_id") or obj.get("song_id") == "") and \
               (not obj.get("rating") or obj.get("rating") == "")


class BaseRatingSchema(Schema):
    user_id = fields.String(description="User UUID")
    song_id = fields.String(description="Song UUID")
    rating = fields.Integer(description="Rating")


# Schéma de notation de modification (user_id, song_id, rating)
class RatingUpdateSchema(BaseRatingSchema):
    # permet de définir dans quelles conditions le schéma est validé ou non
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("user_id" in data and data["user_id"] != "") or
                ("song_id" in data and data["song_id"] != "") or
                ("rating" in data and data["rating"] != "")):
            raise ValidationError("at least one of ['user_id','song_id','rating'] must be specified")
