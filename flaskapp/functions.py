import docker

def get_docker_ip_port(user):
    client = docker.from_env()
    container = client.containers.get(user.container_id)
    ip, port = container.ports['8080/tcp'][0]['HostIp'], container.ports['8080/tcp'][0]['HostPort']
    return ip, port