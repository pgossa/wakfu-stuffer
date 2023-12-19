<script setup lang="ts">
import Level from '../components/Level.vue'
import Restrictions from '../components/Restrictions.vue'
import { ref } from 'vue'

const level = ref(230)
const mandatoryItems = ref('')
const forbiddenItems = ref('')

const setLevel = (levelInput: number) => {
  level.value = levelInput
  console.log(level.value)
}

const updateForbiddenItems = (updatedForbiddenItem: string) => {
  forbiddenItems.value = updatedForbiddenItem
}

const updateMandatoryItems = (updatedMandatoryItem: string) => {
  mandatoryItems.value = updatedMandatoryItem
}
</script>

<template>
  <div v-if="!isRedirected">
    <div class="grid grid-rows-4 grid-flow-col justify-center gap-4">
      <div id="modulo" class="row-span-4 flex flex-col">
        <Level @levelEmit="setLevel" />
      </div>
    </div>

    <div id="restrictions" class="row-span-2 flex flex-col space-y-4">
      <Restrictions
        @forbiddenItemsEmit="updateForbiddenItems"
        @mandatoryItemsEmit="updateMandatoryItems"
      />
    </div>
    <div class="flex justify-center space-x-20 pt-8">
      <button
        type="button"
        @click="getBasicBuild"
        class="focus:outline-none text-white bg-green-700 hover:bg-green-800 focus:ring-4 focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800"
      >
        Calculate Build
      </button>
      <button
        @click="clearStats"
        type="button"
        class="focus:outline-none text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-900"
      >
        Reset Selection
      </button>
    </div>
  </div>
</template>
