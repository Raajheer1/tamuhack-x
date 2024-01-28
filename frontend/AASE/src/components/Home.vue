<template>
  <Transition name="nested">
    <div class="flex flex-col pt-10 bg-aa-blue h-screen outer">
      <img class="absolute" src="../assets/globe.svg" alt="World" />
      <div class="mx-10 mt-16 mb-14">
        <div class="flex justify-between">
          <p class="mt-0.5 text-xl text-white tracking-wider">Upcoming Flights</p>
           <button @click="accountPage">
              <i class="h-5 w-5 p-2 rounded-full my-auto fa-solid fa-user text-white border border-white"></i>
           </button>

        </div>
        <h1 class="text-4xl text-white font-semibold mt-8 tracking-wider">Your next adventure awaits</h1>
      </div>
      <div class="max-h-screen overflow-auto bg-white rounded-t-3xl py-10 px-6 inner">
        <div class="flex flex-col" v-for="seat in props.seats">
          <div :key="seat.id" @click="clicked(seat.id.toString())">
            <SeatCard :seat="seat" />
          </div>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import SeatCard from "./SeatCard.vue";

const emit = defineEmits<{
  (e: 'account'): void
  (e: 'seat', id: string): void
}>()

const clicked = (id: string) => {
  emit('seat', id)
}

const accountPage = () => {
  emit('account')
}

const props = defineProps(['seats'])


</script>

<style scoped>
.nested-enter-active .inner,
.nested-leave-active .inner {
  transition: all 0.3s ease-in-out;
}

.nested-enter-from .inner,
.nested-leave-to .inner {
  transform: translateX(30px);
  opacity: 0;
}

.nested-enter-active .inner {
  transition-delay: 1s;
}
</style>