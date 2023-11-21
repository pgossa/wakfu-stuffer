import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import BasicView from '../views/BasicView.vue'
import WeightView from '../views/WeightView.vue'
import AdvancedView from '../views/AdvancedView.vue'
import DisplayView from '@/views/DisplayView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/basic',
      name: 'basic',
      component: BasicView
    },
    {
      path: '/weight',
      name: 'weight',
      component: WeightView
    },
    {
      path: '/advanced',
      name: 'advanced',
      component: AdvancedView
    },
    {
      path: '/display',
      name: 'display',
      component: DisplayView
    },
  ]
})

export default router
