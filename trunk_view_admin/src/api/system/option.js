import request from '@/utils/system/request'



let api_url_prefix = "/option";



const DeleteMany = (data) => {
  return request({
    url: api_url_prefix,
    method: 'delete',
    data,
  })
}


function GetPage(params) {
  return request({
    url: api_url_prefix,
    method: 'get',
    //baseURL: '/mock',
    params
  })
}

export default {
  DeleteMany,
  GetPage,
}

