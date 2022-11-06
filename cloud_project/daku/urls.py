from django.urls import path
from .views import home

app_name = "daku"
urlpatterns = [
    path('', home, name='daku-home')
]