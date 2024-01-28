<template>
  <Transition name="nested">
    <div class="flex flex-col pt-10 bg-aa-blue h-screen outer">
      <img class="absolute" src="../assets/globe.svg" alt="World" />
      <div class="mx-10 mt-16 mb-14">
        <div class="flex justify-between">
          <p class="mt-0.5 text-xl text-white tracking-wider">Upcoming Flights</p>
          <i class="h-5 w-5 p-2 rounded-full my-auto fa-solid fa-user text-white border border-white"></i>
        </div>
        <h1 class="text-4xl text-white font-semibold mt-8 tracking-wider">Your next adventure awaits</h1>
      </div>
      <div class="max-h-screen overflow-auto bg-white rounded-t-3xl py-10 px-6 inner">
        <div class="flex flex-col" v-for="seat in seats">
          <div :key="seat.id" @click="clicked(seat.id.toString())">
            <SeatCard :seat="seat" />
          </div>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import {Seat} from "../types";
import { ref } from 'vue';
import SeatCard from "./SeatCard.vue";

const emit = defineEmits<{
  (e: 'seat', id: string): void
}>()

const clicked = (id: string) => {
  emit('seat', id)
}

const seats = ref<Seat[]>([
  {
    id: 1,
    seat_number: "1A",
    flight: {
      id: 1,
      origin: "SEA",
      destination: "ORD",
      aircraft_type: "737",
      schedule_departure_time: "2021-10-10T12:00:00.000Z",
      schedule_arrival_time: "2021-10-10T18:00:00.000Z",
    },
    assigned_pnr: "ABCDEF",
    bid_time_end: new Date().toString(),
    buy_now_price: 0,
    current_bid: 0,
    flight_id: 0,
    for_sale: false,
    last_bidder: "",
    min_price: 0,
  },
  {
    id: 2,
    seat_number: "3B",
    flight: {
      id: 66,
      origin: "DFW",
      destination: "NYC",
      aircraft_type: "737",
      schedule_departure_time: "2021-10-10T12:00:00.000Z",
      schedule_arrival_time: "2021-10-10T18:00:00.000Z",
    },
    assigned_pnr: "DEFGFF",
    bid_time_end: new Date().toString(),
    buy_now_price: 0,
    current_bid: 0,
    flight_id: 0,
    for_sale: false,
    last_bidder: "",
    min_price: 0,
  },
  {
    id: 3,
    seat_number: "11F",
    flight: {
      id: 46,
      origin: "IAH",
      destination: "CLL",
      aircraft_type: "737",
      schedule_departure_time: "2021-10-10T12:00:00.000Z",
      schedule_arrival_time: "2021-10-10T18:00:00.000Z",
    },
    assigned_pnr: "DEFGFF",
    bid_time_end: new Date().toString(),
    buy_now_price: 0,
    current_bid: 0,
    flight_id: 0,
    for_sale: false,
    last_bidder: "",
    min_price: 0,
  },
  {
    id: 4,
    seat_number: "21F",
    flight: {
      id: 46,
      origin: "MIA",
      destination: "SFO",
      aircraft_type: "737",
      schedule_departure_time: "2021-10-10T12:00:00.000Z",
      schedule_arrival_time: "2021-10-10T18:00:00.000Z",
    },
    assigned_pnr: "DEFGFF",
    bid_time_end: new Date().toString(),
    buy_now_price: 0,
    current_bid: 0,
    flight_id: 0,
    for_sale: false,
    last_bidder: "",
    min_price: 0,
  }
]);


</script>

<style scoped>
.nested-enter-active .inner,
.nested-leave-active .inner {
  transition: all 0.3s ease-in-out;
}

.nested-enter-from .inner,
.nested-leave-to .inner {
  transform: translateX(30px);
  opacity: 0;
}

.nested-enter-active .inner {
  transition-delay: 1s;
}
</style>