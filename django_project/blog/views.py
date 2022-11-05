from django.shortcuts import render
# Create your views here.

# logic for how to handle routes comes here

# create some fake posts
posts = [
    {
        'author': 'coreyms',
        'title': 'blog post 1',
        'content': 'first post content',
        'date': 'aug 27, 2018'
    },
    {
        'author': 'coreyms2',
        'title': 'blog post 2',
        'content': 'second post content',
        'date': 'aug 30, 2018'
    }
]

def home(request):
    context = {
        'title': 'Wow!',
        'posts': posts
    }

    return render(request, 'blog/home.html', context)

def about(request):
    return render(request, 'blog/about.html')