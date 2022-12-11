import socket
import json

from django.db import transaction
from rest_framework.generics import ListAPIView, RetrieveUpdateAPIView
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status

from django.core.exceptions import ObjectDoesNotExist

from project.exceptions import ValidationExceptionWithMessage
from project.constants import ORDER_STATUS
from shop.serializers import (
    MessagesSerializer,
    UsersSerializer,
    OrderSerializer,
    UpdateOrderSerializer,
    CreateOrderSerializer
)
from shop.models import Message, User, Order


class GetUsers(ListAPIView):
    queryset = User.objects.all()
    serializer_class = UsersSerializer


class GetAllOrders(ListAPIView):
    serializer_class = OrderSerializer

    def get_queryset(self):
        return Order.objects.filter(is_closed=self.request.query_params.get('closed', False) == 'true').order_by("-date")

    def post(self, request, *args, **kwargs):
        data = CreateOrderSerializer(data=request.data)
        data.is_valid(raise_exception=True)
        obj = Order.objects.create(user_id=User(user_id=data.validated_data.get('user_id')),
                                                order_value=data.validated_data.get('order_value'), status=1)
        return Response(data={"message": f"Заказ №{obj.order_id} успешно создан.", "type": "success"}, status=status.HTTP_200_OK)


class GetSpecificOrder(RetrieveUpdateAPIView):
    serializer_class = OrderSerializer
    queryset = Order.objects.all()

    def patch(self, request, *args, **kwargs):
        data = UpdateOrderSerializer(data=request.data)
        data.is_valid(raise_exception=True)

        user_id = data.validated_data.get("user_id")
        order_status = data.validated_data.get("status")
        is_closed = data.validated_data.get("is_closed", None)
        order_id = kwargs.get('pk')

        with transaction.atomic():
            try:
                obj = Order.objects.get(user_id=user_id, order_id=order_id)
            except ObjectDoesNotExist:
                raise ValidationExceptionWithMessage('Ошибка изменения статуса')

            obj.status = order_status
            if is_closed is not None:
                obj.is_closed = is_closed
            obj.save()

            try:
                if is_closed:
                    order_status = 4
                message = {
                    'user_id': user_id,
                    'message_text': f'[Уведомление]\nСтатус заказа №{order_id} изменен на: {ORDER_STATUS[order_status - 1]}'
                }
                SendMessage(message)
                return Response(data={"message": "Статус заказа изменен", "type": "success"}, status=status.HTTP_200_OK)
            except:
                raise ValidationExceptionWithMessage('Ошибка уведомления пользователя')


class GetMessages(ListAPIView):
    queryset = Message.objects.all()
    serializer_class = MessagesSerializer

    def get_queryset(self):
        return self.queryset.filter(user_id=self.kwargs["user_id"]).order_by("-date")[0:100]


class CreateMessages(APIView):
    def post(self, request, *args, **kwargs):
        try:
            data = request.data
            user_id = data.get("user_id")
            order_id = data.get("order_id")
            message_text = data.get("message_text")
            is_sender = data.get("is_sender")
            activity = Order.objects.get(order_id=order_id, user_id__id=user_id)
            if activity.get_chat_activity:
                Message.objects.create(user_id=User(user_id=user_id), order_id=Order(order_id=order_id), message_text=message_text, is_sender=is_sender)
                if not is_sender:
                    try:
                        message = {
                            'user_id': int(user_id),
                            'message_text': f'[Оператор к заказу №{order_id}]\n{message_text}'
                        }
                        SendMessage(message)
                    except:
                        return Response(data={"message": "Ошибка уведомления пользователя", "type": "error"}, status=status.HTTP_400_BAD_REQUEST)
                return Response(status=status.HTTP_200_OK)
            else:
                return Response(data={"message": "Чат неактивен", "type": "error"}, status=status.HTTP_400_BAD_REQUEST)
        except:
            return Response(status=status.HTTP_400_BAD_REQUEST)


def SendMessage(message):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as client:
        client.connect(("192.168.1.63", 8001))
        client.send(json.dumps(message).encode("utf-8"))
