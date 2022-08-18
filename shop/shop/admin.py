from django.contrib import admin
from .models import *

admin.site.register(Users)
admin.site.register(Messages)
admin.site.register(Orders)
admin.site.register(MessagesActivity)