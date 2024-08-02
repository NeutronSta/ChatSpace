<template>
  <view class="container">
    <!-- 顶部导航栏 -->
    <view class="navbar">
      <u-icon name="arrow-left" size="20" @click="goBack" />
      <view class="navbar-title">添加好友</view>
      <view></view> <!-- 占位元素，保持标题居中 -->
    </view>

    <!-- 添加好友表单 -->
    <view class="form-container">
      <input v-model="username" placeholder="输入好友ID或名称" class="input-field"/>
      <textarea v-model="message" placeholder="想对TA说点什么" class="textarea-field"></textarea>
      <button @click="addFriend" class="submit-btn">添加好友</button>
    </view>
  </view>
</template>

<script>
import { addFriend } from '../../api/user';
export default {
  data() {
    return {
      username: '', // 绑定输入框的数据
      message: '' // 用于存储交友信息的数据
    };
  },
  methods: {
    goBack() {
      // 返回上一页
      uni.navigateTo({
      	url:'/pages/friends/FriendsList'
      });
    },
    addFriend() {
      if (!this.username) {
        uni.showToast({
          title: '请输入好友ID或名称',
          icon: 'none'
        });
        return;
      }
      console.log('添加好友:', this.username, '附言:', this.message);
      addFriend({username: this.username})
	  .then((res) => {
		  uni.showToast({
		    title: '添加请求已发送',
		    icon: 'success'
		  });
	  })
	  .catch((err) => {
		  console.error('好友添加失败', err);
	  })
      this.username = ''; // 清空输入框
      this.message = ''; // 清空消息框
    }
  }
};
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  height: 80px;
  background-color: #FFFFFF;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  z-index: 1000;
  padding: 0 10px;
}

.navbar-title {
  font-size: 18px;
  font-weight: bold;
}

.form-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding-top: 80px;
}

.input-field, .textarea-field {
  width: 90%;
  padding: 10px;
  margin: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

.textarea-field {
  height: 100px; /* 设置合适的高度 */
}

.submit-btn {
  width: 90%;
  padding: 10px;
  background-color: #007BFF;
  color: white;
  border: none;
  border-radius: 5px;
  margin-top: 10px;
}
</style>
