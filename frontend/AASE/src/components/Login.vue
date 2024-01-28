
<template>
    
    <div class="flex mx-auto" id="login">
        
        <div class="mx-auto">

            <h1 class="text-4xl font-medium">Sign In</h1>
            <p class="text-gray-400">Start Your Journey with affordable price</p>
            
            <form>
                <ul>
                    <li>
                        <label for="">AAAdvantage #</label>
                        <blockquote></blockquote>
                        <input placeholder="Enter AAAdvantage #" v-model="aaAdvantageNum">
                    </li>
                    
                    <li>
                        <label for="">Password</label>
                        <blockquote></blockquote>
                        <input placeholder="Enter Your Password" v-model="password">
                    </li>

                    <li>
                    <button class="w-1/2 text-xl bg-blue-500 text-white py-3 rounded-xl gap-x-1">
                        Login
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

</script>