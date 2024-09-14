import request from '@/utils/system/request'



let api_url_prefix = "/setup";



const DeleteMany = (data) => {
  return request({
    url: api_url_prefix,
    method: 'delete',
    data,
  })
}


const GetOne = (id) => {
  return request({
    url: api_url_prefix + `/${id}`,
    method: 'get',
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
  GetOne,
  GetPage,
}

