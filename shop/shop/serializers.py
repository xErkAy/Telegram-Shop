from rest_framework import serializers
from .models import *


class UsersSerializer(serializers.ModelSerializer):
    class Meta:
        model = Users
        fields = "__all__"

class OrderSerializer(serializers.ModelSerializer):
    user = serializers.JSONField(source='get_user')

    class Meta:
        model = Orders
        exclude = ("user_id", )

class MessagesSerializer(serializers.ModelSerializer):
    class Meta:
        model = Messages
        exclude = ("message_id", )