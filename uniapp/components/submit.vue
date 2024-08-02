<template>
	<view>
		<view class="submit">
			<view class="submit-chat">
				<textarea auto-height="true" cols="30" rows="10"></textarea>
			</view>
			<view class="bt-img">
				<image src="@/static/add.svg" mode=""></image>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		name:"submit",
		data() {
			return {
				
			};
		}
	}
</script>

<style lang="scss">
.submit {
	background: rgba(244,244,244,0.96);
	border-top: 1px solid $uni-border-color;
	width: 100%;
	height: 100rpx;
	position: fixed;
	bottom: 0;
	z-index: 100;
}
</style>