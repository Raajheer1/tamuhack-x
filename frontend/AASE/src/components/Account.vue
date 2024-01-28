<template>
    <div class="flex flex-col h-screen">
        <!-- top information tag -->
        <div class="w-full h-32 bg-aa-blue rounded-b-3xl">
            <div class="flex justify-between">
                <div class="text-white p-1.5">
                    <h1 class="text-2xl">{{FirstName}}John {{LastName}}Doe</h1>
                    <h1 class = "text-l">AAdvantage {{aaAdvantageNum}}#8QUJ879</h1>
                    <h1 class="text-l">Member Since {{joinMonth}}/{{joinYear}}</h1>
                </div>
                <img src="../assets/aa-tail.svg" alt="American Airlines Logo" class="h-20 w-20 rounded-full mt-4 mr-4">
            </div>
        </div>
        <!-- start balance area -->
        <div class="w-stretch h-48 bg-aa-blue rounded-3xl m-3">
            hi
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