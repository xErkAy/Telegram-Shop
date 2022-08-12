<template>
  <div class="container">
    <v-btn
      class="mb-7 mt-4"
      @click="getAllOrdersOpened()"
    >Открытые заказы</v-btn>
    <div v-for="item in items" class="mb-2">
        <div class="row pt-lg-1 rounded-3 border shadow-lg">
          <div class="p-3 p-lg-3 pt-lg-3">
            <h3 class="fw-bold">Заказ №{{ item.order_id }}</h3>
            <div><span class="fw-bold">Имя:</span> {{ item.user.first_name }}</div>
            <div><span class="fw-bold">Статус:</span> <span style="color: red">завершен</span></div>
            <div class="d-flex justify-space-between align-center">
              <div><span class="fw-bold">Заказано</span> {{ convertDate(item.date) }}</div>
              <v-btn
                @click="goToSpecificOrder(item.order_id)"
              >Подробнее</v-btn>
            </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {
  getAllOrders,
} from '@/api/index'

export default {
  name: 'AllOrdersClosed',
  data() {
    return {
      updatedTime: 5000,
      intervalId: undefined,
      items: [],
    }
  },
  mounted() {
    this.getAllOrdersList()
    this.updateOrders()
  },
  destroyed() {
    this.intervalId ? clearInterval(this.intervalId) : '';
  }, 
  methods: {
    async getAllOrdersList() {
      const response = await getAllOrders({ closed:true })
      this.items = response.data
    },
    convertDate(date) {
      date = new Date(date).toLocaleString().replaceAll('/', '.').split(', ')
      return `${date[0]} в ${date[1]}`
    },
    goToSpecificOrder(id) {
      this.$router.push({path: '/' + id})
    },
    getAllOrdersOpened() {
      this.$router.push({path: '/'})
    },
    updateOrders() {
      this.intervalId = setInterval(() => {
        this.getAllOrdersList()
      }, this.updatedTime)
    }
  }
}
</script>