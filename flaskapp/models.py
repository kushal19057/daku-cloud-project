from flaskapp import db, login_manager
from flask_login import UserMixin

@login_manager.user_loader
def load_user(user_id):
    return User.query.get(int(user_id))

class User(db.Model, UserMixin):
    id = db.Column(db.Integer, primary_key=True)
    email = db.Column(db.String(120), unique=True, nullable=False)
    password = db.Column(db.String(60), nullable=False)
    container_id = db.Column(db.String(120), nullable=False)
    ip_address = db.Column(db.String(30), nullable=False)
    port_number = db.Column(db.Integer, nullable=False)

    def __repr__(self):
        return f"User('{self.email}, {self.container_id}, {self.ip_address}, {self.port_number}')"

