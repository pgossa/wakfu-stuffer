<script setup lang="ts">
import router from '@/router'
import Request from '../services/request'
import type { Build } from '@/types/buildType'
import type { BuildRequest } from '@/types/buildRequetsType'
import buildUtils from '@/services/buildUtils'
import { useBuildStore } from '@/stores/build'
import { ref, onMounted, onUnmounted } from 'vue'
import Level from '../components/Level.vue'
import Restrictions from '../components/Restrictions.vue'

const buildStore = useBuildStore()
const isRedirected = ref(false)
const level = ref(230)
const mandatoryItems = ref('')
const forbiddenItems = ref('')
const statsSelected = ref([''])
const windowWidth = ref(window.innerWidth)

const updateForbiddenItems = (updatedForbiddenItem: string) => {
  forbiddenItems.value = updatedForbiddenItem
}
const updateMandatoryItems = (updatedMandatoryItem: string) => {
  mandatoryItems.value = updatedMandatoryItem
}
const handleResize = () => {
  console.log(window.innerWidth)
  windowWidth.value = window.innerWidth
}

onMounted(() => {
    window.addEventListener('resize', handleResize)
})
onUnmounted(() => {
    window.removeEventListener('resize', handleResize)
})

function toggleStat(stat: string) {
  if (statsSelected.value.includes(stat)) {
    statsSelected.value.splice(statsSelected.value.indexOf(stat), 1)
  } else {
    statsSelected.value.push(stat)
  }
}

const setLevel = (levelInput: number) => {
  level.value = levelInput
}

function clearStats() {
  statsSelected.value = []
  mandatoryItems.value = ''
  forbiddenItems.value = ''
  level.value = 230
}

async function getBasicBuild() {
  console.log(level)
  console.log(level.value)
  const buildRequest: BuildRequest = buildUtils.createEasyBuildRequestBody(
    level.value,
    statsSelected.value,
    mandatoryItems.value,
    forbiddenItems.value
  )
  isRedirected.value = true
  const build: Build = await Request.post<Build>(
    'http://localhost:5000/v1/api/build/weightRanking',
    JSON.stringify(buildRequest)
  )
  buildStore.setBuild(build)
  router.push({ name: 'display' })
}
</script>

