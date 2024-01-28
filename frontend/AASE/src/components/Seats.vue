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
                        12:00 PM
                    </div>
                    <div class="text-sm text-opacity-70 text-gray-700 flex flex-row">
                        <!-- Departure date -->
                        28/01/2024
                    </div>
                </div>
                <div class="flex flex-col mx-4"> <!-- Add margin to create a symmetrical gap between time and city details -->
                    <div class="text-sm text-black flex flex-row">
                        <!-- Departure Airport Code -->
                        SEA
                    </div>
                    <div class="text-sm text-opacity-70 text-gray-700 flex flex-row">
                        <!-- Departure city -->
                        Seattle
                    </div>
                </div>
            </div>

            <!-- Centered Plane Icon and American Airlines Logo -->
            <div class="flex items-center justify-center flex flex-col">
                <img src="../assets/aa-tail.svg" alt="American Airlines Logo" class="h-20 w-20 rounded-full">
                <div class="text-sm text-black flex flex-row font-weight-bold">
                    Seat: 8B
                </div>
            </div>

            <!-- Arrival Details -->
            <div class="flex flex-row">
                <div class="flex flex-col mx-4"> <!-- Add margin to create a symmetrical gap between time and city details -->
                    <div class="text-sm text-black flex flex-row">
                        <!-- Arrival Airport Code -->
                        ORD
                    </div>
                    <div class="text-sm text-opacity-70 text-gray-700 flex flex-row">
                        <!-- Arrival City with 70% opacity -->
                        Chicago
                    </div>
                </div>
                <div class="flex flex-col">
                    <div class="text-sm text-black flex flex-row">
                        <!-- Arrival time -->
                        2:00 PM
                    </div>
                    <div class="text-sm text-opacity-70 text-gray-700 flex flex-row">
                        <!-- Arrival date with 70% opacity -->
                        28/01/2024
                    </div>
                </div>
            </div>
        </div>
    </div>

      <!-- seatmap -->
    <div class="flex flex-col p-4 relative z-10 text-white">
      <p> Selected Seat: {{ mySeat }} (Green)</p>
      <p class="flex">Available Seats: <span class="ml-1 max-w-36 truncate">{{ availableSeats.toString() }}</span> (Blue)</p>
    </div>
    <div class="mt-4 p-4 h-stretch bg-white overflow-y-auto w-stretch m-3 rounded-3xl shadow-md relative z-10">
  <!-- First Class Section -->
      <div>
        <p class="text-2xl font-bold text-black text-center underline decoration-aa-blue decoration-2 mb-2">First Class</p>
        <div v-for="row in Math.floor(firstClass/4)" :key="row" class="flex justify-between my-1">
          <div class="w-1/3 flex gap-x-2.5">
            <div class="w-12 h-12 rounded-xl flex"
                 :class="mySeat == `${row}A` ? 'bg-green-500' : availableSeats.includes(`${row}A`) ? 'bg-aa-blue' : 'bg-gray-500'"
            >
              <p class="my-auto mx-auto text-white font-semibold">A</p>
            </div>
            <div class="w-12 h-12 rounded-xl flex"
                 :class="mySeat == `${row}B` ? 'bg-green-500' : availableSeats.includes(`${row}B`) ? 'bg-aa-blue' : 'bg-gray-500'">
              <p class="my-auto mx-auto text-white font-semibold">B</p>
            </div>
          </div>
          <div>
            <h4 class="font-semibold text-gray-500 text-2xl mt-1.5">{{ row }}</h4>
          </div>
          <div class="w-1/3 flex gap-x-2.5">
            <div class="w-12 h-12  rounded-xl flex"
                 :class="mySeat == `${row}E` ? 'bg-green-500' : availableSeats.includes(`${row}E`) ? 'bg-aa-blue' : 'bg-gray-500'">
              <p class="my-auto mx-auto text-white font-semibold">E</p>
            </div>
            <div class="w-12 h-12 rounded-xl flex"
                 :class="mySeat == `${row}F` ? 'bg-green-500' : availableSeats.includes(`${row}F`) ? 'bg-aa-blue' : 'bg-gray-500'">
              <p class="my-auto mx-auto text-white font-semibold">F</p>
            </div>
          </div>
        </div>
        <p class="text-2xl font-bold text-black text-center underline decoration-aa-blue decoration-2 mb-2 mt-4">Economy Class</p>
        <div v-for="row in Math.floor(economyClass/6)" :key="row" class="flex justify-between my-1">
          <div class="w-3/7 flex gap-x-1.5">
            <div class="w-10 h-10 rounded-xl flex"
                 :class="mySeat == `${row + offset}A` ? 'bg-green-500' : availableSeats.includes(`${row + offset}A`) ? 'bg-aa-blue' : 'bg-gray-500'">
              <p class="my-auto mx-auto text-white font-semibold">A</p>
            </div>
            <div class="w-10 h-10 rounded-xl flex"
                 :class="mySeat == `${row + offset}B` ? 'bg-green-500' : availableSeats.includes(`${row + offset}B`) ? 'bg-aa-blue' : 'bg-gray-500'">
              <p class="my-auto mx-auto text-white font-semibold">B</p>
            </div>
            <div class="w-10 h-10 rounded-xl flex"
                 :class="mySeat == `${row + offset}C` ? 'bg-green-500' : availableSeats.includes(`${row + offset}C`) ? 'bg-aa-blue' : 'bg-gray-500'">
              <p class="my-auto mx-auto text-white font-semibold">C</p>
            </div>
          </div>
          <div>
            <h4 class="font-semibold text-gray-500 text-2xl mt-1.5">{{ row + offset }}</h4>
          </div>
          <div class="w-3/7 flex gap-x-1.5">
            <div class="w-10 h-10 rounded-xl flex"
                 :class="mySeat == `${row + offset}D` ? 'bg-green-500' : availableSeats.includes(`${row + offset}D`) ? 'bg-aa-blue' : 'bg-gray-500'">
              <p class="my-auto mx-auto text-white font-semibold">D</p>
            </div>
            <div class="w-10 h-10 rounded-xl flex"
                 :class="mySeat == `${row + offset}E` ? 'bg-green-500' : availableSeats.includes(`${row + offset}E`) ? 'bg-aa-blue' : 'bg-gray-500'">
              <p class="my-auto mx-auto text-white font-semibold">E</p>
            </div>
            <div class="w-10 h-10 rounded-xl flex"
                 :class="mySeat == `${row + offset}F` ? 'bg-green-500' : availableSeats.includes(`${row + offset}F`) ? 'bg-aa-blue' : 'bg-gray-500'">
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
import {ref} from 'vue';

const firstClass = ref<number>(16);
const offset = ref<number>(Math.floor(firstClass.value / 4));
const economyClass = ref<number>(144);

const props = defineProps(['selectedSeat', 'emptySeats']);

const mySeat = ref<string>(props.selectedSeat);
const availableSeats = ref<string[]>(props.selectedSeat);

</script>
