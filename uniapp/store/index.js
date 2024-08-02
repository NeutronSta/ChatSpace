import Vue from 'vue'
import Vuex from 'vuex'
import user from './modules/user.js'
import getters from './getters.js'

Vue.use(Vuex);

export default new Vuex.Store({  // 注意 Store 首字母大写
	modules: {
		user
	},
	state: {
		currentTab: 0,
		applyNum: 1,
		token: uni.getStorageSync('token') || null,
	},
	mutations: {
		updateTab(state, index) {
			// console.log(`updating tab from ${state.currentTab} to ${index}`);
			if (state.currentTab !== index) {
				state.currentTab = index;
			}
		},
		SET_TOKEN(state, token) {
			state.token = token;
			uni.setStorageSync('token', token);
		},
		CLEAR_TOKEN(state) {
			state.token = null;
			uni.removeStorageSync('token');
		},
		SET_APPLY_NUM(state, num) {
			state.applyNum = num;
		}
	},
	actions: {
	},
	getters
})
