<script setup lang="ts">
import router from '@/router';
import Request from '../services/request'
import type { Build } from '@/types/buildType'
import { useBuildStore } from '@/stores/build';
import { ref } from 'vue'
import { FwbSpinner } from 'flowbite-vue'

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
    <button @click="getBasicBuild">Click me(Basic)</button>
  </div>
  <div v-else>
    <p>Searching best build...</p>
    <div class="flex justify-center">
      <fwb-spinner size="12" />
    </div>
  </div>

</template>