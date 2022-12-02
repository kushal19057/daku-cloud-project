from flask import Flask, render_template, flash, redirect, url_for
from forms import RegistrationForm, LoginForm

app = Flask(__name__)

app.config["SECRET_KEY"] = '1f1e1210f56638c8d6536435c6f19b5a'

@app.route("/")
@app.route("/home")
def home():
    return render_template("home.html")

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