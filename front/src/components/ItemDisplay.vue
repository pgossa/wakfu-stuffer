<script setup lang="ts">
import type { Item } from '@/types/itemType'
import itemUtils from '../services/itemUtils'
import Popper from "vue3-popper";

const props = defineProps({
  item: Object as () => Item
})
</script>

<template>
  <Popper hover arrow>
    <div class="flex flex-col justify-center rounded-lg border-4" v-bind:class="itemUtils.getColorByRarity(props.item!)">
      <img :src="itemUtils.getImageUrl(props.item!)" :alt="item?.Title.En">
      <p>{{ props.item?.Title.En }}</p>
    </div>
    <template #content>
      <div class="flex flex-col bg-slate-500 rounded-lg border-4" v-bind:class="itemUtils.getColorByRarity(props.item!)">
        <div v-for="effect in props.item?.EquipEffects" class="flex flex-row">
          <img :src="itemUtils.getIconByEffectId(effect.ActionId)" class="w-6"/> : {{ effect.Quantity }}
        </div>
      </div>
    </template>
  </Popper>
</template>