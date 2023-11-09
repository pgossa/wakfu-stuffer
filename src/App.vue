<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { inject, onBeforeMount, onMounted } from 'vue'

onMounted(() => {})
</script>

<template>
  <header>
    <!-- <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" /> -->
    <!---->
    <!-- <div class="wrapper"> -->
    <!--   <HelloWorld msg="You did it!" /> -->
    <!---->
    <!--   <nav> -->
    <!--     <RouterLink to="/">Home</RouterLink> -->
    <!--     <RouterLink to="/about">About</RouterLink> -->
    <!--   </nav> -->
    <!-- </div> -->
  </header>

  <RouterView v-slot="{ Component }">
    <template v-if="Component">
      <KeepAlive>
        <Suspense>
          <component :is="Component"></component>
          <template #fallback> Loading... </template>
        </Suspense>
      </KeepAlive>
    </template>
  </RouterView>
</template>

<style scoped></style>
