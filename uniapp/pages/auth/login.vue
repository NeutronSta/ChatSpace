<template>
  <view class="container">
    <view class="title">登录</view>
    <view class="form">
      <view class="form-item">
        <u-input style="border: 1px solid #ccc;" v-model="form.name" placeholder="请输入用户名" />
      </view>
      <view class="form-item">
        <u-input style="border: 1px solid #ccc;" v-model="form.password" type="password" placeholder="请输入密码" />
      </view>
      <view class="button-container">
        <u-button type="primary" @click="handleLogin" class="login-button">登录</u-button>
      </view>
    </view>
    <view class="register-link">
      <text @click="navigateToRegister">注册用户</text>
    </view>
  </view>
</template> 

<script>
import { login } from '../../api/auth';
import store from '@/store/index.js'; // 引入 Vuex store

export default {
  data() {
    return {
      form: {
        name: '',
        password: ''
      },
      isLoggedIn: false
    };
  },
  methods: {
    handleLogin() {
      login(this.form).then((res) => {
        console.log('返回的res是:', res);
		const token = res.token
        store.commit('SET_TOKEN', token); // 提交 token 到 Vuex store
        console.log('登录中...');
        this.checkToken(); // 登录成功后检查token并跳转
      }).catch((error) => {
        console.error('登录失败:', error);
      });
    },
    navigateToRegister() {
      console.log('navigateToRegister function called');
      uni.navigateTo({
        url: '/pages/auth/register'
      });
    },
    checkToken() {
      console.log('checkToken function called');
      const token = uni.getStorageSync('token');
      if (token) {
        this.isLoggedIn = true;
        uni.navigateTo({
          url: '/pages/friends/FriendsList'
        });
      } else {
        this.isLoggedIn = false;
      }
    }
  },
  onShow() {
    this.checkToken();
  }
};
</script>


<style scoped>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 50vh;
  padding: 20px;
}

.title {
  font-size: 24px;
  margin-bottom: 10px;
}

.form {
  width: 100%;
  max-width: 400px;
  padding: 20px;
}

.form-item {
  margin-bottom: 20px;
}

.button-container {
  display: flex;
  justify-content: center;
  margin-top: 20px;
}

.login-button {
  width: 100%;
  padding: 10px;
  border: none;
  border-radius: 25px;
  background: linear-gradient(45deg, #1e90ff, #1c86ee);
  color: white;
  font-size: 16px;
  text-align: center;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  cursor: pointer;
}

.register-link {
  margin-top: 20px;
  color: #1e90ff;
  cursor: pointer;
  font-size: 14px;
}
</style>
