import docker
import json
import random

def get_docker_ip_port(user):
    # ip, port = container.ports['8080/tcp'][0]['HostIp'], container.ports['8080/tcp'][0]['HostPort']
    return user.ip_address, user.port_number

def get_container_ip()->str:
    with open("./flaskapp/ip.json", "r") as f:
        file = json.load(f)

    ips = file["server_ips"]
    return random.choice(ips)