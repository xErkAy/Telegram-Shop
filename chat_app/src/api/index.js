import axios from 'axios'

export const instance = axios.create({
    baseURL: 'http://localhost:8000/api/',
})

export const getAllOrders = payload => instance.get(`orders/${payload}`)
export const getSpecificOrder = id => instance.get(`orders/${id}`)

export const NewMessage = payload => instance.post('message/', payload)