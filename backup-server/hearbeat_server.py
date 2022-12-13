import requests
from multiprocessing import Process
import time
import json
import random
import docker

users = []


RESTORE_PATH = "/app/"

import sqlite3
con = sqlite3.connect("../instance/site.db")


def get_ip_port()->str:
    with open("../flaskapp/ip.json", "r") as f:
        file = json.load(f)

    ips = file["server_ips"]
    return random.choice(ips)


def restore_func(c_id, ip, port, pk_id):
    cur = con.cursor()
    ip = get_ip_port()
    url = ip + ":2375"
    client = docker.DockerClient(base_url=url)
    container = client.containers.run("kushal19057/my-go-app", ports={8080:None}, detach=True)
    container.reload()
    try:
        backup_name = f"{c_id}.tar"
        with open(backup_name, 'rb') as backup_f:
            container.put_archive(RESTORE_PATH, backup_f)
    except:
        print("failed to upload backup data")
    
    port = container.ports['8080/tcp'][0]['HostPort']
    statement = f"""UPDATE user SET container_id='{container.id}', port_number='{port}', ip_address='{ip}' WHERE id={pk_id}"""
    print(statement)
    res = cur.execute(statement)
    print(res)
    con.commit()

    return


while True:
    cur = con.cursor()
    res = cur.execute("SELECT * FROM user")
    users = res.fetchall()

    procs = []
    for user in users:
        print(user)

        user_id = user[0]
        c_id = user[3]
        ip = user[4]
        port = user[5]
        url = f"http://{ip}:{port}/heartbeat"
        try:
            x = requests.get(url, timeout=10)
            print(type(x.status_code), x.status_code)
            if not str(x.status_code) == "202":
                print(f"{user} is down")
                proc = Process(target=restore_func, args=(c_id, ip, port, user_id))
                proc.start()
                procs.append(proc)
            
                   
        except:
            print(f"{user} did not reply back")
            print(f"{user} is down")
            proc = Process(target=restore_func, args=(c_id, ip, port, user_id))
            proc.start()
            procs.append(proc)
            

    for proc in procs:
        print(proc)
        proc.join()

    time.sleep(3600)