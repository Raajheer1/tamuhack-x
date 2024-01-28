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
          <Home @seat="toSeat" @account="toAccount" />
        </div>
        <div v-if="page === 'account'" :key="page">
          <Account @home="toHome" @home2="toHome2" @login="toLogin" :seats="seats"/>
        </div>
        <div v-if="page === 'seats'" :key="page">
          <Seats :selectedSeat="selectedSeat" :emptySeats="emptySeats" />
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
  import {API} from "./utils/api.ts";

  import {Seat} from "./types";

// welcome, login, home, account, seat, confirmation
const page = ref<string>("welcome");

const AAAdvantageID = ref<string>("");

const selectedSeat = ref<Seat>();
const emptySeats = ref<Seat[]>([]);

const seats = ref<Seat[]>([]);

const toLogin = () => {
  page.value = "login";
}

const toWelcome = () => {
  page.value = "welcome";
}

const toHome = (AAdvantageID: string) => {
  AAAdvantageID.value = AAdvantageID;
  fetchSeats();
  page.value = "home";
}

const toHome2 = () => {
  page.value = "home";
}

const toAccount = () => {
  page.value = "account"
}

const toSeat = (SeatID: string) => {
  seats.value.forEach((seat) => {
    if (seat.id.toString() === SeatID) {
      selectedSeat.value = seat
    }
  })
  page.value = "seats"
}

// Fetch the data from the backend
const fetchSeats = async () => {
  const response = await API.get("/account/fetchall/" + AAAdvantageID.value);
  seats.value = response.data;

  // TODO - Fetch the empty seats on this plane
  const secondary = await API.get("/seat/flight/" + selectedSeat.value?.flight_id);
  secondary.data.forEach((seat: Seat) => {
    if(seat.for_sale) {
      emptySeats.value.push(seat);
    }
  })
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