// api/setting.js
import request from '@/utils/request';

export function changePWD(data) {
	return request({
		url:'/user/update_password',
		method: 'POST',
		data: data
	})
}
