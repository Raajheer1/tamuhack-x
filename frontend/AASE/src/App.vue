<template>
  <transition name="slide-fade" mode="out-in">
    <div id="app" :key="page">
        <div v-if="page === 'welcome'" :key="page">
          <Welcome @next="toLogin" />
        </div>
        <div v-if="page === 'login'" :key="page">
          <Login @back="toWelcome" @home="toHome" />
        </div>
        <div v-if="page === 'home'" :key="page">
          <Home @seat="toSeat"  />
        </div>
        <div v-if="page === 'account'" :key="page">
          <Account />
        </div>
        <div v-if="page === 'seats'" :key="page">
          <Seats />
        </div>
        <div v-if="page === 'confirmation'" :key="page">
          <Confirmation @home="toHome" @account="toAccount"/>
        </div>
    </div>
  </transition>
</template>


<script setup lang="ts">
  import { ref } from 'vue';
  import Welcome from "./components/Welcome.vue";
  import Login from "./components/Login.vue";
  import Home from "./components/Home.vue";
  import Confirmation from './components/Confirmation.vue';
  import Account from "./components/Account.vue";
  import Seats from "./components/Seats.vue";

// welcome, login, home, account, seat, confirmation
// const page = ref<string>("seats");
// welcome, login, home, account, seat, confirmation
const page = ref<string>("account");



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

const toSeat = (id: string) => {
  page.value = "seats"
}


</script>

<style scoped>
.slide-fade-enter-active {
  transition: all 0.1s ease-out;
}

.slide-fade-leave-active {
  transition: all 0.3s cubic-bezier(1, 0.5, 0.8, 1);
}

.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateX(20px);
  opacity: 0;
}
</style>