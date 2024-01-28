<!-- src/views/SeatsPage.vue -->
<template>
    <ion-page>
      <ion-header>
        <ion-toolbar>
          <ion-title>Seat Map</ion-title>
        </ion-toolbar>
      </ion-header>
      <ion-content class="ion-padding">
        <!-- Seat Map -->
        <div class="seat-map-container">
          <div v-for="row in seatRows" :key="row" class="seat-row">
            <div
              v-for="column in seatColumns"
              :key="`${row}-${column}`"
              class="seat"
            >
              {{ row }}{{ column }}
            </div>
          </div>
        </div>
  
        <!-- Toggle Drawer Button -->
        <button @click="toggleDrawer" class="bg-blue-500 text-white p-2 mt-4">
          Toggle Drawer
        </button>
  
        <!-- Bottom Drawer -->
        <div v-show="isDrawerVisible" class="bottom-drawer">
          <div class="drawer-content">
            <!-- Drawer content goes here -->
            <p>This is the drawer content</p>
          </div>
        </div>
      </ion-content>
    </ion-page>
  </template>
  
  <script setup lang="ts">
  import { ref } from "vue";
  const seatRows = ["A", "B", "C", "D", "E", "F"];
  const seatColumns = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 
                        11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
                        21, 22, 23, 24, 25, 26, 27, 28, 29, 30];
  
  const toggleDrawer = () => {
    isDrawerVisible.value = !isDrawerVisible.value;
  };
  
  const isDrawerVisible = ref(false);
  </script>
  
  <style scoped>
  
  .seat-map-container {
    display: flex;
    flex-wrap: wrap;
    gap: 10px; 
  }
  
  .seat-row {
    display: flex;
    justify-content: space-around;
  }
  
  .seat {
    width: 30px;
    height: 30px;
    border: 1px solid #ccc;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
  }
  
  @media (max-width: 640px) {
    .seat {
      width: 20px;
      height: 20px;
    }
  }
  
  .bottom-drawer {
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0;
    background-color: #fff;
    box-shadow: 0 -2px 4px rgba(0, 0, 0, 0.1);
    padding: 16px;
    transition: transform 0.3s ease-in-out;
    transform: translateY(100%);
  }
  
  .drawer-content {
    /* Make the text black */
    color: #000;
  }
  </style>
  