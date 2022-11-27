import docker
import requests
import os

from .models import Container


def get_docker_ip_port(user):
    # using remote containers now
    c = Container.objects.get(user=user)
    ip, port = get_available_docker_ip_port()
    url = f"tcp://{ip}:{port}"
    # client = docker.from_env()
    client = docker.DockerClient(base_url=url, tls=False, version='auto')

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

def get_available_docker_ip_port():
    """
    returns ip, port
    """
    return "localhost", "2375"