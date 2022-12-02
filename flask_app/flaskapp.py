from flask import Flask, render_template, flash, redirect, url_for
from flask_sqlalchemy import SQLAlchemy
from forms import RegistrationForm, LoginForm
from datetime import datetime

app = Flask(__name__)

app.config["SECRET_KEY"] = '1f1e1210f56638c8d6536435c6f19b5a'
app.config["SQLALCHEMY_DATABASE_URI"] = 'sqlite:///site.db'

db = SQLAlchemy(app)

class User(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    username = db.Column(db.String(20), unique=True, nullable=False)
    email = db.Column(db.String(120), unique=True, nullable=False)
    password = db.Column(db.String(60), nullable=False)
    container_id = db.Column(db.Integer, db.ForeignKey('container.id'), nullable=False)
    plan = db.Column(db.String(20), nullable=False, default="BASIC")

    def __repr__(self):
        return f"User('{self.username}', '{self.email}')"

class Container(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    container_id = db.Column(db.String(120), unique=True, nullable=False)
    time_created = db.Column(db.DateTime, nullable=False, default=datetime.utcnow)
    user = db.relationship('User', backref='owner', lazy=True)

    def __repr__(self):
        return f"Container('{self.container_id}', '{self.user}')"


@app.route("/")
@app.route("/home")
def home():
    return render_template("home.html")

@app.route("/about")
def about():
    return render_template("about.html")

@app.route("/register", methods=['GET', 'POST'])
def register():
    form = RegistrationForm()
    if form.validate_on_submit():
        flash(f'Account created for {form.username.data}!', "success")
        return redirect(url_for('home'))
    return render_template("register.html", title='Register', form=form)

@app.route("/login", methods=['GET', 'POST'])
def login():
    form = LoginForm() 
    if form.validate_on_submit():
        if form.email.data == "admin@blog.com" and form.password.data == 'password':
            flash(f'You have been logged in!', "success")
            return redirect(url_for('home'))
        else:
            flash(f'Login Unsuccessful. Check Username and Password', "danger")

    return render_template("login.html", title='Login', form=form)

if __name__ == '__main__':
    app.run(debug=True)