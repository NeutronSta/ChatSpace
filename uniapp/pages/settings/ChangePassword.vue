<template>
  <view class="container">
    <!-- 顶部导航栏 -->
    <view class="navbar">
      <u-icon name="arrow-left" size="20" @click="goBack" />
      <view class="navbar-title">修改密码</view>
      <u-icon style="opacity: 0;" name="arrow-left" size="20"/>
    </view>

    <!-- 修改密码表单 -->
    <view class="form-container">
      <view class="input-group">
        <u-input v-model="password.oldPassword" password placeholder="请输入旧密码" />
      </view>
      <view class="input-group">
        <u-input v-model="password.newPassword" password placeholder="请输入新密码" />
      </view>
      <view class="input-group">
        <u-input v-model="password.confirmNewPassword" password placeholder="请确认新密码" />
      </view>
      <u-button type="primary" @click="changePassword" class="submit-btn">确认修改</u-button>
    </view>
  </view>
</template>

<script>
import { changePWD } from '../../api/setting';

export default {
  data() {
    return {
      password: {
        oldPassword: '',
        newPassword: '',
        confirmNewPassword: ''
      }
    };
  },
  methods: {
    goBack() {
      uni.navigateTo({
        url: '/pages/settings/profile'
      });
    },
    changePassword() {
      const _editpwd = {
        old_password: this.password.oldPassword,
        new_password: this.password.newPassword
      };
	  console.log(_editpwd)
      // 调用修改密码接口
      changePWD(_editpwd).then((res) => {
        // 修改成功后的处理
        uni.showToast({
          title: '密码修改成功',
          icon: 'success'
        });

        setTimeout(() => {
          // 等待两秒后执行
          uni.removeStorageSync('token'); // 移除本地 token
          uni.reLaunch({
            url: '/pages/auth/login' // 跳转到登录页面
          });
        }, 2000); // 等待时间：2000毫秒（即2秒）
      }).catch((error) => {
        console.error('密码修改失败', error);
        uni.showToast({
          title: '密码修改失败，请稍后重试',
          icon: 'none'
        });
      });
    }
  }
};
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
}

.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  height: 60px;
  background-color: #FFFFFF;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  padding: 0 10px;
  z-index: 1000;
}

.nav-placeholder {
  flex: 1;
  opacity: 0; /* 完全透明 */
}

.navbar-title {
  font-size: 18px;
  font-weight: bold;
  flex: 2; /* 使标题更加容易居中 */
  text-align: center; /* 确保标题文本居中 */
}

.form-container {
  display: flex;
  flex-direction: column;
  padding: 20px;
  margin-top: 60px; /* 保留顶部导航栏的空间 */
  box-sizing: border-box; /* 确保padding不会影响到容器的总宽度 */
}

.input-group {
  margin-bottom: 15px;
}

.submit-btn {
  margin-top: 20px;
}
</style>
