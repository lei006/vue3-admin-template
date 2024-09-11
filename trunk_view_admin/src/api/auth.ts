import request from '@/utils/system/request'

/** 登录api */
export function loginApi(data: object) {
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
    url: '/auth/out',
    method: 'post',
  })
}

/** 获取用户信息Api */
export function passwordChange(data: object) {
  return request({
    url: '/auth/passwordChange',
    method: 'post',
    data
  })
}
