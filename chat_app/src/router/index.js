import Vue from 'vue'
import VueRouter from 'vue-router'
import AllOrders from '../views/AllOrders.vue'
import AllOrdersClosed from '../views/AllOrdersClosed.vue'
import SpecificOrder from '../views/SpecificOrder.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'AllOrders',
    component: AllOrders
  },
  {
    path: '/closed/',
    name: 'AllOrdersClosed',
    component: AllOrdersClosed
  },
  {
    path: "/:id/",
    name: "SpecificOrder",
    component: SpecificOrder,
  },
]

const router = new VueRouter({
  routes
})

export default router
