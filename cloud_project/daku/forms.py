from django import forms

class FileUploadForm(forms.Form):
    filename = forms.CharField(label="enter file name")
    filepath = forms.CharField(label="enter file path")
    file = forms.FileField()

class BeastUploadForm(forms.Form):
    file = forms.FileField()