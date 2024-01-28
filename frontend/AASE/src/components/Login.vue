
<template>
    
    <div class="flex mx-auto bg-[url('./src/assets/american_bg.png')] h-screen" id="login">
        
        <div class="mx-auto px-10">
            <div class="my-10">
                <button @click="backPage">
                    <i class="fa-solid fa-arrow-left"></i> 
                </button>
            </div>
            <h1 class="text-4xl font-medium my-8 text-center">American Airlines Seat Exchange</h1>
            <h1 class="text-4xl font-medium my-6">Sign In</h1>
            <p class="text-gray-400 my-6">Start Your Journey with affordable price</p>
            
            <form v-on:submit.prevent="attemptLogin">
                <ul>
                    <li class="my-6">
                        <label>AADVANTAGE #</label>
                        <blockquote></blockquote>
                        <input name="aaAdvantageNum" placeholder="Enter AAAdvantage #" v-model="aaAdvantageNum">
                    </li>
                    
                    <li class="my-6">
                        <label>PASSWORD</label>
                        <blockquote></blockquote>
                        <input name="password" placeholder="Enter Your Password" v-model="password">
                    </li>

                    <li>
                    <button type="submit" class="w-full text-xl bg-blue-500 text-white py-3 rounded-xl gap-x-1">
                        Sign in 
                        <span><i class="fa-solid fa-circle-check"></i></span>
                    </button>
                    </li>
                </ul>
            </form>
            
        </div>
    </div>

</template>

<script setup lang="ts">
    import {ref} from "vue"
    
    const aaAdvantageNum = ref<string>("")
    const password = ref<string>("")
    
    const emit = defineEmits(["back", "home"]);
    
    const backPage = () => {
        emit("back")
    }

    const attemptLogin = () => {
        // emit("home")
        console.log("submit")
        emit("home")
    }

</script>