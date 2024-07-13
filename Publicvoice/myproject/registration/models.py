from django.db import models

# Create your models here.
# accounts/models.py

from django.db import models
from django.contrib.auth.models import User

class UserProfile(models.Model):
    CATEGORY_CHOICES = [
        ('president', 'President'),
        ('governor', 'Governor'),
        ('mp', 'MP'),
        ('senator', 'Senator'),
        ('citizen', 'Citizen'),
        ('other', 'Other'),
    ]

    user = models.OneToOneField(User, on_delete=models.CASCADE)
    full_name = models.CharField(max_length=100)
    id_number = models.CharField(max_length=20)
    category = models.CharField(max_length=10, choices=CATEGORY_CHOICES)

    def __str__(self):
        return self.full_name
