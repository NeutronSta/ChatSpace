// 配置服务器的IP地址
const development = {
    apiUrl: 'http://114.116.204.137:8080',
	wssUrl: ''
};

const production = {
    apiUrl: 'http://114.116.204.137:8080',
	wssUrl: ''
};

const test = {
    apiUrl: 'http://114.116.204.137:8080',
	wssUrl: ''
};

const config = process.env.NODE_ENV === 'production' 
    ? production 
    : process.env.NODE_ENV === 'test'
    ? test
    : development;

export default config;
