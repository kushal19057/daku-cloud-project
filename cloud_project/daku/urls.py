from django.urls import path
from .views import daku_home, daku_file_upload

app_name = "daku"

urlpatterns = [
    path('', daku_home, name='daku-home'),
    path('upload/', daku_file_upload, name='daku-file-upload')
]