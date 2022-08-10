from django.urls import path
from . import views


urlpatterns = [
    path('users/', views.GetUsers.as_view()),
    path('messages/<int:user_id>', views.GetMessages.as_view()),
    path('orders/', views.GetAllOrders.as_view()),
    path('orders/<int:order_id>', views.GetSpecificOrder.as_view()),
]