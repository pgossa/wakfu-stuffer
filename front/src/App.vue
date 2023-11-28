<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import NavBar from './components/NavBar.vue'
</script>

<template>
  <NavBar></NavBar>
  <main>
    <Suspense>
      <Router-view />
    </Suspense>
  </main>
</template>

<style scoped>
main {
  padding-top: 4rem;
}
</style>
