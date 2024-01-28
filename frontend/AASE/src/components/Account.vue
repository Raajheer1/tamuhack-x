<template>
    <div class="flex flex-col h-screen">
        <div class="w-120 h-32 bg-aa-blue rounded-b-3xl">
            <div class="text-white p-1.5">
                <h1 class="text-2xl">{{FirstName}}John {{LastName}}Doe</h1>
                <h1 class = "text-l">AAdvantage {{aaAdvantageNum}}#8QUJ879</h1>
                <h1 class="text-l">Member Since {{joinMonth}}/{{joinYear}}</h1>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">

    import {ref} from "vue"

    const FirstName = ref<string>("")

    const LastName = ref<string>("")

    const aaAdvantageNum = ref<string>("")

    const joinDate = randomDate(new Date(1981, 5), new Date())
    const joinMonth = joinDate.getMonth()
    const joinYear = joinDate.getFullYear()

    function randomDate(start, end) {
        return new Date(start.getTime() + Math.random() * (end.getTime() - start.getTime()));
    }

    const Money = ref()

</script>