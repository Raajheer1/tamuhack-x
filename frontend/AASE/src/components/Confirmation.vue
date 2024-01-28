<template>
    <div class="flex flex-col pt-10 bg-aa-blue h-screen">
        <img class="absolute mx-10 flex items-center z-5" src="../assets/globe.svg" alt="World" />
        <div class="ml-10 z-10">
            <button @click="homePage"> <!--might make it navigate to seats again if they still want to trade-->
                <i class="fa-solid fa-arrow-left text-white"></i> 
            </button>
        </div>
        <img class="pt-40 flex items-center" src="../assets/confirmation.svg" alt="confirmation checkmark"/>
        <div class="text-white text-3xl pt-10">
            <h1 class="flex justify-center ">Seat Confirmed!</h1>
            <h2 class="flex justify-center">1B</h2>
        </div>
        <button @click="accountPage" class="rounded-full bg-white py-3 my-10 mx-20 text-xl">
            See Details
        </button>
    </div>
</template>

<style scoped>
@font-face {
  font-family: 'FrutigerBold';
  src: url('../assets/FrutigerBoldRegular.ttf');
}
/* 
.custom-font {
  font-family: 'FrutigerBold';
} */
</style>

<script setup lang="ts">
    import {ref} from "vue"
    
    const emit = defineEmits(["home","account"]);
    
    const accountPage = () => {
        emit("account")
        // console.log("accountpage button works")
    }

    const homePage = () => {
        emit("home")
        // console.log("homepage button works")
    }

</script>