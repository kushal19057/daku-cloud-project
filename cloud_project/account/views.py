from django.shortcuts import render, redirect
from django.contrib import messages
from django.contrib.auth.decorators import login_required
from django.urls import reverse


from .forms import RegistrationForm
from .models import UserBase
# Create your views here.
@login_required
def dashboard(request):
    return render(request, 'account/dashboard/dashboard.html', {'section': 'profile'})

def account_register(request):
    if request.user.is_authenticated:
        return redirect('account:account-dashboard')

    registerForm = RegistrationForm(request.POST)

    if registerForm.is_valid():
    
        user = registerForm.save(commit=False)
        user.user_name = registerForm.clean_username()
        user.set_password(registerForm.cleaned_data['password1'])
        
        # create container here
            
        user.save()
        messages.success(request, f'Account created for {user.user_name}! Try logging in!')
        
        return render(request, 'account/registration/register.html', {'usr': user})

    else:
        registerForm = RegistrationForm()
    
    return render(request, 'account/registration/register.html', {'form': registerForm})