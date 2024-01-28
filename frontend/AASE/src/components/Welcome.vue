<template>
  <div class="flex flex-col">
    <div class="px-6 mb-6 mt-8">
      <img v-if="tab == 0" src="../assets/welcome1.svg" alt="Welcome to AASE" class="w-full" />
      <img v-else-if="tab == 1" src="../assets/welcome1.svg" alt="Welcome to AASE" class="w-full" />
      <img v-else src="../assets/welcome1.svg" alt="Welcome to AASE" class="w-full" />
    </div>
    <div class="flex mx-auto gap-x-2.5 my-4">
      <div class="w-3 h-2 bg-gray-300 rounded-3xl" :class="tab == 0 ? 'w-16 bg-blue-500' : ''">
      </div>
      <div class="w-3 h-2 bg-gray-300 rounded-3xl" :class="tab == 1 ? 'w-16 bg-blue-500' : ''">
      </div>
      <div class="w-3 h-2 bg-gray-300 rounded-3xl" :class="tab == 2 ? 'w-16 bg-blue-500' : ''">
      </div>
    </div>
    <div class="mt-8 mx-8">
      <h1 class="font-semibold text-5xl leading-tight">
        Sell your <br> upgraded seats today!
      </h1>
    </div>
    <div class="flex m-8 gap-x-4">
      <button v-if="tab == 0" class="w-1/2 text-xl text-gray-700">
        Skip
      </button>
      <button v-else @click="prevTab" class="w-1/2 text-xl text-gray-700">
        Back
      </button>
      <button @click="nextTab" class="w-1/2 text-xl bg-blue-500 text-white py-3 rounded-xl gap-x-1">
        Next
        <i class="fa-solid fa-arrow-right"></i>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const tab = ref<number>(0);

const nextTab = () => {
  tab.value++;
}

const prevTab = () => {
  tab.value--;
}

</script>

<style scoped>

</style>