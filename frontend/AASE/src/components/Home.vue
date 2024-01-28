<template>
  <transition name="slide-fade" mode="out-in">
    <div class="flex flex-col pt-10 bg-aa-blue h-screen">
      <img class="absolute" src="../assets/globe.svg" alt="World" />
      <div class="mx-10 mt-16 mb-14">
        <div class="flex justify-between">
          <p class="mt-0.5 text-xl text-white tracking-wider">Upcoming Flights</p>
          <i class="h-5 w-5 p-2 rounded-full my-auto fa-solid fa-user text-white border border-white"></i>
        </div>
        <h1 class="text-4xl text-white font-semibold mt-8 tracking-wider">Your future adventure awaits</h1>
      </div>
      <div class="h-full bg-white rounded-t-3xl">

      </div>
    </div>
  </transition>
</template>

<script setup lang="ts">
// import { ref } from 'vue';


</script>

<style scoped>
.slide-fade-enter-active {
  transition: all 0.3s ease-out;
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