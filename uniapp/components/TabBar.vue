<!-- /components/TabBar.vue -->
<template>
  <u-tabbar :value="currentTab">
    <u-tabbar-item text="首页" icon="home" @click="switchTab('/pages/friends/FriendsList')"></u-tabbar-item>
    <u-tabbar-item text="好友申请" icon="man-add" :badge="applyNum" @click="switchTab('/pages/friends/FriendsRequests')"></u-tabbar-item>
    <u-tabbar-item text="我的" icon="account" @click="switchTab('/pages/settings/profile')"></u-tabbar-item>
  </u-tabbar>
</template>

<script>
import { mapState } from 'vuex';
export default {
  computed: {
	...mapState(['currentTab']),
	applyNum() {
		return this.$store.state.applyNum;
	}  
  },
  methods: {
    switchTab(path) {
      uni.navigateTo({ url: path });
    }
  }
};
</script>
