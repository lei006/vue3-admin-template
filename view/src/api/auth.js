import request from '@/utils/system/request'

/** 登录api */
function login(data) {
  return request({
    url: '/auth/login',
    method: 'post',
    data
  })
}
function regedit(data) {
  return request({
    url: '/auth/regedit',
    method: 'post',
    data
  })
}

/** 获取用户信息Api */
function info() {
  return request({
    url: '/auth/info',
    method: 'get',
  })
}

/** 退出登录Api */
function logout() {
  return request({
    url: '/auth/logout',
    method: 'delete',
  })
}

/** 获取用户信息Api */
function password(data) {
  return request({
    url: '/auth/password',
    method: 'patch',
    data
  })
}

function captcha() {
  return request({
    url: '/auth/captcha',
    method: 'get'
  })
}



export default {regedit, login, logout, password, info, captcha}
