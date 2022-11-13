from django.shortcuts import render, redirect
from django.contrib.auth.decorators import login_required
from .forms import FileUploadForm, BeastUploadForm
from django.http import HttpResponse
from .functions import handle_uploaded_file, handle_beast_file

# Create your views here.

@login_required
def daku_home(request):
    return render(request, 'daku/home.html')

@login_required
def daku_file_upload(request):
    if request.method == "POST":
        form = FileUploadForm(request.POST, request.FILES)
        if form.is_valid():
            handle_uploaded_file(request.FILES["file"], request.user)
            return HttpResponse("File uploaded successfully")
    else:
        form = FileUploadForm()
        return render(request, "daku/file_upload.html", {'form': form})

@login_required
def daku_beast(request):
    if request.method == "POST":
        form = BeastUploadForm(request.POST, request.FILES)
        if form.is_valid():
            handle_beast_file(request.FILES["file"], request.user)
            return HttpResponse("File Uploaded and sent to run")
    else:
        form = BeastUploadForm()
        return render(request, "daku/beast_upload.html", {'form': form})