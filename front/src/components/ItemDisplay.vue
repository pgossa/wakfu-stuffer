<script setup lang="ts">
import type { Item } from '@/types/itemType'
import itemUtils from '../services/itemUtils'
import Popper from "vue3-popper";
import hp from "../assets/wakfu/hp.png"

const props = defineProps({
  item: Object as () => Item
})
</script>

<template>
  <Popper hover arrow>
    <div class="flex flex-col justify-center rounded-lg border-4" v-bind:class="itemUtils.getColorByRarity(props.item!)">
      <!-- <img :src="itemUtils.getImageUrl(props.item!)" :alt="item?.Title.En"> -->
      <p>{{ props.item?.Title.En }}</p>
    </div>
    <template #content>
      <div class="flex flex-col bg-slate-500 rounded-lg border-4" v-bind:class="itemUtils.getColorByRarity(props.item!)">
        <div v-for="effect in props.item?.EquipEffects" class="flex flex-row gap-2">
          <div class="" >{{ effect.Quantity }}</div> <div class="" >{{ itemUtils.getDescByEffect(effect) }}</div> <img :src="itemUtils.getIconByEffect(effect)" class="relative w-6 right-0"/>
        </div>
      </div>
    </template>
  </Popper>
</template>