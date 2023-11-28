<script setup lang="ts">
import type { Item } from '@/types/itemType'
import itemUtils from '../services/itemUtils'
import Popper from "vue3-popper";
import envyclopedia from "../assets/wakfu/encyclopedia.webp";
// TODO: Properly redirect to wakfu encyclopedia

const props = defineProps({
  item: Object as () => Item
})
</script>

<template>
  <Popper  closeDelay="0">
    <div class="flex flex-col justify-center rounded-lg border-4" v-bind:class="itemUtils.getColorByRarity(props.item!)">
      <img :src="itemUtils.getImageUrl(props.item!)" class="max-w-xxs">
      <p>{{ props.item?.Title.En }}</p>
    </div>
    <template #content>
      <div class="flex flex-col bg-slate-500 rounded-lg border-4" v-bind:class="itemUtils.getColorByRarity(props.item!)">
        <div class="grid grid-cols-3 border-b-4">
          <label class="text-center border-r-4 col-start-1 col-end-3 pr-2">{{ props.item?.Title.En }}</label>
          <label class="text-center col-start-3 border-r-4 col-end-4">level{{ props.item?.Level }}</label>
          <a class="col-start-4 col-end-5 w-6" href="https://www.google.com/" target="_blank" rel="noopener noreferrer"> 
          <img class="" :src=envyclopedia /></a>
        </div>
        <div v-for="effect in props.item?.EquipEffects" class="grid grid-cols-2 gap-2">
          <div class="col-start-1 col-end-2" >{{ effect.Quantity }} {{ itemUtils.getDescByEffect(effect) }}</div> <img :src="itemUtils.getIconByEffect(effect)" class="relative w-6 col-start-3 col-end-4"/>
        </div>
      </div>
    </template>
  </Popper>
</template>

<style scoped>
.max-w-xxs{
  max-width: 5rem;
}
@media(max-width: 1024px) {
  .max-w-xxs{
    max-width: 3rem;
  }
}
</style>