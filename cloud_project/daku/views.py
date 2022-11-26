from django.shortcuts import render
from django.contrib.auth.decorators import login_required
from django.http import HttpResponse, JsonResponse
from .functions import get_docker_ip_port, upload_file_to_container, run_beast_on_container
import json

@login_required
def daku_home(request):
    ip, port = get_docker_ip_port(request.user)
    return render(request, 'daku/home.html', {'docker_ip': ip, 'docker_port': port})

@login_required
def daku_file_upload(request):
    # handle POST request
    if request.method == "POST":
        # upload_file_to_container(request.user, request.body['upload_file'])
        data = {"message": "SUCCESS",}
        return JsonResponse(data)
    else:
        # handle GET request
        return render(request, 'daku/file_upload.html')
    
@login_required
def daku_file_editor(request):
    # handle POST request
    if request.method == "POST":
        run_beast_on_container(request.user, request.FILES)
        data = {"message": "SUCCESS"}
        return JsonResponse(data)
    else:
        # handle GET request
        return render(request, "daku/beast_run.html")