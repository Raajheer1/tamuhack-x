<template>
  <transition name="slide-fade" mode="out-in">
    <div class="flex flex-col" :key="tab">
      <div class="px-6 mb-6 mt-8">
        <img v-if="tab == 0" src="../assets/welcome1.svg" alt="Welcome to AASE" class="w-full" />
        <img v-else-if="tab == 1" src="../assets/welcome2.svg" alt="Welcome to AASE" class="w-full" />
        <img v-else src="../assets/welcome3.svg" alt="Welcome to AASE" class="w-full" />
      </div>
      <div class="flex mx-auto gap-x-2.5 my-4" :key="tab">
        <div class="h-2 rounded-3xl" :class="tab == 0 ? 'w-16 bg-aa-blue' : 'w-3 bg-gray-300'">
        </div>
        <div class="h-2 rounded-3xl" :class="tab == 1 ? 'w-16 bg-aa-blue' : 'w-3 bg-gray-300'">
        </div>
        <div class="h-2 rounded-3xl" :class="tab == 2 ? 'w-16 bg-aa-blue' : 'w-3 bg-gray-300'">
        </div>
      </div>
      <div class="mt-8 mx-8" :key="tab">
        <h1 v-if="tab == 0" class="font-semibold text-5xl leading-tight">
          Sell your upgraded seats today!
        </h1>
        <h1 v-else-if="tab == 1" class="font-semibold text-5xl leading-tight">
          Buy a cheap seat upgrade today!
        </h1>
        <h1 v-else class="font-semibold text-5xl leading-tight">
          Welcome to AASE, get started now!
        </h1>
      </div>
      <div class="flex m-8 gap-x-4">
        <button v-if="tab == 0" class="w-1/2 text-xl text-gray-700">
          Skip
        </button>
        <button v-else @click="prevTab" class="w-1/2 text-xl text-gray-700">
          Back
        </button>
        <button @click="nextTab" class="w-1/2 text-xl bg-aa-blue text-white py-3 rounded-xl gap-x-1">
          Next
          <i class="fa-solid fa-arrow-right"></i>
        </button>
      </div>
    </div>
  </transition>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const emit = defineEmits(['next']);
const tab = ref<number>(0);

const nextTab = () => {
  if(tab.value == 2){
    emit('next');
  } else {
    tab.value++;
  }
}

const prevTab = () => {
  tab.value--;
}

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