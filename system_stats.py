# run this file on the server
# https://medium.com/codex/setup-a-python-script-as-a-service-through-systemctl-systemd-f0cc55a42267
# https://stackoverflow.com/questions/64993415/python-infinite-loop-do-something-every-x-seconds-without-sleep

import psutil
import redis
import datetime

IP = '' # change this to the IP of the server on which it is deployed
r = redis.Redis(host='redis-18253.c264.ap-south-1-1.ec2.cloud.redislabs.com', port=18253, db=0, password='oc01To5gRD86hOcEoBD4RMFdpLvk8Ikz')

next_time = datetime.datetime.now()
delta = datetime.timedelta(seconds=2)

while True:
    period = datetime.datetime.now()
    
    # do something here
    
    if period >= next_time:
        # every 30 seconds do something here 1 time, like saving in database
        r.set(f"{IP}-cpu_percent", psutil.cpu_percent())
        r.set(f"{IP}-virtual_memory_percent", psutil.virtual_memory().percent)
        r.set(f"{IP}-disk_usage", psutil.disk_usage('/')[3])
        # print(period)
        next_time += delta
        # Or, depending on how long the main body took
        # and when you want to do this again,
        # next_time = period + delta

