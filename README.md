# daku-cloud-project
cloud computing course project DAKU

initial goals :

- [X] `register` new user
  - [X] mail id registration; provide username and password - get this from FCS 
- [X] `login user` and user home page
  - [X] get this from FCS
- [X] `create container` on first login
  - [X] see `os.subprocess` and docker python SDK
- [ ] ability to `upload files` and `view` or `delete` uploaded files
  - [ ] see FCS repo
- [ ] ability to `run` files and stream outputs using any script (beast ?)

---

note : always run the website in private window or incognito mode. 

---

## Steps to run :

- create a new virtual env
- install everything from  `requirements.txt`
- run django server :

`python manage.py runserver`

---
Related to dockerizing go app and running that container from python and interacting with it from postman.

Some important links (kushal) :


- https://stackoverflow.com/questions/69720234/retrieve-the-host-port-of-a-container-launched-from-docker-sdk-for-python
- https://realpython.com/python-subprocess/#basic-usage-of-the-python-subprocess-module
- https://stackoverflow.com/questions/32451748/how-to-bind-ports-with-docker-py
- https://tutorialedge.net/golang/go-docker-tutorial/
- https://stackoverflow.com/questions/39037049/how-to-upload-a-file-and-json-data-in-postman
- https://gabrieltanner.org/blog/golang-file-uploading/
- https://github.com/zupzup/golang-http-file-upload-download/blob/main/main.go
- https://www.zupzup.org/go-http-file-upload-download/
- https://freshman.tech/file-upload-golang/

```
history | grep docker
docker build -t my-go-app .
docker images
docker run -p 8080:8081 -it my-go-app
docker ps
docker stop <container id>
```

---

- [ ] able to run files inside go container. create simple demo app. see online editor code repo for example
- [ ] show file structure and folders (search create dropbox clone using node js)
- [ ] create files using online text editor and save files using that.
- [ ] delete files, move files
- [ ] upload folder
- [ ] Improve UI
- [ ] Run containers on college pc instead of laptop 
- [ ] deploy on PC 1, containers on PC 2, 3
- [ ] 