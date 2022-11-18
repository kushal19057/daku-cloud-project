from django.shortcuts import render
from django.contrib.auth.decorators import login_required
from django.http import HttpResponse
from .functions import get_docker_url


@login_required
def daku_home(request):
    return render(request, 'daku/home.html', {'docker_details': get_docker_url(request.user)})

@login_required
def daku_file_upload(request):
    return render(request, 'daku/file_upload.html', {'docker_details': get_docker_url(request.user)})
    
@login_required
def daku_file_editor(request):
    return render(request, "daku/create_file_using_editor.html", {'docker_details': get_docker_url(request.user)})