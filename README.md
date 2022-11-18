# daku-cloud-project

## How to run the code for the first time ?

- ensure that a python virtualenv is setup and docker setup is done on PC
- in the folder where  `manage.py` is present, run :
  - `python manage.py makemigrations`
  - `python manage.py migrate --run-syncdb`
- the above commands will setup the database and stuff. Now to run the django webserver, run `python manage.py runserver`.

---

some docker commands :
- `docker ps` : list running containers
- `docker ps -al` : list all containers (running or created or exited)
- `docker build -t <name of image> .` : build an image named `<name of image>` from `Dockerfile`
- `docker run -it -d -p :8080 <name of image>` : to run a docker image