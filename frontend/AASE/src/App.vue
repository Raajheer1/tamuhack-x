<template>
  <div id="app">
    <div v-if="page === 'welcome'">
      <Welcome @next="toLogin" />
    </div>
    <div v-else-if="page === 'login'">
      Login Page
    </div>
    <div v-else-if="page === 'home'">
      Home Page
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

// welcome, login, home, account, seat, confirmation
const page = ref<string>("welcome");

const toLogin = () => {
  page.value = "login";
}
</script>
