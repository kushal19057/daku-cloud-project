from flask_wtf import FlaskForm
from wtforms import StringField, PasswordField, SubmitField
from wtforms.validators import DataRequired, Email, EqualTo
from flaskapp.models import User


class RegistrationForm(FlaskForm):
    email = StringField(
        "Email", 
        validators=[DataRequired(), Email()]
        )
    password = PasswordField(
        'Password', 
        validators=[DataRequired()]
        )
    confirm_password = PasswordField(
        'Confirm Password', 
        validators=[DataRequired(), EqualTo('password')]
        )
    submit = SubmitField("Sign Up")


class LoginForm(FlaskForm):
    email = StringField(
        "Email", 
        validators=[DataRequired(), Email()]
        )
    password = PasswordField(
        'Password', 
        validators=[DataRequired()]
        )
    submit = SubmitField("Login")

