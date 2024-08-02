<template>
  <view class="container">
    <!-- 顶部导航栏 -->
    <view class="navbar">
      <u-icon name="arrow-left" size="20" @click="goBack" />
      <view class="navbar-title">我的资料</view>
      <view></view> <!-- 占位元素，保持标题居中 -->
    </view>

    <!-- 个人资料显示区域 -->
    <view class="profile-container">
      <view class="avatar-container">
	    <image class="avatar" :src="userInfo.avatar" />
        <u-button class="custom-style" size="normal" type="primary" @click="editAvatar">修改头像</u-button>
      </view>
      <view class="info-item">
        <text class="info-label">昵称: {{ userInfo.name }}</text>
        <u-button class="custom-style" size="normal" type="primary" @click="editNickname">修改</u-button>
      </view>
      <u-button size="normal" type="primary" @click="navigateToChangePassword">修改密码</u-button>
    </view>
	
	<tab-bar />
	
  </view>
</template>


<script>
import TabBar from '@/components/TabBar'
import { uploadFile } from '../../utils/request';
import { mapGetters } from 'vuex';
export default {
  components: {
	  TabBar
  },
  data() {
    return {
    };
  },
  onShow() {
  	this.$store.commit('updateTab', 2)
  },
  computed: {
	  ...mapGetters(['userInfo'])
  },
  methods: {
    goBack() {
      uni.navigateTo({
      	url:'/pages/friends/FriendsList'
      })
    },
    editAvatar() {
      uni.chooseImage({
        count: 1,
        sizeType: ['original', 'compressed'],
        sourceType: ['album', 'camera'],
        success: (chooseImageRes) => {
          const tempFilePaths = chooseImageRes.tempFilePaths;
		  const tempFiles = chooseImageRes.tempFiles;
		  console.log(tempFiles)
          uploadFile({
            url: '/user/avatar', // 上传接口路径
            // filePath: tempFilePaths[0],
			files: tempFiles,
            // name: 'file',
          }).then((data) => {
            this.user.avatar = data.url; // 假设服务器返回新头像的URL
			this.$store.commit('SET_USER',this.user);
            uni.showToast({
              title: '头像更新成功',
              icon: 'success'
            });
          }).catch((error) => {
            console.error('Upload failed:', error);
          });
        }
      });
    },
    editNickname() {
      // 触发修改昵称逻辑
    },
    navigateToChangePassword() {
      uni.navigateTo({ url: '/pages/settings/ChangePassword' }); // 导航到修改密码页面
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

.custom-style {
	width: 100px;
	margin-left: 0;
	margin-right: 0px;
}

.profile-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 90px; /* 保留顶部导航栏空间 */
  margin: 90px 15px 0 15px;
}

.avatar-container {
  width: 100%;
  justify-content: space-between;
  display: flex;
  align-items: center;
  margin-bottom: 20px;
}

.avatar {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  margin-bottom: 10px; /* 头像下方留空隙 */
}

.info-item {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
  width: 100%;
  justify-content: space-between;
}

.info-label {
  font-size: 16px;
  color: #333;
}

.info-content {
  font-size: 16px;
  flex-grow: 1;
  text-align: right;
  margin-right: 10px;
}

.edit-btn, .change-password-btn {
  background-color: #007BFF;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 5px;
}
</style>
