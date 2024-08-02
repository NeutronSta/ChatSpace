import config from '@/env.js'
import store from '@/store/index.js'; // 引入 Vuex store

export default function request(options) {
    const url = config.apiUrl + options.url;
    const token = store.state.token; // 获取 token
    return new Promise((resolve, reject) => {
        uni.request({
            url,
            method: options.method || 'GET',
            data: options.data || {},
            header: {
                ...options.header,
                'Authorization': `Bearer ${token}`
            },
            success: (res) => {
                if (res.statusCode === 401) {
                    uni.showToast({
                        title: '登录过期，请重新登录',
                        icon: 'none'
                    });
                    uni.removeStorageSync('token');
                    reject(new Error('Token expired'));
                    // 可以选择重定向到登录页面
                    uni.redirectTo({ url: '/pages/auth/login' });
                } else if (res.statusCode === 200) {
                    resolve(res.data);
                } else {
                    uni.showToast({
                        title: `请求失败：${res.data.message || '未知错误'}`,
                        icon: 'none'
                    });
                    reject(res);
                }
            },
            fail: (err) => {
                uni.showToast({
                    title: '请求失败，请稍后重试',
                    icon: 'none'
                });
                reject(err);
            }
        });
    });
}

export function uploadFile(options) {
    const url = config.apiUrl + options.url;
    const token = store.state.token; // 获取 token
    // const filePath = options.filePath;
	const files = options.files;
    // const name = options.name;
	const uploadDetails = {
		url,
		// filePath,
		// name,
		files,
		formData: options.data || {},
		header: {
			...options.header,
			'Authorization': `Bearer ${token}`
		}
	};
	console.log(uploadDetails)
    return new Promise((resolve, reject) => {
        uni.uploadFile({
			...uploadDetails,
            success: (res) => {
                if (res.statusCode === 401) {
                    uni.showToast({
                        title: '登录过期，请重新登录',
                        icon: 'none'
                    });
                    uni.removeStorageSync('token');
                    reject(new Error('Token expired'));
                    // 可以选择重定向到登录页面
                    uni.redirectTo({ url: '/pages/auth/login' });
                } else if (res.statusCode === 200) {
                    resolve(JSON.parse(res.data)); // 确保返回的是JSON对象
                } else {
                    uni.showToast({
                        title: `上传失败：${JSON.parse(res.data).message || '未知错误'}`,
                        icon: 'none'
                    });
                    reject(res);
                }
            },
            fail: (err) => {
                uni.showToast({
                    title: '上传失败，请稍后重试',
                    icon: 'none'
                });
                reject(err);
            }
        });
    });
}
