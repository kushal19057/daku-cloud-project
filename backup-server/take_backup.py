import time
import docker
users = []

import sqlite3
con = sqlite3.connect("../instance/site.db")

BACKUP_PATH = "/app/tmp"

while True:
    cur = con.cursor()
    res = cur.execute("SELECT * FROM user")
    users = res.fetchall()
    print(users)

    for user in users:
        print(user)
        c_id = user[3]
        ip = user[4]

        url = ip + ":2375"

        try:
            client = docker.DockerClient(base_url=url, timeout=5)

            print(c_id, ip, url)
            container = client.containers.get(c_id)
            
            if container.status == "running":
                # take backup and override prev backup
                f = open(f"{c_id}.tar", "wb")
                bits, stat = container.get_archive(BACKUP_PATH)
                print(stat)
                for chunk in bits:
                    f.write(chunk)
                f.close()
        except:
            print("cannot connect to docker container")
    
    time.sleep(3600)