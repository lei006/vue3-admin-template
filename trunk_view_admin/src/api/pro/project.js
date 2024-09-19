import request from '@/utils/system/request'



let api_url_prefix = "/project";


// 新增
const AddOne = (data) => {
  return request({
    url: api_url_prefix,
    method: 'post',
    data
  })
}


const DeleteMany = (data) => {
  return request({
    url: api_url_prefix,
    method: 'delete',
    data,
  })
}

const PutOne = (data) => {
  return request({
    url: api_url_prefix + `/${data.id}`,
    method: 'put',
    data
  })
}

  // 更新指定字段
const PatchOne = (id, field, data) => {
  let tmp_data = {field, data};
  return request({
    url: api_url_prefix + `/${id}`,
    method: 'patch',
    data:tmp_data
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
  AddOne,
  DeleteMany,
  PutOne,
  PatchOne,
  GetOne,
  GetPage,
}

