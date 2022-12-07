from django.urls import path
from shop.views import (
    GetUsers,
    CreateMessages,
    GetMessages,
    GetAllOrders,
    GetSpecificOrder
)

urlpatterns = [
    path('users/', GetUsers.as_view()),
    path('messages/', CreateMessages.as_view()),
    path('messages/<int:user_id>/', GetMessages.as_view()),
    path('orders/', GetAllOrders.as_view()),
    path('orders/<int:pk>/', GetSpecificOrder.as_view()),
]