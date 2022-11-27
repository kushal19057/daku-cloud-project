from django.shortcuts import render
from django.contrib.auth.decorators import login_required
from django.http import HttpResponse, JsonResponse
from .functions import get_docker_ip_port
import json

@login_required
def daku_home(request):
    ip, port = get_docker_ip_port(request.user)
    return render(request, 'daku/home.html', {'docker_ip': ip, 'docker_port': port})

@login_required
def daku_file_upload(request):
    ip, port = get_docker_ip_port(request.user)
    return render(request, 'daku/file_upload.html', {"docker_ip": ip, "docker_port": port})
    
@login_required
def daku_file_editor(request):
    ip, port = get_docker_ip_port(request.user)
    return render(request, "daku/beast_run.html", {"docker_ip": ip, "docker_port": port})