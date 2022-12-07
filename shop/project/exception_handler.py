from django.http import Http404
from rest_framework.exceptions import APIException, ValidationError
from rest_framework.views import exception_handler
from rest_framework import status


def custom_exception_handler(exc, context):
    response = exception_handler(exc, context)
    if response is not None:
        if isinstance(exc, APIException):
            if isinstance(exc, ValidationError):
                response.data = {'message': response.data}
                exc.default_code = '2010'
            else:
                response.data['message'] = exc.detail
                del response.data['detail']
            response.data = {
                'success': False,
                'code': exc.default_code,
                'message': response.data.get('message')
            }
            response.status_code = exc.status_code
        elif isinstance(exc, Http404):
            response.data['message'] = response.reason_phrase
            response.data.update({'success': False, 'code': response.status_code})
            response.status_code = status.HTTP_404_NOT_FOUND
        return response
