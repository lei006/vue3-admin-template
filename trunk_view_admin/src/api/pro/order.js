import request from '@/utils/system/request'



let api_url_prefix = "/order";



function GetPage(params) {
  return request({
    url: api_url_prefix,
    method: 'get',
    //baseURL: '/mock',
    params
  })
}

export default {
  GetPage,
}


