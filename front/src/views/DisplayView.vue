<script setup lang="ts">
import router from '@/router'
import { onBeforeMount } from 'vue'
import type { Build } from '@/types/buildType'
import { useBuildStore } from '@/stores/build'
import ItemDisplay from '../components/ItemDisplay.vue'
import HomeView from './HomeView.vue'

const buildStore = useBuildStore()
const build: Build = buildStore.getBuild

onBeforeMount(() => {
  if (build.Mount === undefined) {
    router.push({ path: '/', name: 'home', component: HomeView })
  }
})
</script>

<template>
  <div class="grid grid-cols-3 gap-4">
    <div class="col-start-1 col-span-1">
      <ItemDisplay :item="build.Mount" />
    </div>
    <div class="col-start-2 col-span-1">
      <ItemDisplay :item="build.Helmet" />
    </div>
    <div class="col-start-3 col-span-1">
      <ItemDisplay :item="build.Pet" />
    </div>

    <div class="col-start-1 col-span-1">
      <ItemDisplay :item="build.Epaulettes" />
    </div>
    <div class="col-start-2 col-span-1">
      <ItemDisplay :item="build.Amulet" />
    </div>
    <div class="col-start-3 col-span-1">
      <ItemDisplay :item="build.Cloak" />
    </div>

    <div class="col-start-2 col-span-1">
      <ItemDisplay :item="build.BreastPlate" />
    </div>

    <div class="col-start-1 col-span-1">
      <ItemDisplay :item="build.LeftHand" />
    </div>
    <div class="col-start-2 col-span-1">
      <ItemDisplay :item="build.Belt" />
    </div>
    <div class="col-start-3 col-span-1">
      <ItemDisplay :item="build.RightHand" />
    </div>

    <div class="col-start-2 col-span-1">
      <ItemDisplay :item="build.Boots" />
    </div>

    <div class="col-start-1 col-span-1">
      <ItemDisplay :item="build.FirstWeapon" />
    </div>
    <div v-if="build.SecondWeapon.Id" class="col-start-2 col-span-1">
      <ItemDisplay :item="build.SecondWeapon" />
    </div>
    <div v-else class="col-start-2 col-span-1">
      <ItemDisplay :item="build.FirstWeapon" />
    </div>
    <div class="col-start-3 col-span-1">
      <ItemDisplay :item="build.Accessory" />
    </div>
  </div>
</template>
