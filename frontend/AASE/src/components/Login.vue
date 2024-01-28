<template>
    
    <div class="flex mx-auto h-screen bg-[url('./src/assets/american_bg.png')] bg-cover" id="login">
        <div class="mx-auto">
            <div class="my-10 mx-10 flex justify-between">
                <button @click="backPage">
                    <i class="fa-solid fa-arrow-left"></i> 
                </button>
                <div class="flex flex-col">
                    <img src="../assets/aaLogoFull.svg">
                    <h1 class="text-right mr-16 text-2xl">Seat Exchange</h1>
                </div>
            </div>

<!-- SEMI OLD CODE -->

            <!-- <h1 class="text-4xl font-medium my-8 text-center">American Airlines Seat Exchange</h1> -->
            <!-- <h1 class="text-4xl font-medium my-6">Sign In</h1>
            <p class="text-gray-400 my-6">Start Your Journey with affordable price</p> -->

            <!-- <div class="absolute mt-80 h-48 w-80 mx-9 bg-slate-200 rounded-2xl opacity-50 text-transparent z-0">  . </div> -->

            <div class="flex flex-col mt-80 w-80 mx-8 z-10">
                <div class="h-full w-full opacity-50">

                </div>
                <input class="mt-16 mb-3 mx-8 h-10 rounded-xl px-3" name="aaAdvantageNum" placeholder="Enter AAdvantage #" v-model="aaAdvantageNum">
                <input class="mx-8 h-10 rounded-xl px-3" name="password" placeholder="Password" v-model="password" type="password">
                <button @click="attemptLogin" type="button" class="w-stretch text-xl bg-blue-500 text-white pt-1.5 my-3 mx-8 mb-8 h-10 rounded-xl gap-x-1">
                    Sign in 
                    <!-- <span><i class="fa-solid fa-circle-check"></i></span> -->
                </button> 
            </div>          
            
        </div>
    </div>

</template>

<script setup lang="ts">
    import {ref} from "vue"
    
    const aaAdvantageNum = ref<string>("")
    const password = ref<string>("")

    const emit = defineEmits<{
      (e: 'home', id: string): void
      (e: 'back'): void
    }>()


    const backPage = () => {
        emit("back")
    }

    const attemptLogin = () => {
        if(password.value == "") {
            alert("Please enter your password")
        } else if(aaAdvantageNum.value == "") {
            alert("Please enter your AAAdvantage #")
        }
        
        // emit("home")
        console.log("submit")

        if(password.value.length >= 4 && aaAdvantageNum.value.length >= 2){

            fetch(`http://localhost:8080/account/fetchall/${aaAdvantageNum.value}`, {
                method: "get",
                mode: "cors",
                headers : {
                    "Content-Type" : "application/json"
                }
            }).then(res => {
                return res.json();
            }).then(data => {
                console.log(data);
                console.log("login sucess");
                emit("home")
            }).catch(err => {
                console.error("Failed to Login : ", err);
            })
        }else{
            console.log("invalid credentials");
        }

    }

</script>

<!-- <style scoped>

#opac50 {
    opacity: 50%;
    background-color: rgba(255, 255, 255, 122.5);
}

#opac100 {
    /* opacity: 100%; */
    background-color: rgba(255, 255, 255, 255);
}


</style> -->