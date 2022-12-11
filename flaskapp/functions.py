import docker
import json
import random
from flaskapp import r


inf=int(1e9)


def get_docker_ip_port(user):
    # ip, port = container.ports['8080/tcp'][0]['HostIp'], container.ports['8080/tcp'][0]['HostPort']
    return user.ip_address, user.port_number

def get_container_ip()->str:
    with open("./flaskapp/ip.json", "r") as f:
        file = json.load(f)

    ips = file["server_ips"]

    dict_load={}

    for ip in ips:
        dict_load[ip]=inf
    

    for ip in ips:
        cpu_percent=float(r.get(f"{ip}-cpu_percent"))
        virtual_memory=float(r.get(f"{ip}-virtual_memory_percent"))
        dict_load[ip]=cpu_percent*virtual_memory


    dict_load=dict(sorted(dict_load.items(), key=lambda item: item[1]))

    print(list(dict_load.keys())[0])
    return str(list(dict_load.keys())[0])