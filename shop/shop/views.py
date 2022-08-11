from rest_framework.generics import ListAPIView, UpdateAPIView
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
        if self.request.query_params.get('all', None) == 'false':
            return Orders.objects.filter(is_closed=False).order_by("-date")
        return Orders.objects.all().order_by("-date")

class GetSpecificOrder(ListAPIView):
    serializer_class = OrderSerializer

    def get_queryset(self):
        return Orders.objects.filter(order_id=self.kwargs['order_id'])

class ChangeOrderStatus(UpdateAPIView):
    def put(self, request, *args, **kwargs):
        data = request.data
        try:
            Orders.objects.filter(user_id=data.get("user_id"), order_id=data.get("order_id")).update(status=data.get("status"))
            return Response(status=status.HTTP_200_OK)
        except ObjectDoesNotExist:
            return Response(status=status.HTTP_404_NOT_FOUND)

class GetMessages(ListAPIView):
    queryset = Messages.objects.all()

    def get(self, request, *args, **kwargs):
        messages = self.queryset.filter(user_id=kwargs["user_id"]).order_by("-date")[0:100]
        return Response(MessagesSerializer(messages, many=True).data)
