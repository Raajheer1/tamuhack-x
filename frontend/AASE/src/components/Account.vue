<template>
    <div class="flex flex-col h-screen">
        <!-- top information tag -->
        <div class="w-full h-32 bg-aa-blue rounded-b-3xl">
            <div class="flex justify-between">
                <div class="text-white pl-3 py-5">
                    <h1 class="text-2xl">{{FirstName}}John {{LastName}}Doe</h1>
                    <h1 class = "text-l">AAdvantage {{aaAdvantageNum}}#8QUJ879</h1>
                    <h1 class="text-l">Member Since {{joinMonth}}/{{joinYear}}</h1>
                </div>
                <img src="../assets/aa-tail.svg" alt="American Airlines Logo" class="h-20 w-20 rounded-full m-5">
            </div>
        </div>
        <!-- start balance area -->
        <div class="w-stretch h-48 bg-slate-200 rounded-3xl m-3 p-2 px-3">
            <h2>Current Balance</h2>
            <h1 class="text-4xl">{{ Money }}$15,123.98</h1>
            <div class="flex justify-between">
                <button class="rounded-full h-16 w-72 bg-white my-7 mx-3 text-xl">Withdraw</button>
                <button class="rounded-full h-16 w-72 bg-black my-7 mx-3 text-xl text-white">Add Funds</button>
            </div>
        </div>
        <!-- starting the scrolling -->
        <div class="h-stretch overflow-auto bg-slate-200 rounded-3xl py-10 px-6 inner w-stretch mx-3 mt-1 mb-3">
            <div class="flex flex-col" v-for="seat in seats">
                <div>
                    <SeatCard :seat="seat" />
                </div>
            </div>
        </div>
    </div>
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

    const FirstName = ref<string>("")

    const LastName = ref<string>("")

    const aaAdvantageNum = ref<string>("")

    const joinDate = randomDate(new Date(1981, 5), new Date())
    const joinMonth = joinDate.getMonth()
    const joinYear = joinDate.getFullYear()

    function randomDate(start, end) {
        return new Date(start.getTime() + Math.random() * (end.getTime() - start.getTime()));
    }

    const Money = ref()

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