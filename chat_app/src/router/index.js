import Vue from 'vue'
import VueRouter from 'vue-router'
import AllOrders from '../views/AllOrders.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'AllOrders',
    component: AllOrders
  },
]

const router = new VueRouter({
  routes
})

export default router
