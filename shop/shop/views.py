import socket
import json
from rest_framework.generics import ListAPIView, UpdateAPIView
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from django.core.exceptions import ObjectDoesNotExist
from .serializers import MessagesSerializer, UsersSerializer, OrderSerializer
from .models import Messages, Users, Orders


class GetUsers(ListAPIView):
    queryset = Users.objects.all()
    serializer_class = UsersSerializer


class GetAllOrders(ListAPIView):
    serializer_class = OrderSerializer

    def get_queryset(self):
        if self.request.query_params.get('closed', None) == 'true':
            return Orders.objects.filter(is_closed=True).order_by("-date")
        return Orders.objects.filter(is_closed=False).order_by("-date")

    def post(self, request, *args, **kwargs):
        data = request.data
        try:
            user_id = data.get("user_id")
            order_value = data.get("order_value")
            Orders.objects.create(user_id=Users(user_id=user_id), order_value=order_value, status=1)
            return Response(data={"message": f"Заказ №{Orders.objects.latest('order_id').order_id} успешно создан."}, status=status.HTTP_200_OK)
        except:
            return Response(data={"message": "Произошла ошибка при создании заказа."}, status=status.HTTP_400_BAD_REQUEST)


class GetSpecificOrder(ListAPIView):
    serializer_class = OrderSerializer

    def get_queryset(self):
        return Orders.objects.filter(order_id=self.kwargs['order_id'])


class ChangeOrderStatus(APIView):
    def post(self, request, *args, **kwargs):
        data = request.data
        try:
            user_id = data.get("user_id")
            order_id = data.get("order_id")
            order_status = data.get("status")
            is_closed = data.get("is_closed", None)

            obj = Orders.objects.get(user_id=user_id, order_id=order_id)
            obj.status = order_status
            if is_closed is not None:
                obj.is_closed = is_closed
            obj.save()

            try:
                if is_closed:
                    order_status = 4
                message = {
                    'user_id': user_id,
                    'message_text': f'[Уведомление]\nСтатус заказа №{order_id} изменен на: {get_order_status[order_status - 1]}'
                }
                SendMessage(message)
                return Response(data={"message": "Статус заказа изменен", "type": "success"}, status=status.HTTP_200_OK)
            except:
                return Response(data={"message": "Ошибка уведомления пользователя", "type": "error"}, status=status.HTTP_400_BAD_REQUEST)
        except ObjectDoesNotExist:
            return Response({"message": "Ошибка изменения статуса", "type": "error"}, status=status.HTTP_400_BAD_REQUEST)


class GetMessages(ListAPIView):
    queryset = Messages.objects.all()
    serializer_class = MessagesSerializer

    def get(self, request, *args, **kwargs):
        messages = self.queryset.filter(user_id=kwargs["user_id"]).order_by("-date")[0:100]
        return Response(self.serializer_class(messages, many=True).data)

class CreateMessages(APIView):
    def post(self, request, *args, **kwargs):
        try:
            data = request.data
            user_id = data.get("user_id")
            order_id = data.get("order_id")
            message_text = data.get("message_text")
            is_sender = False
            activity = Orders.objects.get(order_id=order_id, user_id=Users(user_id=user_id))
            if activity.get_chat_activity:
                Messages.objects.create(user_id=Users(user_id=user_id), order_id=Orders(order_id=order_id), message_text=message_text, is_sender=is_sender)
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

get_order_status = ['', 'готовится', 'готов к выдаче', 'выдан']


def SendMessage(message):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as client:
        client.connect(("192.168.88.57", 8001))
        client.send(json.dumps(message).encode("utf-8"))
