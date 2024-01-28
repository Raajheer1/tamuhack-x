<template>
    <div class="flex flex-col h-screen">
        <div class="w-120 h-32 bg-aa-blue rounded-b-3xl">
            <div class="text-white p-1.5">
                <h1 class="text-2xl">{{FirstName}}John {{LastName}}Doe</h1>
                <h1 class = "text-xl">AAdvantage {{aaAdvantageNum}}#8QUJ879</h1>
                <h1 class="text-xl">Member Since</h1>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">

    import {ref} from "vue"

    const FirstName = ref<string>("")

    const LastName = ref<string>("")

    const aaAdvantageNum = ref<string>("")

    // const joinDate =

    const Money = ref()

</script>