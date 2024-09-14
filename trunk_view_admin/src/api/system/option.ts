import request from '@/utils/system/request'



let api_url_prefix = "/user";


// 获取数据api
export function getData(data: object) {
  return request({
    url: '/system/user/list',
    method: 'post',
    baseURL: '/mock',
    data
  })
}

// 新增
export function add(data: object) {
  return request({
    url: '/system/user/add',
    method: 'post',
    baseURL: '/mock',
    data
  })
}

// 编辑
export function update(data: object) {
  return request({
    url: '/system/user/update',
    method: 'post',
    baseURL: '/mock',
    data
  })
}

// 状态变更
export function updateStatus(data: object) {
  return request({
    url: '/system/user/updateStatus',
    method: 'post',
    baseURL: '/mock',
    data
  })
}

// 删除
export function del(data: object) {
  return request({
    url: '/system/user/del',
    method: 'post',
    baseURL: '/mock',
    data
  })
}


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

