import { MockMethod } from 'vite-plugin-mock'
const users = [
  { username: 'admin', password: '123456', token: 'admin', info: {
    name: '系统管理员'
  }},
  { username: 'editor', password: '123456', token: 'editor', info: {
    name: '编辑人员'
  }},
  { username: 'test', password: '123456', token: 'test', info: {
    name: '测试人员'
  }},
]
export default [
  {
    url: `/mock/auth/admin/login`,
    method: 'post',
    response: ({ body }) => {
      const user = users.find(user => {
        return body.name === user.username && body.password === user.password
      })
      if (user) {
        return {
          code: 200,
          data: {
            token: user.token,
          },
        };
      } else {
        return {
          code: 401,
          data: {},
          msg: '用户名或密码错误'
        };
      }
      
    }
  },
  {
    url: `/mock/auth/admin/info`,
    method: 'post',
    response: ({ body }) => {
      const { token } = body
      const info = users.find(user => {
        return user.token === token
      }).info
      if (info) {
        return {
          code: 200,
          data: {
            info: info
          },
        };
      } else {
        return {
          code: 403,
          data: {},
          msg: '无访问权限'
        };
      }
      
    }
  },
  {
    url: `/mock/auth/admin/logout`,
    method: 'post',
    response: () => {
      return {
        code: 200,
        data: {},
        msg: 'success'
      };
    }
  },
  {
    url: `/mock/auth/admin/SetPassword`,
    method: 'post',
    response: () => {
      return {
        code: 200,
        data: {},
        msg: 'success'
      };
    }
  },
] as MockMethod[]