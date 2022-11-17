from django import forms

class FileUploadForm(forms.Form):
    filepath = forms.CharField(label="enter full file path")
    file = forms.FileField()
