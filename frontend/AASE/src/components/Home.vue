<template>
  <transition name="slide-fade" mode="out-in">
    <div class="flex flex-col py-10 bg-blue-500 h-screen">
      <img class="absolute" src="../assets/globe.svg" alt="World" />
      <div class="m-10">
        <i class="fa-solid fa-less-than text-white"></i>
      </div>
    </div>
  </transition>
</template>

<script setup lang="ts">
import { ref } from 'vue';



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