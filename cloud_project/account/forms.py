
from django import forms
from django.contrib.auth.forms import AuthenticationForm, UserCreationForm
from .models import UserBase

class UserLoginForm(AuthenticationForm):

    username = forms.CharField(widget=forms.TextInput(
        attrs={'class': 'form-control mb-3', 'placeholder': 'Username', 'id': 'login-username'}))
    password = forms.CharField(widget=forms.PasswordInput(
        attrs={
            'class': 'form-control',
            'placeholder': 'Password',
            'id': 'login-pwd',
        }
    ))

class RegistrationForm(UserCreationForm):
    user_name = forms.CharField(min_length=8, max_length=16)
    email = forms.EmailField(max_length=50)

    class Meta:
        model = UserBase
        fields = ('user_name', 'email', 'password1', 'password2')

    def clean_username(self):
        user_name = self.cleaned_data['user_name'].lower().strip()
        r = UserBase.objects.filter(user_name=user_name)
        if r.count():
            raise forms.ValidationError('Please use another username, that is already taken')
        return user_name

    def clean_email(self):
        email = self.cleaned_data['email'].strip()
        r = UserBase.objects.filter(email=email)
        if r.count():
            raise forms.ValidationError('Please use another Email, that is already taken')
        return email