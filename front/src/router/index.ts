import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import QuickView from '../views/QuickView.vue'
import CustomView from '../views/CustomView.vue'
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
      path: '/quick',
      name: 'quick',
      component: QuickView
    },
    {
      path: '/custom',
      name: 'custom',
      component: CustomView
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
    }
  ]
})

export default router
