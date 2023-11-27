<script setup lang="ts">
import router from '@/router';
import Request from '../services/request'
import type { Build } from '@/types/buildType'
import { useBuildStore } from '@/stores/build';
import { ref } from 'vue'

const buildStore = useBuildStore()
const isRedirected = ref(false) 

async function getBasicBuild(){
  isRedirected.value = true
  const build: Build = await Request.post<Build>("http://localhost:5000/v1/api/build/weightRanking", fakeGenerateBody())
  console.log(build)
  buildStore.setBuild(build)
  router.push({ name: 'display'})
}

function fakeGenerateBody(): string {
  return `{
   "Level":230,
   "WeightSpec":{
      "HP":500,
      "HPWeight":1,
      "AP":3,
      "APWeight":6000,
      "MP":0,
      "MPWeight":3000,
      "WP":0,
      "WPWeight":0,
      "PO":0,
      "POWeight":0,
      "ArmorGiven":0,
      "ArmorGivenWeight":0,
      "ArmorReceived":0,
      "ArmorReceivedWeight":0,
      "ElementaryDamages":{
         "Fire":0,
         "FireWeight":0,
         "Earth":30,
         "EarthWeight":11,
         "Water":30,
         "WaterWeight":11,
         "Air":13,
         "AirWeight":11
      },
      "ElementaryResistances":{
         "Fire":10,
         "FireWeight":10,
         "Earth":10,
         "EarthWeight":10,
         "Water":10,
         "WaterWeight":10,
         "Air":10,
         "AirWeight":10
      },
      "CriticalChance":10,
      "CriticalChanceWeight":1000,
      "Block":0,
      "BlockWeight":0,
      "Initiative":0,
      "InitiativeWeight":0,
      "Dodge":10,
      "DodgeWeight":3,
      "Lock":0,
      "LockWeight":0,
      "Prospecting":0,
      "ProspectingWeight":0,
      "Wisdom":0,
      "WisdomWeight":0,
      "Control":0,
      "ControlWeight":0,
      "Will":0,
      "WillWeight":0,
      "MCritical":100,
      "MCriticalWeight":20,
      "RCritical":0,
      "RCriticalWeight":0,
      "MBack":0,
      "MBackWeight":0,
      "RBack":0,
      "RBackWeight":0,
      "MMelee":250,
      "MMeleeWeight":33,
      "MDistance":0,
      "MDistanceWeight":0,
      "MHeal":0,
      "MHealWeight":0,
      "MBerzerk":0,
      "MBerzerkWeight":0
   },
   "ForbiddenItems":[
      
   ],
   "MandatoryItems":[
      
   ]
}`
}
</script>

<template>
  <div v-if="! isRedirected">
    <div class="flex justify-center"> 
      <div id="modulo">
        
      </div>
    </div>
    <button type="button" @click="getBasicBuild" class="focus:outline-none text-white bg-green-700 hover:bg-green-800 focus:ring-4 focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800">Green</button>
  </div>
  <div v-else>
    <p>Searching best build...</p>
    <div role="status" class="flex justify-center">
    <svg aria-hidden="true" class="inline w-16 h-16 text-gray-200 animate-spin dark:text-gray-600 fill-green-400" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="currentColor"/>
        <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentFill"/>
    </svg>
    <span class="sr-only">Loading...</span>
</div>
  </div>

</template>