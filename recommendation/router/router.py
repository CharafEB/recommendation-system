from django.urls import path
from view import views
urlpatterns = [
    path("" , views.home , name="home")
]
