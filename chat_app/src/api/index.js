import axios from 'axios'

export const instance = axios.create({
    baseURL: 'http://192.168.88.57:8000/api/',
})

export const getAllOrders = params => instance.get('orders/', { params })
export const getSpecificOrder = id => instance.get(`orders/${id}/`)
export const changeOrderStatus = payload => instance.post(`orders/changestatus/`, payload)

export const NewMessage = payload => instance.post('message/', payload)