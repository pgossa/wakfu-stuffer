<script setup lang="ts">
import type { Build } from '@/types/buildType'
import { useBuildStore } from '@/stores/build';
import { FwbSpinner } from 'flowbite-vue'
import ItemDisplay from '../components/ItemDisplay.vue'

const buildStore = useBuildStore()
const build: Build = buildStore.getBuild
</script>

<template>
  <div class="grid grid-cols-5 gap-4">

    <div class="col-start-3 col-span-1">
      <ItemDisplay :item="build.Helmet"/>
    </div>
    <div class="col-end-6 col-span-1">
      <ItemDisplay :item="build.Accessory"/>
    </div>
    

    <div class="col-start-1 col-span-1">
      <ItemDisplay :item="build.Epaulettes"/>
    </div>
    <div class="col-start-3 col-span-1">
      <ItemDisplay :item="build.Amulet"/>
    </div>
    <div class="col-start-5 col-span-1">
      <ItemDisplay :item="build.Cloak"/>
    </div>

    <div class="col-start-3 col-span-1">
      <ItemDisplay :item="build.BreastPlate"/>
    </div>

    <div class="col-start-1 col-span-1">
      <ItemDisplay :item="build.LeftHand"/>
    </div>
    <div class="col-start-3 col-span-1">
      <ItemDisplay :item="build.Belt"/>
    </div>
    <div class="col-start-5 col-span-1">
      <ItemDisplay :item="build.RightHand"/>
    </div>

    <div class="col-start-3 col-span-1">
      <ItemDisplay :item="build.Boots"/>
    </div>

    <div class="col-start-1 col-span-1">
      <ItemDisplay :item="build.FirstWeapon"/>
    </div>
    <div v-if="build.SecondWeapon.Id" class="col-start-2 col-span-1">
      <ItemDisplay :item="build.SecondWeapon"/>
    </div>
    <div class="col-start-4 col-span-1">
      <ItemDisplay :item="build.Mount"/>
    </div>
    <div class="col-start-5 col-span-1">
      <ItemDisplay :item="build.Pet"/>
    </div>

  </div>
</template>