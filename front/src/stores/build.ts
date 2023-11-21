import type { Build } from '@/types/buildType'
import { defineStore } from 'pinia'

export const useBuildStore = defineStore('build', {
  state: () => {
    return {
      data: {} as string
    }
  },
  getters: {
    getBuild() : string{
      return this.data
    }
  },
  actions: {
    setBuild(data: string) {
      this.data = data
    }
  }
})