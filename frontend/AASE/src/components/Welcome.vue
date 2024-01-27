<template>
  <div class="flex">
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