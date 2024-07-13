# accounts/forms.py

from django import forms
from django.contrib.auth.models import User
from django.contrib.auth.forms import UserCreationForm

class UserRegisterForm(UserCreationForm):
    full_name = forms.CharField(max_length=100)
    id_number = forms.CharField(max_length=20)
    CATEGORY_CHOICES = [
        ('president', 'President'),
        ('governor', 'Governor'),
        ('mp', 'MP'),
        ('senator', 'Senator'),
        ('citizen', 'Citizen'),
        ('other', 'Other'),
    ]
    category = forms.ChoiceField(choices=CATEGORY_CHOICES)

    class Meta:
        model = User
        fields = ['username', 'full_name', 'id_number', 'category', 'password1', 'password2']
