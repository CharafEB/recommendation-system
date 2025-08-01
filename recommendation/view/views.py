from django.shortcuts import HttpResponse

def home(req):
    return HttpResponse("hello world")