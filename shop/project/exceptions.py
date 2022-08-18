from rest_framework.exceptions import APIException


class ExceptionSendMessage(BaseException):
    status_code = 503
    default_detail = 'Service temporarily unavailable, try again later.'
    default_code = 'service_unavailable'


class AppAPIException(APIException):

    def __init__(self, detail=None, status=None):
        self.status_code = status if status is not None else self.status_code
        self.default_detail = detail if detail is not None else self.default_detail


    status_code = 400
    default_detail = 'Service temporarily unavailable, try again later.'
    default_code = 'service_unavailable'