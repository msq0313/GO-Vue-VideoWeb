import axios from 'axios';

//用户登录
const login = form => axios.post('/api/v1/user/login', form).then(res => res.data);
//用户登出
const logout = form => axios.delete('/api/v1/user/logout', form).then(res => res.data);
export{
    login,
    logout,
};