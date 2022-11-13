from django.shortcuts import render, redirect
from django.contrib.auth.decorators import login_required
from .forms import FileUploadForm
from django.http import HttpResponse
from .functions import handle_uploaded_file

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