import Vue from 'vue'
import VueRouter from 'vue-router'
import questionPage from './pages/question_page/index.vue'

// 要告诉 vue 使用 vueRouter
Vue.use(VueRouter)

const routes = [
  {
    path: '/questionaire',
    component: questionPage
  }
]

var router = new VueRouter({
  routes
})
export default router
