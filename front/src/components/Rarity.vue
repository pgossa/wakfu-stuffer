<script setup lang="ts">
import { ref } from 'vue'

const emit = defineEmits(['rarityEmit'])

const rarityAllowed = ref([
  'Common',
  'Unusual',
  'Rare',
  'Mythical',
  'Legendary',
  'Souvenir',
  'Relic',
  'Epic'
])

function toggleRarity(item: string) {
  if (rarityAllowed.value.includes(item)) {
    rarityAllowed.value = rarityAllowed.value.filter((i) => i !== item)
  } else {
    rarityAllowed.value.push(item)
  }
  emit('rarityEmit', rarityAllowed.value)
}
</script>

<template>
  <h1 class="mx-auto text-lg underline">Rarity</h1>
  <div class="flex space-x-1 mx-auto">
    <button
      @click="toggleRarity('Common')"
      class="w-6"
      :class="{ 'border-2 rounded border-emerald-500': rarityAllowed.includes('Common') }"
    >
      <img src="https://vertylo.github.io/wakassets/rarities/0.png" alt="Common" class="w-6" />
    </button>
    <button
      @click="toggleRarity('Unusual')"
      class="w-6"
      :class="{ 'border-2 rounded border-emerald-500': rarityAllowed.includes('Unusual') }"
    >
      <img src="https://vertylo.github.io/wakassets/rarities/1.png" alt="Unusual" class="w-6" />
    </button>
    <button
      @click="toggleRarity('Rare')"
      class="w-6"
      :class="{ 'border-2 rounded border-emerald-500': rarityAllowed.includes('Rare') }"
    >
      <img src="https://vertylo.github.io/wakassets/rarities/2.png" alt="Rare" class="w-6" />
    </button>
    <button
      @click="toggleRarity('Mythical')"
      class="w-6"
      :class="{ 'border-2 rounded border-emerald-500': rarityAllowed.includes('Mythical') }"
    >
      <img src="https://vertylo.github.io/wakassets/rarities/3.png" alt="Mythical" class="w-6" />
    </button>
    <button
      @click="toggleRarity('Legendary')"
      class="w-6"
      :class="{ 'border-2 rounded border-emerald-500': rarityAllowed.includes('Legendary') }"
    >
      <img src="https://vertylo.github.io/wakassets/rarities/4.png" alt="Legendary" class="w-6" />
    </button>
    <button
      @click="toggleRarity('Souvenir')"
      class="w-6"
      :class="{ 'border-2 rounded border-emerald-500': rarityAllowed.includes('Souvenir') }"
    >
      <img src="https://vertylo.github.io/wakassets/rarities/6.png" alt="Souvenir" class="w-6" />
    </button>

    <button
      @click="toggleRarity('Relic')"
      class="w-6"
      :class="{ 'border-2 rounded border-emerald-500': rarityAllowed.includes('Relic') }"
    >
      <img src="https://vertylo.github.io/wakassets/rarities/5.png" alt="Relic" class="w-6" />
    </button>
    <button
      @click="toggleRarity('Epic')"
      class="w-6"
      :class="{ 'border-2 rounded border-emerald-500': rarityAllowed.includes('Epic') }"
    >
      <img src="https://vertylo.github.io/wakassets/rarities/7.png" alt="Epic" class="w-6" />
    </button>
  </div>
</template>
