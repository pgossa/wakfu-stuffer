<script setup lang="ts">
import type { Build } from '@/types/buildType'
import { useBuildStore } from '@/stores/build';
import Request from '../services/request'
import { ref, onMounted } from 'vue'

const buildStore = useBuildStore()
const build = ref("")
const buildBuild = ref({} as Build)
build.value = buildStore.getBuild
await Request.getBuild(build.value).then((res) => {
  console.log(res)
  let buildBuild = res
  console.log(buildBuild)
})
</script>

<template>
  <p>{{ buildBuild }}</p>
  <div v-if="build">
    <!-- <ul v-for="item in build.value">
      <li>{{ item.title.en }}</li>
      <li>{{ item.rarity }}</li>
    </ul> -->
  </div>
</template>