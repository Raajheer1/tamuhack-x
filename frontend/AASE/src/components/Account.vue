<template>
    <div class="flex flex-col h-screen">
        <!-- top information tag -->
        <div class="w-full h-32 bg-aa-blue rounded-b-3xl">
            <div class="flex justify-between">
                <div class="text-white pl-3 py-5">
                    <h1 class="text-2xl">{{ user.first_name }} {{ user.last_name }}</h1>
                    <h1 class = "text-l">AAdvantage #{{ user.id }}</h1>
                    <h1 class="text-l">Member Since {{joinMonth}}/{{joinYear}}</h1>
                </div>
                <img src="../assets/aa-tail.svg" alt="American Airlines Logo" class="h-20 w-20 rounded-full m-5">
            </div>
        </div>
        <!-- start balance area -->
        <div class="w-stretch h-48 bg-slate-200 rounded-3xl m-3 p-2 px-3">
            <h2>Current Balance</h2>
            <h1 class="text-4xl">${{ user.money }}</h1>
            <div class="flex justify-between">
                <button @click="decrement" class="rounded-full h-16 w-72 bg-white my-3 mx-3 text-xl">Withdraw</button>
                <button @click="increment" class="rounded-full h-16 w-72 bg-black my-3 mx-3 text-xl text-white">Add Funds</button>
            </div>
        </div>
        <!-- starting the scrolling -->
        <div class="h-stretch overflow-auto bg-slate-200 rounded-3xl pb-3 pt-3 px-6 inner w-stretch mx-3 mt-1">
            <h1 class="flex justify-center mb-2 pr-14 text-xl">Your Previous Adventures</h1>
            <div class="flex flex-col" v-for="seat in props.seats">
                <div :key="seat.id">
                    <SeatCard :seat="seat" />
                </div>
            </div>
        </div>
        <div class="flex justify-between">
            <button @click="homePage" class="w-40 h-20 ml-6 rounded-full my-5 bg-aa-blue text-2xl text-white">Home</button>
            <button @click="loginPage" class="w-40 h-20 mr-6 rounded-full my-5 bg-aa-red text-2xl text-white">Logout</button>
        </div>
    </div>
</template> 

<script setup lang="ts">

import {Account} from "../types";
    import { ref } from 'vue';
    import SeatCard from "./SeatCard.vue";

    const emit = defineEmits<{
      (e: 'login'): void
      (e: 'home2'): void
    }>()

    const props = defineProps(['me', 'seats']);
    const user = ref<Account>(props.me);

    if(user.value.money > 1000){
      user.value.money = 750
    }

    const loginPage = () => {
        emit("login")
    }

    const homePage = () => {
        emit("home2")
    }

    const decrement = () => {
        user.value.money = 0
    }

    const increment = () => {
        user.value.money = user.value.money + 100
    }


    const joinDate = randomDate(new Date(1981, 5), new Date())
    const joinMonth = joinDate.getMonth()
    const joinYear = joinDate.getFullYear()

    function randomDate(start: Date, end: Date) {
        return new Date(start.getTime() + Math.random() * (end.getTime() - start.getTime()));
    }


</script>