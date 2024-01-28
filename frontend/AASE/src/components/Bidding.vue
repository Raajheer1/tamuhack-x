<template>
    <div class="max-w-md mx-auto mt-8 p-4 relative">
      <div class="mb-4 flex flex-col space-y-2">
        <h1 class="text-2xl font-bold mb-2">Seat Auction for AA 1245</h1>
        <div class="flex flex-row space-x-4 justify-between">
          
        <p class="text-gray-600">Previous Seat: {{ currentSeat.seat_number }}</p>
        <p class="text-gray-600">Updated Seat: {{ newSeat.seat_number }}</p>
    </div>
        <p class="text-gray-600  mt-3">Current Bid: ${{ newSeat.current_bid }}</p>
      </div>
  
      <div class="mb-4 relative">
        <label class="block text-sm font-semibold mb-2">Enter Your Bid:</label>
        <input v-model="userBid" class="w-full h-16 px-3 py-2 border rounded-2xl relative z-10" type="number" />
        <!-- Conditional rendering based on the comparison of userBid and currentBid -->
        <div class="absolute inset-y-9 right-2 w-20 h-12 border rounded-2xl bg-white z-20 flex justify-center items-center text-green-500">
          +{{ calculateDifference() }}
        </div>
      </div>

      <div class="mb-4">
        <p>Sub Total: ${{ userBid }}</p>
        <p>Tax (7%): ${{ (userBid * 0.08).toFixed(2) }}</p>
        <p class="font-bold text-lg">Total: ${{ userBid + (userBid * 0.08).toFixed(2) }}</p>
      </div>
  
      <div class="flex flex-col space-y-4">
        <button @click="placeBid" class="bg-aa-blue text-white px-4 py-2 rounded">Place Bid ${{ userBid }}</button>
        <button @click="buyNow" class="bg-white-500 text-aa-blue px-4 py-2 rounded">Buy Now ${{ newSeat.buy_now_price }}</button>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue';
  import {Seat} from "../types";
  import {API} from "../utils/api.ts";

  const props = defineProps(['myseat', 'selectedseat']);

  const emit = defineEmits<{
    (e: 'recompute'): void
    (e: 'recompute2'): void
    (e: 'confirm'): void
  }>()

  const currentSeat = ref<Seat>(props.myseat);
  const newSeat = ref<Seat>(props.selectedseat);

  const userBid = ref<number>(newSeat.value.current_bid + 5);

  const calculateDifference = () => {
    return userBid.value - newSeat.value.current_bid;
  }

  const placeBid = async () => {
    if (userBid.value < newSeat.value.current_bid) {
      alert("Your bid must be higher than the current bid!");
    } else {
      newSeat.value.current_bid = userBid.value;
      const request = {
        bought_seat_id: newSeat.value.id,
        traded_seat_id: currentSeat.value.id,
        bid: userBid.value
      }
      const response = await API.post("/account/bid", request);
      if (response.status == 200) {
        emit('recompute')
      } else {
        alert("Bid failed!");
      }
    }
  }

  const buyNow = async () => {
    const request = {
      bought_seat_id: newSeat.value.id,
      traded_seat_id: currentSeat.value.id,
    }
    const response = await API.post("/account/buy", request);
    if (response.status == 200) {
      emit('recompute2')
      emit('confirm')
    } else {
      alert("Womp Womp!");
    }
  }
  </script>
  