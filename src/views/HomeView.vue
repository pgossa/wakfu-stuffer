<script setup lang="ts">
import TheWelcome from '../components/TheWelcome.vue'
import { inject, onBeforeMount, onMounted } from 'vue'
import listUtils from '../services/listUtils.ts'

const WURL = inject('WURL')
const version = await fetch(WURL + 'config.json')
  .then((r) => r.json())
  .then((r) => r.version)
console.log('version', version)
const items = await fetch(WURL + version + '/' + 'items.json').then((r) => r.json())
console.log('items', items)
const actions = await fetch(WURL + version + '/' + 'actions.json').then((r) => r.json())
console.log('actions:', actions)
const equipmentItemTypes = await fetch(WURL + version + '/' + 'equipmentItemTypes.json').then((r) =>
  r.json()
)
console.log('equipmentItemTypes:', equipmentItemTypes)
const itemProperties = await fetch(WURL + version + '/' + 'itemProperties.json').then((r) =>
  r.json()
)
console.log('itemProperties:', itemProperties)
</script>

<template>
  <main>
    <p>Version: {{ version }}</p>
    <li v-for="item in items">
      {{ item.title.fr }}
    </li>
    <TheWelcome />
  </main>
</template>
