from django.contrib.auth import views as auth_views
from django.urls import path
from . import views
from .forms import UserLoginForm

app_name = "account"

urlpatterns = [
    path("login/", auth_views.LoginView.as_view(template_name='account/registration/login.html', form_class=UserLoginForm, next_page="/account/dashboard"), name='account-login'),
    path("logout/", auth_views.LogoutView.as_view(next_page='/account/login/'), name='account-logout'),
    path("register/", views.account_register, name='account-register'),
    # user register
    path("dashboard/", views.dashboard, name='account-dashboard'),
]