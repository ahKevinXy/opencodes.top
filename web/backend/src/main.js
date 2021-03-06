import Vue from 'vue'
import App from './App.vue'
Vue.config.productionTip = false
// 引入封装的router
import router from '@/router/index'
// time line css
import '../node_modules/timeline-vuejs/dist/timeline-vuejs.css'
import '@/permission'
import { store } from '@/store/index'

// 路由守卫
import Bus from '@/utils/bus.js'
Vue.use(Bus)

import APlayer from '@moefe/vue-aplayer'
Vue.use(APlayer, {
  defaultCover: 'https://github.com/u3u.png',
  productionTip: true
})

import { auth } from '@/directive/auth'
// 按钮权限指令
auth(Vue)

import uploader from 'vue-simple-uploader'
Vue.use(uploader)

export default new Vue({
  render: h => h(App),
  router,
  store
}).$mount('#app')
