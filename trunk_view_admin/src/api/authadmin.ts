import request from '@/utils/system/request'



let api_url_prefix = "/auth/admin";


/** 登录api */
export function loginApi(data: object) {
  return request({
    url: api_url_prefix + '/login',
    method: 'post',
    data
  })
}

/** 获取用户信息Api */
export function getInfoApi(data: object) {
  return request({
    url: api_url_prefix + '/info',
    method: 'post',
    data
  })
}

/** 退出登录Api */
export function loginOutApi() {
  return request({
    url: api_url_prefix + '/logout',
    method: 'post',
  })
}

/** 获取用户信息Api */
export function SetPassword(data: object) {
  return request({
    url: api_url_prefix + '/setpassword',
    method: 'patch',
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
