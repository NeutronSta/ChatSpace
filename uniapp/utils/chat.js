import store from "@/store/index.js"
import config from '@/env.js'
// 连接
let socketTask = null
// 是否主动关闭连接
let meClose = false
// 地址 写你的后端连接地址
let url = config.wssUrl
let token = store.state.token || uni.getStorageSync('token');
// 重连定时器
let Time = null
// WebSocket 连接次数
let connectCount = 0;
// 最大重连次数
const maxConnectCount = 5;
// 心跳定时器
let XTime = null
// 开启连接
const sokcet = (type) => {
	//判断是否有websocet对象，有的话清空
	if (socketTask) {
		socketTask.close()
		socketTask = null;
	}
	let userId = `?userId=${store.getters['user/userInfo'].id}`
	// 进行连接
	socketTask = uni.connectSocket({
		url: url + userId,
		success(data) {
			clearInterval(XTime) //关闭心跳定时器
			console.log("创建连接!");
		}
	});
	socketTask.onOpen((res) => {
		console.log('连接成功！获取离线信息', res);
		// sendMsg("1000")
		clearInterval(Time) //关闭定时器
		clearInterval(XTime) //关闭心跳定时器
		Time = null
		// // 5秒发送一次心跳//后端检测是否在线
		XTime = setInterval(() => {
			// console.log("心跳");
			sendMsg(JSON.stringify({
				ping: true,
				type: 2
			}))
		}, 5000)

	});
	// 监听连接失败
	socketTask.onError((err) => {
		if (!meClose && Time == null) { //判断是否主动关闭进行重新连接
			console.log('连接失败，请检查');
			reconnect()
		}
	})
	// 监听连接关闭close
	socketTask.onClose((e) => {
		if (!meClose && Time == null) { //判断是否主动关闭进行重新连接
			console.log('连接关闭！', meClose);
			reconnect()
		}
	})
	// 监听收到信息
	socketTask.onMessage(function(res) {
		// 接收数据后回调
		let data = JSON.parse(res.data).data
		console.log('服务器内容:', data);
		store.commit("getMessage", data)
	});
}
// 重新连接
const reconnect = () => {
	connectCount++
	console.log("开始断线重连！！！！！！！！！！！");
	if (connectCount > maxConnectCount) {
		console.log('WebSocket 重连次数已达上限,开始每隔30秒重连一次');
		Time = setInterval(() => {
			sokcet()
		}, 30000)
		return;
	}
	setTimeout(() => {
		// 确保已经关闭后再重新打开
		socketTask.close()
		socketTask = '';
		console.log("重新连接中...");
		sokcet()
	}, 5000);

}
//向后端发送信息
const sendMsg = (msg) => {
	// console.log(msg);
	//通过 WebSocket 连接发送数据
	socketTask.send({
		data: msg,
		success() {
			// console.log("成功");
		},
		fail() {
			console.log("失败");
			uni.showLoading({
				title: "加载中..."
			})
			setTimeout(() => {
				uni.hideLoading()
			}, 1000)
			if (!meClose && Time == null) {
				reconnect()
			}
		},
	});
}
// 手动关闭连接
const stop = () => {
	// 主动关闭连接
	meClose = true
	socketTask.close({
		success() {
			console.log("手动关闭成功！");
			clearInterval(Time) //关闭定时器
			clearInterval(XTime) //关闭心跳定时器
			Time = null
			// 确保已经关闭
			socketTask = null;
		}
	})
}
// 导出方法
export const websocetObj = {
	sokcet, //连接
	stop, //关闭
	sendMsg, //发送
};
