<template>
    <div class="max-w-md mx-auto mt-8 p-4 relative">
      <div class="mb-4 flex flex-col space-y-2">
        <h1 class="text-2xl font-bold mb-2">Seat Auction for AA 1245</h1>
        <div class="flex flex-row space-x-4 justify-between">
          
        <p class="text-gray-600">Previous Seat: {{ previousSeat }}</p>
        <p class="text-gray-600">Updated Seat: {{ updatedSeat }}</p>
    </div>
        <p class="text-gray-600  mt-3">Current Bid: ${{ currentBid }}</p>
      </div>
  
      <div class="mb-4 relative">
        <label class="block text-sm font-semibold mb-2">Enter Your Bid:</label>
        <input v-model="userBid" @input="validateBid" class="w-full h-16 px-3 py-2 border rounded-2xl relative z-10" type="number" />
        <!-- Conditional rendering based on the comparison of userBid and currentBid -->
        <div class="absolute inset-y-9 right-2 w-20 h-12 border rounded-2xl bg-white z-20 flex justify-center items-center text-green-500">
          +{{ calculateDifference() }}
        </div>
      </div>

      <div class="mb-4">
        <p>Sub Total: ${{ calculateSubTotal() }}</p>
        <p>Tax (7%): ${{ calculateTax().toFixed(2) }}</p>
        <p class="font-bold text-lg">Total: ${{ calculateTotal() }}</p>
      </div>
  
      <div class="flex flex-col space-y-4">
        <button @click="placeBid" class="bg-aa-blue text-white px-4 py-2 rounded">Place Bid ${{ userBid }}</button>
        <button @click="buyNow" class="bg-white-500 text-aa-blue px-4 py-2 rounded">Buy Now ${{ buyBid }}</button>
      </div>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent } from 'vue';
  
  export default defineComponent({
    data() {
      return {
        previousSeat: 'A12',
        updatedSeat: 'B14',
        currentBid: 200,
        userBid: 200,
        buyBid: 500,
      };
    },
    methods: {
      validateBid() {
        if (this.userBid < this.currentBid) {
          this.userBid = this.currentBid;
        }
      },
      calculateDifference() {
        return this.userBid - this.currentBid;
      },
      calculateSubTotal() {
        return this.currentBid + this.calculateDifference();
      },
      calculateTax() {
        return Number((0.07 * this.calculateSubTotal()).toFixed(2));
      },
      calculateTotal() {
        return this.calculateSubTotal() + this.calculateTax();
      },
      placeBid() {
        // Add logic to handle placing a bid
        console.log('Bid placed:', this.userBid);
      },
      buyNow() {
        // Add logic to handle buying now
        console.log('Buy now clicked');
      },
    },
  });
  </script>
  