from django.contrib import admin

# Register your models here.
# accounts/admin.py

from django.contrib import admin
from .models import UserProfile

admin.site.register(UserProfile)
