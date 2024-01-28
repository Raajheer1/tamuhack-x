
<template>
    
    <div class="flex mx-auto h-screen bg-[url('./src/assets/american_bg.png')] bg-cover" id="login">
        <div class="mx-auto px-10">
            <div class="my-10">
                <button @click="backPage">
                    <i class="fa-solid fa-arrow-left"></i> 
                </button>
            </div>
            <h1 class="text-4xl font-medium my-8 text-center">American Airlines Seat Exchange</h1>
            <h1 class="text-4xl font-medium my-6">Sign In</h1>
            <p class="text-gray-400 my-6">Start Your Journey with affordable price</p>
            
            <div >
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
                    <button @click="attemptLogin" type="button" class="w-full text-xl bg-blue-500 text-white py-3 rounded-xl gap-x-1">
                        Sign in 
                        <span><i class="fa-solid fa-circle-check"></i></span>
                    </button>
                    </li>
                </ul>
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