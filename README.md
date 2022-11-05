# daku-cloud-project
cloud computing course project DAKU


---

```
Commands run :
django-admin startproject django_project

to run web server on localserver :

python manage.py runserver

to create an app :

python manange.py startapp blog

in order to map urls to functions, create a new file `urls.py` in the `blog` directory
touch blog/urls.py

create a `templates` directory inside the `blog` directory
django_project/ mkdir blog/templates

create a `blog` directory inside the `templates` directory
django_project/ mkdir blog/templates/blog

create 2 html files `home.html` and `about.html`

Add blog app to the list of installed apps in settings.py

create new html file `base.html`; this will be used for template inheritance

```