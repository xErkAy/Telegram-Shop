from rest_framework.exceptions import APIException


class ValidationExceptionWithMessage(APIException):
    status_code = 400
    default_code = '1300'
    default_detail = 'error message'