<template>
  <button @click="handleResize">Test</button>
  <div v-if="!isRedirected">
    <div class="grid grid-rows-4 grid-flow-col justify-center gap-4" :class="{ 'grid-flow-row-dense' : windowWidth < 800 }">
      <div id="modulo" class="row-span-4 flex flex-col">
        <Level @levelEmit="setLevel" />
      </div>

      <div id="stats" class="row-span-4 flex flex-col space-y-4">
        <h1 class="mx-auto text-lg underline">Stats</h1>
        <div class="flex flex-col">
          <h1 class="text-base mx-auto">Primary stats</h1>
          <button
            type="button"
            @click="toggleStat('Action Point')"
            class="text-white bg-blue-700 hover:bg-blue-800 focus:outline-none focus:ring-4 focus:ring-blue-300 font-medium rounded-full text-sm px-5 py-2.5 text-center me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          >
            Action point
          </button>
          <button
            type="button"
            @click="toggleStat('Movement Point')"
            class="text-white bg-green-700 hover:bg-green-800 focus:outline-none focus:ring-4 focus:ring-green-300 font-medium rounded-full text-sm px-5 py-2.5 text-center me-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800"
          >
            Movement point
          </button>
          <button
            type="button"
            @click="toggleStat('Wakfu Point')"
            class="text-white bg-purple-700 hover:bg-purple-800 focus:outline-none focus:ring-4 focus:ring-purple-300 font-medium rounded-full text-sm px-5 py-2.5 text-center me-2 mb-2 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900"
          >
            Wakfu point
          </button>
        </div>
        <div class="flex flex-col">
          <h1 class="text-base mx-auto">Offensive stats</h1>
          <button
            type="button"
            @click="toggleStat('Elemental masteries')"
            class="text-white bg-blue-700 hover:bg-blue-800 focus:outline-none focus:ring-4 focus:ring-blue-300 font-medium rounded-full text-sm px-5 py-2.5 text-center me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          >
            Elementary masteries
          </button>
          <button
            type="button"
            @click="toggleStat('Melee masteries')"
            class="text-white bg-purple-700 hover:bg-purple-800 focus:outline-none focus:ring-4 focus:ring-purple-300 font-medium rounded-full text-sm px-5 py-2.5 text-center me-2 mb-2 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900"
          >
            Melee masteries
          </button>
          <button
            type="button"
            @click="toggleStat('Distance masteries')"
            class="text-white bg-green-700 hover:bg-green-800 focus:outline-none focus:ring-4 focus:ring-green-300 font-medium rounded-full text-sm px-5 py-2.5 text-center me-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800"
          >
            Distance masteries
          </button>
          <button
            type="button"
            @click="toggleStat('Rear masteries')"
            class="text-white bg-purple-700 hover:bg-purple-800 focus:outline-none focus:ring-4 focus:ring-purple-300 font-medium rounded-full text-sm px-5 py-2.5 text-center me-2 mb-2 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900"
          >
            Rear masteries
          </button>
          <button
            type="button"
            @click="toggleStat('Critical')"
            class="text-white bg-blue-700 hover:bg-blue-800 focus:outline-none focus:ring-4 focus:ring-blue-300 font-medium rounded-full text-sm px-5 py-2.5 text-center me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          >
            Critical (chance & masteries)
          </button>
          <button
            type="button"
            @click="toggleStat('Heal masteries')"
            class="text-white bg-purple-700 hover:bg-purple-800 focus:outline-none focus:ring-4 focus:ring-purple-300 font-medium rounded-full text-sm px-5 py-2.5 text-center me-2 mb-2 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900"
          >
            Heal masteries
          </button>
        </div>
        <div class="flex flex-col">
          <h1 class="text-base mx-auto">Defensive stats</h1>
          <button
            type="button"
            @click="toggleStat('Health Points')"
            class="text-white bg-blue-700 hover:bg-blue-800 focus:outline-none focus:ring-4 focus:ring-blue-300 font-medium rounded-full text-sm px-5 py-2.5 text-center me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          >
            Health Points
          </button>
          <button
            type="button"
            @click="toggleStat('Elemental resistances')"
            class="text-white bg-purple-700 hover:bg-purple-800 focus:outline-none focus:ring-4 focus:ring-purple-300 font-medium rounded-full text-sm px-5 py-2.5 text-center me-2 mb-2 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900"
          >
            Elementary resistances
          </button>
          <button
            type="button"
            @click="toggleStat('Block')"
            class="text-white bg-blue-700 hover:bg-blue-800 focus:outline-none focus:ring-4 focus:ring-blue-300 font-medium rounded-full text-sm px-5 py-2.5 text-center me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          >
            Block
          </button>
        </div>
        <div class="flex flex-col">
          <h1 class="text-base mx-auto">Secondary stats</h1>
          <button
            type="button"
            @click="toggleStat('Dodge')"
            class="text-white bg-blue-700 hover:bg-blue-800 focus:outline-none focus:ring-4 focus:ring-blue-300 font-medium rounded-full text-sm px-5 py-2.5 text-center me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          >
            Dodge
          </button>
          <button
            type="button"
            @click="toggleStat('Lock')"
            class="text-white bg-purple-700 hover:bg-purple-800 focus:outline-none focus:ring-4 focus:ring-purple-300 font-medium rounded-full text-sm px-5 py-2.5 text-center me-2 mb-2 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900"
          >
            Lock
          </button>
          <button
            type="button"
            @click="toggleStat('Initiative')"
            class="text-white bg-green-700 hover:bg-green-800 focus:outline-none focus:ring-4 focus:ring-green-300 font-medium rounded-full text-sm px-5 py-2.5 text-center me-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800"
          >
            Initiative
          </button>
          <button
            type="button"
            @click="toggleStat('Range')"
            class="text-white bg-purple-700 hover:bg-purple-800 focus:outline-none focus:ring-4 focus:ring-purple-300 font-medium rounded-full text-sm px-5 py-2.5 text-center me-2 mb-2 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900"
          >
            Range
          </button>
          <button
            type="button"
            @click="toggleStat('Control')"
            class="text-white bg-blue-700 hover:bg-blue-800 focus:outline-none focus:ring-4 focus:ring-blue-300 font-medium rounded-full text-sm px-5 py-2.5 text-center me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          >
            Control
          </button>
        </div>
      </div>

      <div id="restrictions" class="row-span-2 flex flex-col space-y-4">
        <Restrictions
          @forbiddenItemsEmit="updateForbiddenItems"
          @mandatoryItemsEmit="updateMandatoryItems"
        />
      </div>

      <div id="displayStats" class="row-span-2 flex flex-col mx-auto">
        <h1 class="text-lg underline mx-auto">Stats Selected</h1>
        <div v-for="stat in statsSelected" :key="stat" class="mx-auto">
          {{ stat }}
        </div>
      </div>
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
  <div v-else>
    <p>Searching best build...</p>
    <div role="status" class="flex justify-center">
      <svg
        aria-hidden="true"
        class="inline w-16 h-16 text-gray-200 animate-spin dark:text-gray-600 fill-green-400"
        viewBox="0 0 100 101"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
          fill="currentColor"
        />
        <path
          d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
          fill="currentFill"
        />
      </svg>
      <span class="sr-only">Loading...</span>
    </div>
  </div>
</template>
