import type { Build } from '@/types/buildType'
import { defineStore } from 'pinia'

export const useBuildStore = defineStore('build', {
  state: () => {
    return {
      data: {} as Build
    }
  },
  getters: {
    getBuild() : Build{
      return this.data
    }
  },
  actions: {
    setBuild(data: Build) {
      this.data = data
    }
  }
})