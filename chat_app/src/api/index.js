import axios from 'axios'

export const instance = axios.create({
    baseURL: 'http://localhost:8000/api/',
})

export const getAllOrders = params => instance.get('orders/', { params })
export const getSpecificOrder = id => instance.get(`orders/${id}/`)
export const changeOrderStatus = payload => instance.patch(`orders/${payload.order_id}/`, payload.data)

export const NewMessage = payload => instance.post('message/', payload)