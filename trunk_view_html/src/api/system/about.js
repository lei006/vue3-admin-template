import request from '@/utils/system/request'



let api_url_prefix = "/about";





const SetLicense = (data) => {
  return request({
    url: api_url_prefix + `/license`,
    method: 'patch',
    data
  })
}



function GetList(params) {
  return request({
    url: api_url_prefix,
    method: 'get',
    //baseURL: '/mock',
    params
  })
}

export default {
  SetLicense,
  GetList,
}

