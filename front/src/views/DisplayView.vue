<script setup lang="ts">
import type { Build } from '@/types/buildType'
import { useBuildStore } from '@/stores/build';
import Request from '../services/request'
import { ref, onMounted } from 'vue'

const buildStore = useBuildStore()
const build = buildStore.getBuild
const buildBuild = ref({} as Build)
console.log("0000000000000000000000")
console.log(build)
buildBuild.value = await Request.post<Build>("http://localhost:5000/v1/api/build/weightRanking", build)
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