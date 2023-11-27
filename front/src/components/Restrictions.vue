<script setup lang="ts">
import { ref } from 'vue'

const emit = defineEmits(['forbiddenItemsEmit', 'mandatoryItemsEmit'])

const forbiddenItems = ref('')
const mandatoryItems = ref('')
</script>

<template>
  <h1 class="mx-auto text-lg underline">Restrictions</h1>
  <div class="flex flex-col">
    <h5 class="mx-auto">Forbidden items:</h5>
    <textarea
      :value="forbiddenItems"
      @input="$emit('forbiddenItemsEmit', $event.target.value)"
      placeholder="Do not include these items"
    ></textarea>
  </div>
  <div class="flex flex-col">
    <h5 class="mx-auto">Mandatory items:</h5>
    <textarea
      :value="forbiddenItems"
      @input="$emit('mandatoryItemsEmit', $event.target.value)"
      placeholder="These items must be included"
    ></textarea>
  </div>
</template>
