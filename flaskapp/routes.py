from flask import render_template, flash, redirect, url_for, request, jsonify
from flaskapp import app, db, bcrypt
from flaskapp.forms import RegistrationForm, LoginForm
from flaskapp.models import User
from flask_login import login_user, current_user, logout_user, login_required
import docker
from flaskapp.functions import get_docker_ip_port

@app.route("/")
@app.route("/home")
def home():
    return render_template("home.html")

@app.route("/profile")
@login_required
def profile():
    docker_ip, docker_port = get_docker_ip_port(current_user)
    return render_template("profile.html", docker_ip=docker_ip, docker_port=docker_port)

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
        client = docker.from_env()
        container = client.containers.run("my-go-app", ports={8080:None}, detach=True)
        container.reload()

        # create user instance
        user = User(username=form.username.data, email=form.email.data, password=hashed_password, container_id=container.id)
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


# routes for container work
@app.route("/upload")
@login_required
def upload_file():
    docker_ip, docker_port = get_docker_ip_port(current_user)
    return render_template("container_file_upload.html", docker_ip=docker_ip, docker_port=docker_port)

@app.route("/files")
@login_required
def list_work_files():
    docker_ip, docker_port = get_docker_ip_port(current_user)
    return render_template("container_list_work_files.html", docker_ip=docker_ip, docker_port=docker_port)

@app.route("/beast")
@login_required
def run_beast():
    docker_ip, docker_port = get_docker_ip_port(current_user)
    return render_template("container_run_beast.html", docker_ip=docker_ip, docker_port=docker_port)