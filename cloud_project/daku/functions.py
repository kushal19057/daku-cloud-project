import docker
import requests
import os

from .models import Container

def handle_uploaded_file(f, user):
    # TODO change static file upload address https://www.javatpoint.com/django-file-upload
    # probably a bad idea to save files like this. add some random hash so that 2 files do not conincide
    print(os.getcwd())
    with open("./static/upload/" + f.name, 'wb') as destination:
        for chunk in f.chunks():
            destination.write(chunk)

    client = docker.from_env()
    c = Container.objects.get(user=user)

    if c is not None:
        container_id = c.container_id
        container = client.containers.get(container_id)
        print(container.ports)
        hostIp = container.ports['8080/tcp'][0]['HostIp']
        hostPort = container.ports['8080/tcp'][0]['HostPort']
        print(container.id + " connection established | status = " + container.status)
        print(hostIp, hostPort)
        
        with open("./static/upload/" + f.name, 'rb') as uf:
            r = requests.post(f"http://{hostIp}:{hostPort}/upload", files={'uploadFile': uf})
            print(r)


def handle_beast_file(f, user):

    with open("./static/upload/" + f.name, 'wb') as destination:
        for chunk in f.chunks():
            destination.write(chunk)

    client = docker.from_env()
    c = Container.objects.get(user=user)

    if c is not None:
        container_id = c.container_id
        container = client.containers.get(container_id)
        print(container.ports)
        hostIp = container.ports['8080/tcp'][0]['HostIp']
        hostPort = container.ports['8080/tcp'][0]['HostPort']
        print(container.id + " connection established | status = " + container.status)
        print(hostIp, hostPort)
        
        with open("./static/upload/" + f.name, 'rb') as uf:
            r = requests.post(f"http://{hostIp}:{hostPort}/upload", files={'uploadFile': uf})
            print(r)
            for out in container.exec_run("bash /app/tmp/beast", stream=True, detach=True, tty=True):
                print(out)