// api/auth.js
import request from '@/utils/request';

// 登录函数
export function login(data) {
    return request({
        url: '/user/login',
        method: 'POST',
		header: {
			'content-type': 'application/x-www-form-urlencoded'
		},
        data: data
    });
}

// 获取用户信息
export function getUserInfo() {
    return request({
        url: '/user/my_info',
        method: 'GET'
    });
}

// 注册函数
export function register(data) {
	return request({
		url: '/user/register',
		method: 'POST',
		header: {
			'content-type': 'application/x-www-form-urlencoded'
		},
		data: data
	})
}