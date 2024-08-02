// #ifndef VUE3
import Vue from 'vue'
import App from './App'
import uView from 'uview-ui'
import store from './store/index.js'  // 确保这个路径是正确的

Vue.prototype.$store = store
Vue.config.productionTip = false

Vue.use(uView)

App.mpType = 'app'

const app = new Vue({
    ...App,
	store  // 这里的 store 不需要用花括号包围
})
app.$mount()
// #endif

// #ifdef VUE3
import { createSSRApp } from 'vue'
import App from './App.vue'
export function createApp() {
  const app = createSSRApp(App)
  return {
    app
  }
}
// #endif
