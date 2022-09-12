<template>
  <div class="container">
    <v-btn
      class="mb-7 mt-4"
      @click="getAllOrdersClosed()"
    >Завершенные заказы</v-btn>
    <div v-for="item in items" class="mb-2">
        <div class="row pt-lg-1 rounded-3 border shadow-lg">
          <div class="p-3 p-lg-3 pt-lg-3">
            <h3 class="fw-bold">Заказ №{{ item.order_id }}</h3>
            <div><span class="fw-bold">Имя:</span> {{ item.user.first_name }}</div>
            <div><span class="fw-bold">Статус:</span> <span :style="`color: ${getOrderColorByStatus(item.status)}`">{{ getOrderStatus(item.status) }}</span></div>
            <div><span class="fw-bold">Заказано</span> {{ convertDate(item.date) }}</div>
            <div style="float: right;">
              <v-btn
                @click="acceptOrder(item)"
                v-if="item.status === 1"
                color="red"
              >Принять заказ</v-btn>
              <v-btn
                v-else
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
  changeOrderStatus,
} from '@/api/index'

export default {
  name: 'AllOrders',
  data() {
    return {
      items: [],
      updatedTime: 5000,
      intervalId: undefined,
      statusNames: [
        { name: 'не взят в работу', value: '1', color: 'black' },
        { name: 'готовится', value: '2', color: 'orange' },
        { name: 'готов к выдаче', value: '3', color: 'green' },
      ],
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
      const response = await getAllOrders()
      this.items = response.data
    },
    convertDate(date) {
      date = new Date(date).toLocaleString().replaceAll('/', '.').split(', ')
      return `${date[0]} в ${date[1]}`
    },
    getOrderStatus(status) {
      return this.statusNames.find(obj => status == obj.value).name
    },
    getOrderColorByStatus(status) {
      return this.statusNames.find(obj => status == obj.value).color
    },
    goToSpecificOrder(id) {
      this.$router.push({path: '/' + id})
    },
    getAllOrdersClosed() {
      this.$router.push({path: '/closed'})
    },
    acceptOrder(item) {
      const payload = {
        user_id: item.user.user_id,
        order_id: item.order_id,
        status: 2,
      }
      changeOrderStatus(payload)
        .then(
          (res) => {
            this.$notify({
              type: 'success',
              text: 'Заказ принят'
            })
            const findObj = this.items.find(i => i.order_id === item.order_id)
            findObj ? findObj.status = payload.status : ''
          },
          (err) => {
            this.$notify({
              type: 'error',
              text: 'Произошла ошибка'
            })
          }
        )
    },
    updateOrders() {
      this.intervalId = setInterval(() => {
        this.getAllOrdersList()
      }, this.updatedTime)
    }
  }
}
</script>