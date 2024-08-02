<template>
  <view class="chat-container">
    <!-- 顶部导航栏 -->
    <view class="navbar">
      <u-icon name="arrow-left" size="20" @click="goBack" />
      <view class="navbar-title">{{ friendName }}</view>
      <view></view> <!-- 占位元素，保持标题居中 -->
    </view>

    <!-- 消息容器 -->
    <scroll-view class="message-container" scroll-y="true">
      <view v-for="msg in messages" :key="msg.id" :class="['message', msg.type]">
        <image :src="determineAvator(msg.type)" class="avatar"/>
        <view class="bubble">
          <view class="message-text">{{ msg.text }}</view>
          <view class="message-time">{{ msg.time }}</view>
        </view>
      </view>
    </scroll-view>

    <!-- 输入区 -->
	<view class="input-container">
	  <u-textarea auto-height v-model="newMessage" class="input-area" />
	  <u-button @click="sendMessage" class="send-button">发送</u-button>
	</view>

  </view>
</template>

<script>
import { mapGetters } from 'vuex';
export default {
  data() {
    return {
      friendName: '好友名称',
      messages: [
        { id: 1, type: 'received', text: '你好！', time: '10:00', avatar: '/path/to/default-avatar.jpg' },
        { id: 2, type: 'sent', text: '你好，怎么了？', time: '10:01', avatar: '/path/to/my-avatar.jpg' }
      ],
      newMessage: '',
	  friendID: null,
	  friendAvatar: ''
    };
  },
  onLoad(options) {
    // 接收传输过来的参数
    this.friendID = options.friendID;
	this.friendName = options.friendName;
	this.friendAvatar = options.friendAvatar;
    this.loadMessages();
  },
  computed: {
	...mapGetters(['userInfo'])
  },
  methods: {
    goBack() {
      uni.navigateTo({
		  url:'/pages/friends/FriendsList'
	  });
    },
    sendMessage() {
      if (!this.newMessage.trim()) return;
      const newMsg = {
        id: this.messages.length + 1,
        type: 'sent',
        text: this.newMessage,
        time: new Date().toLocaleTimeString().slice(0, 5),
        avatar: '/path/to/my-avatar.jpg'
      };
      this.messages.push(newMsg);
      this.newMessage = '';
    },
	loadMessages() {
	  // 根据friendID加载聊天记录
	  console.log('加载好友ID为:', this.friendID, '的聊天记录');
	  // 此处模拟加载数据
	},
	determineAvator(messageType) {
		return messageType == 'sent' ? this.userInfo.avatar : this.friendAvatar
	}
  }
};
</script>

<style scoped>
.chat-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  position: fixed; /* 固定位置 */
  top: 0; /* 保持在顶部 */
  left: 0;
  right: 0;
  height: 50px; /* 适当的高度 */
  padding: 10px;
  background-color: #FFFFFF;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  z-index: 1000; /* 确保它在最上面 */
}

.message-container {
  flex: 1; /* 占据除输入栏外的所有空间 */
  overflow-y: auto; /* 允许消息滚动 */
  padding: 80px 10px 0 10px;
  background-color: #f9f9f9;
}

.message {
  display: flex;
  align-items: flex-end;
  margin-bottom: 10px;
}

.message.received {
  flex-direction: row; /* 确保布局方向 */
}

.message.sent {
  flex-direction: row-reverse; /* 发送消息，头像在右 */
}

.avatar {
  width: 40px;
  height: 40px;
  border-radius: 20px;
  margin-left: 10px; /* 为头像增加外间距 */
  margin-right: 10px;
}

.bubble {
  max-width: 70%;
  padding: 10px;
  border-radius: 15px;
  background-color: #e0f7fa;
  word-break: break-all;
  word-wrap: break-word;
  white-space: normal;
}

.message.sent .bubble {
  background-color: #b2dfdb;
}

.input-container {
  display: flex;
  align-items: center;
  padding: 10px;
  background-color: #FFFFFF;
  position: fixed;
  bottom: 0;
  border-top: 1px solid #ccc;	
  height: 80px;
  width: 100%;
}

input {
  flex: 2;
  margin-right: 10px;
  padding: 8px 10px; /* 确保输入框有足够的内边距，好看且实用 */
  border: 1px solid #ccc; /* 给输入框添加边框 */
  border-radius: 4px; /* 轻微圆角 */
}

.send-button {
	width: auto;
}

.input-area {
	border: 1px solid #ccc;
}

</style>