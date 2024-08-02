<template>
  <view class="container">
	  
	<view class="navbar">
	<u-icon name="arrow-left" size="20" @click="goBack" />
	<view class="navbar-title">我的资料</view>
	<view></view> <!-- 占位元素，保持标题居中 -->
	</view>
	
    <scroll-view class="list-container" scroll-y="true">
      <view v-for="request in friendRequests" :key="request.id" class="request-item">
        <!-- 使用 img 标签代替 u-avatar 组件 -->
        <image class="avatar" :src="request.url"></image>
        <view class="info">
          <text class="name">{{ request.username }}</text>
          <text class="message">{{ request.message }}</text>
        </view>
        <view class="actions">
          <u-button size="normal" type="primary" class="custom-style" @click="acceptRequest(request.id)">接受</u-button>
          <u-button size="normal" type="warning" class="custom-style" @click="rejectRequest(request.id)">拒绝</u-button>
        </view>
      </view>
    </scroll-view>
	
	<!-- 底部导航栏 -->
	<tab-bar/>
  </view>
</template>




<script>
import TabBar from '@/components/TabBar';
import { mapState } from 'vuex';
import { getApplyList } from '../../api/user';
export default {
  components: {
	  TabBar
  },
  data() {
    return {
      friendRequests: [
      ]
    };
  },
  onShow() {
  	this.$store.commit('updateTab', 1);
	getApplyList().then((res) => {
		this.friendRequests = res;
	})
  },
  computed: {
	  ...mapState(['currentTab'])
  },
  methods: {
    goBack() {
      uni.navigateTo({
      	url:'/pages/friends/FriendsList'
      });
    },
    acceptRequest(id) {
      // 处理接受请求的逻辑
      console.log('Accept request:', id);
    },
    rejectRequest(id) {
      // 处理拒绝请求的逻辑
      console.log('Reject request:', id);
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

.list-container {
  flex: 1;
  padding-top: 80px;
}

.request-item {
  display: flex;
  align-items: center;
  margin-bottom: 20px;  /* 增加底部间隔 */
  padding: 10px;
  background-color: #FFFFFF;
  border-radius: 10px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.avatar {
  width: 60px;  /* 头像宽度 */
  height: 60px; /* 头像高度 */
  border-radius: 30px;  /* 圆形头像 */
  margin-right: 20px; /* 增加头像和名称之间的间隔 */
}

.info {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.name {
  font-size: 16px;
  font-weight: bold;
  color: #333; /* 文本颜色加深，增加可读性 */
}

.message {
  font-size: 14px;
  color: #666;  /* 略微调整信息颜色 */
  margin-top: 5px; /* 增加名字和消息之间的间隔 */
}

.actions {
  display: flex;
}

.action-button {
  margin-left: 10px;  /* 按钮之间增加间隔 */
}

.custom-style {
  margin: 0 5px;
}
</style>

