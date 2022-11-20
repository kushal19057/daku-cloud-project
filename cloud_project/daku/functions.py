import docker
import requests
import os

from .models import Container


def get_docker_ip_port(user):
    client = docker.from_env()
    c = Container.objects.get(user=user)

    ip = None
    port = None

    if c is not None:
        container_id = c.container_id
        container = client.containers.get(container_id)
        print(container.ports)
        hostIp = container.ports['8080/tcp'][0]['HostIp']
        hostPort = container.ports['8080/tcp'][0]['HostPort']
        print(container.id + " connection established | status = " + container.status)
        print(hostIp, hostPort)

        ip = hostIp
        port = hostPort
    
    return ip, port