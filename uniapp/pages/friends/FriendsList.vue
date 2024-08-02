<template>
  <view class="container">
    <!-- 遮罩层 -->
    <view v-if="showMenu" class="overlay" @click="toggleMenu"></view>

    <!-- 顶部导航栏 -->
    <view class="navbar">
      <u-icon name="arrow-left" size="20" style="opacity: 0;"/>
      <view class="navbar-title">好友列表</view>
      <u-icon name="plus" size="20" @click="toggleMenu" />
    </view>
    
    <!-- 浮动菜单 -->user
    <view :style="menuStyle" class="menu">
      <view class="menu-item" @click="navigateToAddFriend">添加好友</view>
      <view class="menu-item" @click="logout">退出登录</view>
    </view>
	
    <!-- 好友列表 -->
    <view class="page-content">
      <scroll-view class="friend-list" scroll-y="true">
        <view class="friend-item" v-for="friend in friends" :key="friend.id" @click="navigateToChat(friend)">
          <image class="friend-avatar" :src="friend.url" />
          <view class="friend-info">
            <view class="friend-name">{{ friend.username }}</view>
            <view class="friend-message">{{ friend.lastMessage }}</view>
          </view>
          <text class="friend-time">{{ friend.time }}</text>
        </view>
      </scroll-view>
    </view>
	
	<!-- 底部导航栏 -->
    <tab-bar/>
	
  </view>
</template>

<script>
import TabBar from '@/components/TabBar'
import { mapState } from 'vuex';
import { getFriendList, getApplyList } from '../../api/user';
export default {
  components: {
	TabBar
  },
  data() {
    return {
      friends: [
        { id: 1, username: '', url: '', lastMessage: '', time: '' },
      ],
	  showMenu: false
    };
  },
  onShow() {
  	this.$store.commit('updateTab', 0);
	getFriendList().then((res) => {
		this.friends = res;
	})
	getApplyList().then((res) => {
		this.$store.commit('SET_APPLY_NUM', res.length);
	})
  },
  mounted() {
	this.$store.dispatch('user/fetchUser')
	  .then(() => {
		  console.log('用户已加载')
	  })
	  .catch((err) => {
        console.error("加载用户信息失败:", error);
	  })
  },
  computed: {
	...mapState(['currentTab']),
    menuStyle() {
      if (this.showMenu) {
        return {
          opacity: '1',
          transform: 'translateY(0)',
          visibility: 'visible'
        };
      } else {
        return {
          opacity: '0',
          transform: 'translateY(-20px)',
          visibility: 'hidden'
        };
      }
    },
  },
  methods: {
    toggleMenu() {
      this.showMenu = !this.showMenu;
    },
    navigateToAddFriend() {
      this.showMenu = false;
      uni.navigateTo({ url: '/pages/friends/FriendsAdd' });
    },
    logout() {
	  uni.removeStorageSync('token');
      this.showMenu = false;
      uni.navigateTo({ url: '/pages/auth/login' });
    },
	navigateToChat(friend) {
	  uni.navigateTo({
		url: `/pages/chat/detail?friendID=${friend.id}&friendName=${friend.username}&friendAvatar=${friend.url}`
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

.overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.15); /* 半透明遮罩 */
  z-index: 1001;
}

/* 导航栏样式 */
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: fixed; /* 固定位置 */
  top: 0; /* 顶部对齐 */
  left: 0; /* 左侧对齐 */
  right: 0; /* 右侧对齐 */
  height: 80px; /* 导航栏高度 */
  background-color: #FFFFFF; /* 背景颜色 */
  box-shadow: 0 2px 4px rgba(0,0,0,0.1); /* 下阴影效果 */
  z-index: 1000; /* 确保导航栏在最上面 */
  padding: 0 10px; /* 内边距 */
}

.page-content {
  padding-top: 80px; /* 导航栏高度 */
  padding-bottom: 50px;
}

.navbar-title {
  font-size: 18px; /* 标题字体大小 */
}

.navbar-title {
  font-size: 18px;
  font-weight: bold;
}

.navbar-icons {
  display: flex;
  align-items: center;
}

.friend-list {
  flex: 1;
  padding-top: 10px;
}

.friend-item {
  display: flex;
  align-items: center;
  padding: 10px;
  background-color: #ffffff;
  border-radius: 10px;
  margin-bottom: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  margin-left: 10px;
  margin-right: 10px;
}

.friend-avatar {
  width: 40px;
  height: 40px;
  border-radius: 20px;
  margin-right: 10px;
}

.friend-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.friend-name {
  font-size: 16px;
  font-weight: bold;
}

.friend-message {
  font-size: 14px;
  color: #888888;
}

.friend-name, .friend-message {
  overflow: hidden; /* 隐藏超出内容 */
  text-overflow: ellipsis; /* 使用省略号表示溢出 */
  white-space: nowrap; /* 不换行 */
}


.friend-time {
  font-size: 12px;
  color: #cccccc;
}

.bottom-navbar {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  height: 50px; /* 根据实际内容调整高度 */
}

.bottom-navbar-item {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.bottom-navbar-icon {
  font-size: 24px;
  color: #333333;
}

.bottom-navbar-text {
  font-size: 12px;
  color: #333333;
}

.menu {
  position: absolute;
  right: 10px;
  top: 60px;
  background-color: white;
  border: 1px solid #ccc;
  border-radius: 5px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
  z-index: 1002;
  transition: opacity 0.3s, transform 0.3s, visibility 0.3s;
}

.menu-show {
  visibility: visible; /* 使菜单可见 */
  opacity: 1;
  transform: translateY(0);
}


.menu::before {
  content: '';
  position: absolute;
  top: -10px;
  right: 10px;
  border-width: 5px;
  border-style: solid;
  border-color: transparent transparent white transparent;
}

.menu-item {
  padding: 20px 10px;
  border-bottom: 1px solid #eee;
}

.menu-item:last-child {
  border-bottom: none;
}

.u-tabbar {
  position: fixed;
  bottom: 0;
  width: 100%;
  height: 50px; /* 根据需要调整 */
}
</style>
