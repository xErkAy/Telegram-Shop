<template>
  <v-app>
    <notifications position="top right" />
    <div class="container">
      <div class="pb-3 mb-4 mt-2 border-bottom">
        <router-link to="/" class="d-flex align-items-center text-dark text-decoration-none">
          <span class="fs-4">Просмотр заказов</span>
        </router-link>
      </div>
    </div>
    <v-main>
      <router-view/>
    </v-main>
    <div class="container">
        <footer>
          <ul class="nav justify-content-center border-bottom pb-3 mb-3"></ul>
          <p class="text-center text-muted">© 2022</p>
        </footer>
    </div>
  </v-app>
</template>

<script>
export default {
  name: 'App',
  data: () => ({
    
  }),
};
</script>
