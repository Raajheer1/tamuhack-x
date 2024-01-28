
<template>
    
    <div class="flex mx-auto" id="login">
        
        <div class="mx-auto">
            <div class="my-10">
                <button @click="backPage">
                    <i class="fa-solid fa-arrow-left"></i> 
                </button>
            </div>

            <h1 class="text-4xl font-medium my-6 ">Sign In</h1>
            <p class="text-gray-400 my-6">Start Your Journey with affordable price</p>
            
            <form>
                <ul>
                    <li class="my-6">
                        <label for="">AAADVANTAGE #</label>
                        <blockquote></blockquote>
                        <input placeholder="Enter AAAdvantage #" v-model="aaAdvantageNum">
                    </li>
                    
                    <li class="my-6">
                        <label for="">PASSWORD</label>
                        <blockquote></blockquote>
                        <input placeholder="Enter Your Password" v-model="password">
                    </li>

                    <li>
                    <button class="w-full text-xl bg-blue-500 text-white py-3 rounded-xl gap-x-1">
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
    
    const emit = defineEmits(['back']);
    
    const backPage = () => {
        emit("back")
    }

</script>