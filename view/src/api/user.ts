import request from '@/utils/system/request'

/** 登录api */
export function loginApi(data: object) {


  console.log("loginApi", data);
  return request({
    url: '/auth/login',
    method: 'post',
    data
  })
}

/** 获取用户信息Api */
export function getInfoApi(data: object) {
  return request({
    url: '/auth/info',
    method: 'post',
    data
  })
}

/** 退出登录Api */
export function loginOutApi() {
  return request({
    url: '/auth/logout',
    method: 'post',
  })
}

/** 获取用户信息Api */
export function passwordChange(data: object) {
  return request({
    url: '/auth/password',
    method: 'post',
    data
  })
}

/** 获取登录后需要展示的菜单 */
export function getMenuApi() {
  return request({
    url: '/menu/list',
    method: 'post',
  })
}