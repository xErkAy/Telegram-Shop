from rest_framework import serializers
from shop.models import *


class UsersSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = "__all__"


class OrderSerializer(serializers.ModelSerializer):
    user = serializers.JSONField(source='get_user')

    class Meta:
        model = Order
        exclude = ("user_id", "is_chat_active")


class MessagesSerializer(serializers.ModelSerializer):
    class Meta:
        model = Message
        exclude = ("message_id", )


class CreateOrderSerializer(serializers.Serializer):
    user_id = serializers.IntegerField()
    order_value = serializers.CharField()


class UpdateOrderSerializer(serializers.Serializer):
    user_id = serializers.IntegerField()
    status = serializers.IntegerField()
    is_closed = serializers.BooleanField(required=False)