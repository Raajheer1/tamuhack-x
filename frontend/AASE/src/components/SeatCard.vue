<template>
  <div class="w-full bg-aa-blue rounded-xl cursor-pointer">
    <div class="flex flex-col px-6 py-4">
        <div class="flex flex-row justify-between text-white">
          <div>
            <h3 class="text-lg font-bold tracking-wider">{{ seat.flight.origin }}</h3>
            <p class="text-xs text-gray-100">Seattle</p>
          </div>
          <div class="border-b-2 border-dashed border-b-white w-full h-6 mx-5"></div>
          <div class="text-right">
            <h3 class="text-lg font-bold tracking-wider">{{ seat.flight.destination }}</h3>
            <p class="text-xs text-gray-100">Chiraq</p>
          </div>
        </div>
        <div class="flex flex-row justify-between py-2 text-white">
          <div>
            <h3 class="text-md tracking-wider">{{ new Date(seat.flight.scheduled_departure_time).toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'}) }}</h3>
            <p class="text-xs text-gray-100">{{ new Date(seat.flight.scheduled_departure_time).toLocaleDateString() }}</p>
          </div>
          <div class="text-right">
            <h3 class="text-md tracking-wider">{{ new Date(seat.flight.scheduled_arrival_time).toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'}) }}</h3>
            <p class="text-xs text-gray-100">{{ new Date(seat.flight.scheduled_arrival_time).toLocaleDateString() }}</p>
          </div>
        </div>
    </div>
  </div>
  <div class="w-full bg-aa-blue rounded-xl border-t-2 border-dotted border-white mb-4 cursor-pointer">
    <div class="flex flex-col">
      <div class="flex flex-row justify-between py-3 px-6 text-white">
        <div>
          <h3 class="text-lg font-semibold tracking-wider">Seat: {{ seat.seat_number }}</h3>
        </div>
        <div class="text-right flex my-auto">
          <h3 class="text-lg font-semibold tracking-wider my-auto mr-2">AA {{ seat.flight.id }}</h3>
          <img class="h-6 my-auto" src="../assets/aa-tail.svg" alt="American Airlines" />
        </div>
      </div>
    </div>
  </div>

</template>

<script setup lang="ts">
import { Seat } from '../types';

defineProps<{ seat: Seat }>()

</script>


<style scoped>

</style>