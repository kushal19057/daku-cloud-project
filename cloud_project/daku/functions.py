import docker
import requests
import os

from .models import Container


def get_docker_ip_port(user):
    # using remote containers now
    c = Container.objects.get(user=user)
    
    client = docker.DockerClient(base_url='tcp://192.168.192.148:2375', tls=False, version='auto')

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


def upload_file_to_container(user, file):
    ip, port = get_docker_ip_port(user)
    post_data = {'uploadFile': file}
    url = f"http://{ip}:{port}/upload_file"
    response = requests.post(url, data=post_data)
    print(response)
    return response.content

def run_beast_on_container(user, files):
    pass

def get_available_docker_ip_port():
    """
    returns ip, port
    """
    return "localhost", "2375"