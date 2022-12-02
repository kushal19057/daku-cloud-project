from flask_wtf import FlaskForm
from wtforms import StringField, PasswordField, SubmitField, BooleanField
from wtforms.validators import DataRequired, Length, Email, EqualTo, ValidationError
from flaskapp.models import User


class RegistrationForm(FlaskForm):
    username = StringField(
        "Username", 
        validators=[DataRequired(), Length(min=6, max=15)]
        )
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


    def validate_username(self, username):
        user = User.query.filter_by(username=username.data).first()
        print(user)
        if user:
            return ValidationError("That username is already taken. Please choose a different username.")
    
    def validate_email(self, email):
        user = User.query.filter_by(email=email.data).first()
        if user:
            return ValidationError("That email is already taken. Please choose a different email.")


class LoginForm(FlaskForm):
    email = StringField(
        "Email", 
        validators=[DataRequired(), Email()]
        )
    password = PasswordField(
        'Password', 
        validators=[DataRequired()]
        )
    remember = BooleanField("Remember Me")
    submit = SubmitField("Login")

