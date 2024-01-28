<template>
  <div id="app">
    <div v-if="page === 'welcome'">
      <Welcome @next="toLogin" />
    </div>
    <div v-else-if="page === 'login'">
      <Login @back="toWelcome" />
    </div>
    <div v-else-if="page === 'home'">
      <Home />
    </div>
    <div v-else-if="page === 'account'">
      Account Page
    </div>
    <div v-else-if="page === 'seat'">
      Seat Page
    </div>
    <div v-else-if="page === 'confirmation'">
      Confirmation Page
    </div>
  </div>
</template>


<script setup lang="ts">
import { ref } from 'vue';
import Welcome from "./components/Welcome.vue";
import Login from "./components/Login.vue";
import Home from "./components/Home.vue";

// welcome, login, home, account, seat, confirmation
const page = ref<string>("home");

const toLogin = () => {
  page.value = "login";
}

const toWelcome = () => {
  page.value = "welcome";
}
</script>