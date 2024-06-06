import service from '@/utils/system/request'

// 对报告的各种操作

let api_url_pre = "/setup";

  // 更新指定字段
const PatchOne = (id, field, data) => {
  let tmp_data = {field, data};
  return service({
    url: api_url_pre + `/${id}`,
    method: 'patch',
    data:tmp_data
  })
}


const GetOne = (id) => {
  return service({
    url: api_url_pre + `/${id}`,
    method: 'get',
  })
}


const GetList = (pageInfo) => {
  return service({
    url: api_url_pre,
    method: 'get',
    pageInfo
  })
}

export default {
  PatchOne,
  GetOne,
  GetList,
}
