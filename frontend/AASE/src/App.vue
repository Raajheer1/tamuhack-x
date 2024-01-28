<template>
  <div id="app">
    <div v-if="page === 'welcome'">
      <Welcome @next="toLogin" />
    </div>
    <div v-else-if="page === 'login'">
      <Login @back="toWelcome" @home="toHome" />
    </div>
    <div v-else-if="page === 'home'">
      <Home />
    </div>
    <div v-else-if="page === 'account'">
      <Account />
    </div>
    <div v-else-if="page === 'seat'">
      Seat Page
    </div>
    <div v-else-if="page === 'confirmation'">
      <Confirmation @home="toHome" @account="toAccount"/> <!--@account="toAccount"-->
    </div>
  </div>
</template>


<script setup lang="ts">
import { ref } from 'vue';
import Welcome from "./components/Welcome.vue";
import Login from "./components/Login.vue";
import Home from "./components/Home.vue";
import Confirmation from './components/Confirmation.vue';
import Account from './components/Account.vue';

// welcome, login, home, account, seat, confirmation
const page = ref<string>("login");

  const toLogin = () => {
    page.value = "login";
  }

  const toWelcome = () => {
    page.value = "welcome";
  }

  const toHome = () => {
    page.value = "home";
  }

  const toAccount = () => {
    page.value = "account"
  }


</script>