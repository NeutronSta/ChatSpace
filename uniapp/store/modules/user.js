// store/modules/user.js
import { getUserInfo } from '../../api/auth';

const state = {
  userInfo: {
	  name: '',
	  avatar: '',
	  id: ''
  }
};

const mutations = {
  SET_USER_INFO(state, userInfo) {
    state.userInfo = userInfo;
  }
};

const actions = {
	fetchUser({ commit, state }) {
		return getUserInfo().then(userInfo => {
		  commit('SET_USER_INFO', userInfo);
		  console.log('用户信息已更新:', state.userInfo);
		}).catch(error => {
		  console.error('Error fetching user info:', error);
		  commit('CLEAR_TOKEN');  // 清除 token 和用户信息
		  uni.redirectTo({ url: '/pages/auth/login' });
		});
	}
};


export default {
  namespaced: true,
  state,
  mutations,
  actions,
};
