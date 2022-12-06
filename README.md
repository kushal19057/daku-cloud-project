# daku-cloud-project

some docker commands :
- `docker ps` : list running containers
- `docker ps -al` : list all containers (running or created or exited)
- `docker build -t <name of image> .` : build an image named `<name of image>` from `Dockerfile`
- `docker run -it -d -p :8080 <name of image>` : to run a docker image
- `docker run -it <container_id> bash` : to open a bash shell inside a running docker container

---
- docker build -t "kushal19057/my-go-app" .
