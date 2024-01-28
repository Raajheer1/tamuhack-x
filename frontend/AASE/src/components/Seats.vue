<template>
  <div class="flex flex-col bg-aa-blue h-screen relative">
        <img class="absolute mx-10 flex items-center z-5" src="../assets/globe.svg" alt="World" />

        <!-- Flight Details -->
        <div class="w-120 h-32 bg-white rounded-b-3xl shadow-md p-4 relative z-10">
            <div class="flex flex-row justify-between h-full">
                <!-- Departure Details -->
                <div class="flex flex-row">
                    <div class="flex flex-col">
                        <div class="text-sm text-black flex flex-row">
                            <!-- Departure time -->
                            12:00 PM
                        </div>
                        <div class="text-sm text-opacity-70 text-gray-700 flex flex-row">
                            <!-- Departure date -->
                            28/01/2024
                        </div>
                    </div>
                    <div class="flex flex-col mx-4"> <!-- Add margin to create a symmetrical gap between time and city details -->
                        <div class="text-sm text-black flex flex-row">
                            <!-- Departure Airport Code -->
                            SEA
                        </div>
                        <div class="text-sm text-opacity-70 text-gray-700 flex flex-row">
                            <!-- Departure city -->
                            Seattle
                        </div>
                    </div>
                </div>

                <!-- Centered Plane Icon and American Airlines Logo -->
                <div class="flex items-center justify-center flex flex-col">
                    <img src="../assets/aa-tail.svg" alt="American Airlines Logo" class="h-20 w-20 rounded-full">
                    <div class="text-sm text-black flex flex-row font-weight-bold">
                        Seat: 8B
                    </div>
                </div>

                <!-- Arrival Details -->
                <div class="flex flex-row">
                    <div class="flex flex-col mx-4"> <!-- Add margin to create a symmetrical gap between time and city details -->
                        <div class="text-sm text-black flex flex-row">
                            <!-- Arrival Airport Code -->
                            ORD
                        </div>
                        <div class="text-sm text-opacity-70 text-gray-700 flex flex-row">
                            <!-- Arrival City with 70% opacity -->
                            Chicago
                        </div>
                    </div>
                    <div class="flex flex-col">
                        <div class="text-sm text-black flex flex-row">
                            <!-- Arrival time -->
                            2:00 PM
                        </div>
                        <div class="text-sm text-opacity-70 text-gray-700 flex flex-row">
                            <!-- Arrival date with 70% opacity -->
                            28/01/2024
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <!-- seatmap -->
        <div class="flex flex-col p-4 relative z-10 text-white">
            Selected Seat: 8B
        </div>
        <div class="mt-4 p-4 h-stretch bg-white overflow-y-auto w-stretch m-3 rounded-3xl shadow-md relative z-10">
    <!-- First Class Section -->
    <div>
      <p class="text-2xl font-bold text-black text-center">First Class</p>
      <div class="grid grid-cols-3">
        <!-- add boxes with rounded corners-->
        <div class="flex flex-col justify-center items-center p-2 mx-2">
          <div class="flex flex-row">
            <div class="w-10 h-10 bg-white rounded-xl p-2 border-4 mb-1"></div>
            <div class="w-10 h-10 bg-white rounded-xl p-2 border-4 mb-1"></div>

          </div>
          <div class="flex flex-row">
            <div class="w-10 h-10 bg-white rounded-xl p-2 border-4 mb-1"></div>
            <div class="w-10 h-10 bg-white rounded-xl p-2 border-4 mb-1"></div>


          </div>
          <div class="flex flex-row">
            <div class="w-10 h-10 bg-white rounded-xl p-2 border-4 mb-1"></div>
            <div class="w-10 h-10 bg-white rounded-xl p-2 border-4 mb-1"></div>
          </div>
        </div>
        <div class="grid grid-cols-3">
        <!-- add boxes with rounded corners-->
        <div class="flex flex-col justify-center items-center p-2">
          <div class="flex flex-row">
            <div class="w-10 h-10 bg-white rounded-xl p-2 border-4 mb-1"></div>
            <div class="w-10 h-10 bg-white rounded-xl p-2 border-4 mb-1"></div>

          </div>
          <div class="flex flex-row">
            <div class="w-10 h-10 bg-white rounded-xl p-2 border-4 mb-1"></div>
            <div class="w-10 h-10 bg-white rounded-xl p-2 border-4 mb-1"></div>


          </div>
          <div class="flex flex-row">
            <div class="w-10 h-10 bg-white rounded-xl p-2 border-4 mb-1"></div>
            <div class="w-10 h-10 bg-white rounded-xl p-2 border-4 mb-1"></div>
          </div>
        </div>
        </div>
        <!-- Add more seats as needed -->
       
      </div>
      </div>
    <!-- Economy Class Section -->
        <!-- A320 -->
  </div>
  </div>
</template>

<script lang="ts">

</script>
