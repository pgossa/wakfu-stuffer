<script setup lang="ts">
import type { Item } from '@/types/itemType'
import itemUtils from '../services/itemUtils'
const props = defineProps({
  item: Object as () => Item
})
</script>

<template>
  <div>
    <p>{{ props.item?.Title.En }}</p>
    <p>{{ props.item?.ItemTypeId }} </p> 
    <p>{{ props.item?.Id }}</p>
    <img :src="itemUtils.getImageUrl(props.item!)" :alt="item?.Title.En">
  </div>
</template>