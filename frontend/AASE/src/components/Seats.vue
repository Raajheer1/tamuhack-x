<template>
  <div class="flex flex-col bg-aa-blue h-screen relative">
    <img class="absolute mx-10 flex items-center z-5" src="../assets/globe.svg" alt="World" />

      <!-- Flight Details -->
    <div class="w-120 h-32 bg-white rounded-b-3xl shadow-md p-4 relative z-10">
        <div class="flex flex-row justify-between h-full">
            <!-- Departure Details -->
            <div class="flex flex-row">
                <div class="flex flex-col">
                    <div class="text-sm text-black flex flex-row">
                        <!-- Departure time -->
                      {{ new Date(mySeat.flight.scheduled_departure_time).toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'}) }}
                    </div>
                    <div class="text-sm text-opacity-70 text-gray-700 flex flex-row">
                        <!-- Departure date -->
                      {{ new Date(mySeat.flight.scheduled_departure_time).toLocaleDateString() }}
                    </div>
                </div>
                <div class="flex flex-col mx-4"> <!-- Add margin to create a symmetrical gap between time and city details -->
                    <div class="text-sm text-black flex flex-row">
                        <!-- Departure Airport Code -->
                      {{ mySeat.flight.origin }}
                    </div>
                    <div class="text-sm text-opacity-70 text-gray-700 flex flex-row">
                        <!-- Departure city -->
                        {{ getCityFromAirportCode(mySeat.flight.origin) }}
                    </div>
                </div>
            </div>

            <!-- Centered Plane Icon and American Airlines Logo -->
            <div class="flex items-center justify-center flex flex-col">
                <img src="../assets/aa-tail.svg" alt="American Airlines Logo" class="h-20 w-20 rounded-full">
                <div class="text-sm text-black flex flex-row font-weight-bold">
                    Seat: {{ mySeat.seat_number }}
                </div>
            </div>

            <!-- Arrival Details -->
            <div class="flex flex-row">
                <div class="flex flex-col mx-4"> <!-- Add margin to create a symmetrical gap between time and city details -->
                    <div class="text-sm text-black flex flex-row">
                        <!-- Arrival Airport Code -->
                      {{ mySeat.flight.destination }}
                    </div>
                    <div class="text-sm text-opacity-70 text-gray-700 flex flex-row">
                        <!-- Arrival City with 70% opacity -->
                        {{ getCityFromAirportCode(mySeat.flight.destination) }}
                    </div>
                </div>
                <div class="flex flex-col">
                    <div class="text-sm text-black flex flex-row">
                        <!-- Arrival time -->
                      {{ new Date(mySeat.flight.scheduled_arrival_time).toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'}) }}
                    </div>
                    <div class="text-sm text-opacity-70 text-gray-700 flex flex-row">
                        <!-- Arrival date with 70% opacity -->
                      {{ new Date(mySeat.flight.scheduled_departure_time).toLocaleDateString() }}
                    </div>
                </div>
            </div>
        </div>
    </div>

      <!-- seatmap -->
    <div class="flex flex-col p-4 relative z-10 text-white">
      <p> Selected Seat: {{ mySeat.seat_number }} (Green)</p>
      <p class="flex">Available Seats: <span class="ml-1 max-w-36 truncate">{{ AvailSeatNumbers.toString() || "None " }}</span> (Blue)</p>
    </div>
    <div class="mt-4 p-4 h-stretch bg-white overflow-y-auto w-stretch m-3 rounded-3xl shadow-md relative z-10">
  <!-- First Class Section -->
      <div>
        <p class="text-2xl font-bold text-black text-center underline decoration-aa-blue decoration-2 mb-2">First Class</p>
        <div v-for="row in Math.floor(firstClass/4)" :key="row" class="flex justify-between my-1">
          <div class="w-1/3 flex gap-x-2.5">
            <div class="w-12 h-12 rounded-xl flex"
                 :class="mySeat.seat_number == `${row}A` ? 'bg-green-500' : isAvailable(`${row}A`) ? 'bg-aa-blue' : 'bg-gray-500'"
            >
              <p class="my-auto mx-auto text-white font-semibold">A</p>
            </div>
            <div class="w-12 h-12 rounded-xl flex"
                 :class="mySeat.seat_number == `${row}B` ? 'bg-green-500' : isAvailable(`${row}B`) ? 'bg-aa-blue' : 'bg-gray-500'">
              <p class="my-auto mx-auto text-white font-semibold">B</p>
            </div>
          </div>
          <div>
            <h4 class="font-semibold text-gray-500 text-2xl mt-1.5">{{ row }}</h4>
          </div>
          <div class="w-1/3 flex gap-x-2.5">
            <div class="w-12 h-12  rounded-xl flex"
                 :class="mySeat.seat_number == `${row}E` ? 'bg-green-500' : isAvailable(`${row}E`) ? 'bg-aa-blue' : 'bg-gray-500'">
              <p class="my-auto mx-auto text-white font-semibold">E</p>
            </div>
            <div class="w-12 h-12 rounded-xl flex"
                 :class="mySeat.seat_number == `${row}F` ? 'bg-green-500' : isAvailable(`${row}F`) ? 'bg-aa-blue' : 'bg-gray-500'">
              <p class="my-auto mx-auto text-white font-semibold">F</p>
            </div>
          </div>
        </div>
        <p class="text-2xl font-bold text-black text-center underline decoration-aa-blue decoration-2 mb-2 mt-4">Economy Class</p>
        <div v-for="row in Math.floor(economyClass/6)" :key="row" class="flex justify-between my-1">
          <div class="w-3/7 flex gap-x-1.5">
            <div class="w-10 h-10 rounded-xl flex"
                 :class="mySeat.seat_number == `${row + offset}A` ? 'bg-green-500' : isAvailable(`${row + offset}A`) ? 'bg-aa-blue' : 'bg-gray-500'">
              <p class="my-auto mx-auto text-white font-semibold">A</p>
            </div>
            <div class="w-10 h-10 rounded-xl flex"
                 :class="mySeat.seat_number == `${row + offset}B` ? 'bg-green-500' : isAvailable(`${row + offset}B`) ? 'bg-aa-blue' : 'bg-gray-500'">
              <p class="my-auto mx-auto text-white font-semibold">B</p>
            </div>
            <div class="w-10 h-10 rounded-xl flex"
                 :class="mySeat.seat_number == `${row + offset}C` ? 'bg-green-500' : isAvailable(`${row + offset}C`) ? 'bg-aa-blue' : 'bg-gray-500'">
              <p class="my-auto mx-auto text-white font-semibold">C</p>
            </div>
          </div>
          <div>
            <h4 class="font-semibold text-gray-500 text-2xl mt-1.5">{{ row + offset }}</h4>
          </div>
          <div class="w-3/7 flex gap-x-1.5">
            <div class="w-10 h-10 rounded-xl flex"
                 :class="mySeat.seat_number == `${row + offset}D` ? 'bg-green-500' : isAvailable(`${row + offset}D`) ? 'bg-aa-blue' : 'bg-gray-500'">
              <p class="my-auto mx-auto text-white font-semibold">D</p>
            </div>
            <div class="w-10 h-10 rounded-xl flex"
                 :class="mySeat.seat_number == `${row + offset}E` ? 'bg-green-500' : isAvailable(`${row + offset}E`) ? 'bg-aa-blue' : 'bg-gray-500'">
              <p class="my-auto mx-auto text-white font-semibold">E</p>
            </div>
            <div class="w-10 h-10 rounded-xl flex"
                 :class="mySeat.seat_number == `${row + offset}F` ? 'bg-green-500' : isAvailable(`${row + offset}F`) ? 'bg-aa-blue' : 'bg-gray-500'">
              <p class="my-auto mx-auto text-white font-semibold">F</p>
            </div>
          </div>
        </div>
      <!-- Economy Class Section -->
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, ref} from 'vue';
import {Seat} from "../types";
import getCityFromAirportCode from "../utils/airportcodes.ts";

const firstClass = ref<number>(16);
const offset = ref<number>(Math.floor(firstClass.value / 4));
const economyClass = ref<number>(144);

const props = defineProps(['selectedSeat', 'emptySeats']);

const mySeat = ref<Seat>(props.selectedSeat);
const availableSeats = ref<Seat[]>([props.selectedSeat]);

const AvailSeatNumbers = computed(() => {
  const seatNumbers: string[] = [];
  availableSeats.value.forEach((seat) => {
    if (seat.seat_number != mySeat.value.seat_number) {
      seatNumbers.push(seat.seat_number);
    }
  })

  return seatNumbers;
})

const isAvailable = (seat: string) => {
  return AvailSeatNumbers.value.includes(seat);
}
</script>
