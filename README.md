# daku-cloud-project

## Setup

### Master `flaskapp` server

The file `flaskapp/ip.json` should contain the list of servers having public IPs. The docker containers will be deployed on these servers.

```json
{
  "server_ips": [
	  "127.0.0.1"
  ]
}
```


```
virtualenv cloud-env
source cloud-env/bin/activate
cd flaskapp
pip install -r requirements.txt

export FLASK_APP=run.py
```

#### setup database by initialising a python shell

```
>>> python

>>> from flaskapp import app, db
>>> with app.app_context():
>>>    db.create_all()
```

#### Measure system utlisation in real time

```
specifiy IP address of  currrent server as :
IP = ''
```

```
# run in background
python system_stats.py
```

### `goapp` Docker Container image

```
cd goapp

# build Dockerfile image
docker build -t kushal19057/my-go-app .
```

### Backup Server

```
# run as background process
python take_backup.py

# run as background process
python heartbeat_server.py
```

### Other useful docker commands to monitor system :

- `docker ps` : list running containers
- `docker ps -al` : list all containers (running or created or exited)
- `docker build -t <name of image> .` : build an image named `<name of image>` from `Dockerfile`
- `docker run -it -d -p :8080 <name of image>` : to run a docker image
- `docker build -t "kushal19057/my-go-app" .` : to build dockerfile
- `docker exec -it <container id>  bash` : to open a shell inside docker container

---

Enjoy using daku cloud app :)