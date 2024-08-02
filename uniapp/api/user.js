// api/user.js
import request from '@/utils/request';

export function addFriend(data) {
	return request({
		url: '/relationship/add_friend',
		method: 'POST',
		header: {
			'content-type': 'application/x-www-form-urlencoded'
		},
		data
	})
}

export function getFriendList() {
	return request({
		url:'/relationship/friend_list',
		method: 'GET'
	})
}

export function getApplyList() {
	return request({
		url:'/relationship/friend_request_list',
		method: 'GET'
	})
}