<template>
  <div class="container">
    <v-checkbox
      v-model="allStatus"
      label="Все заказы"
      @click="getAllOrdersList(allStatus.toString())"
    ></v-checkbox>
    <div v-for="item in data" class="mb-2">
        <div class="row pt-lg-1 rounded-3 border shadow-lg">
          <div class="p-3 p-lg-3 pt-lg-3">
            <h3 class="fw-bold">Заказ №{{ item.order_id }}</h3>
            <div><span class="fw-bold">Имя:</span> {{ item.user.first_name }}</div>
            <div><span class="fw-bold">Статус:</span> {{ getOrderStatus(item.status, item.is_closed) }}</div>
            <div><span class="fw-bold">Заказано</span> {{ convertDate(item.date) }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {
  getAllOrders,
  getSpecificOrder,
} from '@/api/index'

export default {
  name: 'AllOrders',
  data() {
    return {
      data: "",
      items: [],
      statusNames: [
        { name: 'не взят в работу', value: '1' },
        { name: 'готовится', value: '2' },
        { name: 'готов к выдаче', value: '3' },
      ],
      allStatus: false,
    }
  },
  mounted() {
    this.getAllOrdersList()
  },
  methods: {
    async getAllOrdersList() {
      const response = await getAllOrders(`?all=${this.allStatus}`)
      this.data = response.data
    },
    convertDate(date) {
      date = new Date(date).toLocaleString().replaceAll('/', '.').split(", ")
      return `${date[0]} в ${date[1]}`
    },
    getOrderStatus(status, is_closed) {
      if (is_closed) {
        return 'закрыт'
      }
      return this.statusNames.find(obj => status == obj.value).name
    }
  }
}
</script>

<style scoped>
.box {
  font-size: 18px;
  border-style: solid;
  border-radius: 2%;
  margin-bottom: 1%;
  padding: 1%;
}
</style>