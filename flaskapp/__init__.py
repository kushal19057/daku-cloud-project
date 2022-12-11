from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from flask_bcrypt import Bcrypt
from flask_login import LoginManager
import redis

app = Flask(__name__)

app.config["SECRET_KEY"] = '1f1e1210f56638c8d6536435c6f19b5a'
app.config["SQLALCHEMY_DATABASE_URI"] = 'sqlite:///site.db'
# app.config["SQLALCHEMY_DATABASE_URI"] = f"postgresql://{username}:{password}@localhost:5432/{database}"


db = SQLAlchemy(app)
bcrypt = Bcrypt(app)
login_manager = LoginManager(app)
login_manager.login_view = 'login'
login_manager.login_message_category = 'info'

r = redis.Redis(host='redis-18253.c264.ap-south-1-1.ec2.cloud.redislabs.com', port=18253, db=0, password='oc01To5gRD86hOcEoBD4RMFdpLvk8Ikz')

from flaskapp import routes

