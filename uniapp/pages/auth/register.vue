<template>
  <view class="container">
    <u-form>
      <u-form-item label="用户名" required>
        <u-input v-model="username" placeholder="请输入用户名" name="username" />
      </u-form-item>
      <u-form-item label="密码" required>
        <u-input v-model="password" type="password" name="password" placeholder="请输入密码" />
      </u-form-item>
      <u-form-item label="确认密码" required>
        <u-input v-model="confirmPassword" type="password" name="confirmPassword" placeholder="请确认密码" />
      </u-form-item>
      <u-button type="primary" @click="register">注册</u-button>
    </u-form>
  </view>
</template>

<script>
import { register } from '../../api/auth';

export default {
  data() {
    return {
      username: '',
      password: '',
      confirmPassword: ''
    };
  },
  methods: {
    register() {
      console.log('Register method called');

      if (this.password !== this.confirmPassword) {
        uni.showToast({
          title: '两次输入的密码不一致',
          icon: 'none'
        });
        return;
      }

      const formData = {
        name: this.username,
        password: this.password,
      };
	  console.log(formData)

      register(formData).then((res) => {
        console.log('注册成功');
        uni.navigateTo({
          url: '/pages/auth/login'
        });
      }).catch((error) => {
        console.error('注册失败:', error);
        uni.showToast({
          title: '注册失败，请稍后重试',
          icon: 'none'
        });
      });
    }
  }
};
</script>

<style>
.container {
  padding: 20px;
}
</style>
