<template>
  <div class="container">
    <div class="row pt-lg-1 rounded-3 border shadow-lg">
      <div class="p-3 p-lg-3 pt-lg-3">
        <div class="d-flex justify-space-between align-center">
          <h3 class="fw-bold">Заказ №{{ item.order_id }}</h3>
          <v-btn
            @click="changeOrderStatus"
          >Сохранить</v-btn>
        </div>
        <div><span class="fw-bold">Имя:</span> {{ item.user ? item.user.first_name : ''}}</div>
        <div><span class="fw-bold">Заказано</span> {{ convertDate(item.date) }}</div>
        <div class="d-flex mt-1">
          <div class="fw-bold mt-2 mr-2">Статус заказа:</div>
          <div style="width: 27%">
            <v-select
              v-on:change='this.changeOrderColor'
              v-model="selectedStatus"
              :items="statusNames"
              item-text="name"
              item-value="value"
              :background-color="order_color"
              label="Изменить статус"
              solo
              dense
            ></v-select>
          </div>
        </div>
        <div><span class="fw-bold">Содержимое заказа:</span></div>
        {{ item.order_value }}
    </div>
  </div>
  </div>
</template>

<script>
import {
  getSpecificOrder,
  changeOrderStatus,
} from '@/api/index'

export default {
  name: 'SpecificOrder',
  data() {
    return {
      item: {},
      statusNames: [
        { name: 'готовится', value: '2', color: 'orange' },
        { name: 'готов к выдаче', value: '3', color: 'green' },
        { name: 'выдан', value: '4', color: 'red' },
      ],
      selectedStatus: '',
      order_color: '',
    }
  },
  mounted() {
    this.getSpecificOrderList()
  },
  methods: {
    async getSpecificOrderList() {
      const response = await getSpecificOrder(this.$route.params.id)
      this.item = response.data
      this.selectedStatus = this.getOrderStatus(this.item.status, this.item.is_closed)
      this.order_color = this.statusNames.find(obj => this.selectedStatus == obj.value).color
    },
    convertDate(date) {
      date = new Date(date).toLocaleString().replaceAll('/', '.').split(", ")
      return `${date[0]} в ${date[1]}`
    },
    getOrderStatus(status, is_closed) {
      if (is_closed) {
        return '4'
      }
      const findObj = this.statusNames.find(obj => status == obj.value)
      return findObj ? findObj.value : ''
    },
    changeOrderColor() {
      this.order_color = this.statusNames.find(obj => this.selectedStatus == obj.value).color
    },
    changeOrderStatus() {
      let is_closed = false
      let status = this.item.status
      if (this.selectedStatus === '4') {
        is_closed = true
      } else {
        status = parseInt(this.selectedStatus)
      }
      const payload = {
        order_id: this.item.order_id,
        data: {
          user_id: this.item.user.user_id,
          status: status,
          is_closed: is_closed
        }
      }
      changeOrderStatus(payload)
        .then(
          (res) => {
            this.$notify({
              type: res.data.type,
              text: res.data.message
            })
          },
          (err) => {
            this.$notify({
              type: err.response.data.type,
              text: err.response.data.message
            })
          }
        )
    },
  }
}
</script>