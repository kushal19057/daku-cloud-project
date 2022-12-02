from flask import render_template, flash, redirect, url_for, request
from flaskapp import app, db, bcrypt
from flaskapp.forms import RegistrationForm, LoginForm
from flaskapp.models import User
from flask_login import login_user, current_user, logout_user, login_required


@app.route("/")
@app.route("/home")
def home():
    return render_template("home.html")

@app.route("/profile")
@login_required
def profile():
    return render_template("profile.html")

@app.route("/register", methods=['GET', 'POST'])
def register():
    if current_user.is_authenticated:
        return redirect(url_for('home'))

    form = RegistrationForm()
    if form.validate_on_submit():
        # logic for account creation

        # first hash password
        hashed_password = bcrypt.generate_password_hash(form.password.data).decode('utf-8')

        # create container here
        # TODO

        # create user instance
        user = User(username=form.username.data, email=form.email.data, password=hashed_password, container_id="0")
        db.session.add(user)

        db.session.commit()

        flash(f'Your account has been created. You are now able to login.', "success")

        return redirect(url_for('login'))
    return render_template("register.html", title='Register', form=form)

@app.route("/login", methods=['GET', 'POST'])
def login():
    if current_user.is_authenticated:
        return redirect(url_for('home'))

    form = LoginForm() 
    if form.validate_on_submit():
        user = User.query.filter_by(email=form.email.data).first()
        if user and bcrypt.check_password_hash(user.password, form.password.data):
            login_user(user, remember=form.remember.data)
            next_page = request.args.get('next')
            flash(f'Login Successful', "success")
            return redirect(next_page) if next_page else redirect(url_for('home'))

        else:
            flash(f'Login Unsuccessful. Check email and Password', "danger")

    return render_template("login.html", title='Login', form=form)

@app.route("/logout")
@login_required
def logout():
    logout_user()
    return redirect(url_for('home'))
