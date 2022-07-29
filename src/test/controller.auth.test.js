import assert from 'assert';
require('dotenv').config()

import app from '../main/module/koa/core/app';
import request from 'supertest';
import { expect } from 'chai';

//import DbServer from '../main/module/mongo/index'
const controllerAuth = require('../main/module/koa/controller/auth.controller')
const serviceUser = require('../main/module/koa/service/user.service')


let test_user_name = '_test_user';
let test_user_id;


describe('main/module/koa/controller/auth.controller', () => {
    let server;
    before(async () => {

      //加入测试用户
      let ret = await serviceUser.delete_real_by_username(test_user_name);
      ret = await serviceUser.createUser(test_user_name,test_user_name,test_user_name);
      test_user_id = ret._id;


      server = app.listen(5000);
    });
    
    after( async () => {
      ////////////////////////////////////////////
      //删除测试用户
      await serviceUser.delete_real_by_id(test_user_id);

      if (server) {
        server.close();
      }
    });

    let token = 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiJhZG1pbl9pZCIsInVzZXJfbmFtZSI6ImFkbWluIiwiZnJpZW5kX25hbWUiOiJhZG1pbiIsImlzX2FkbWluIjp0cnVlLCJpc19kZWxldGVkIjpmYWxzZSwiaWF0IjoxNjU4OTY5OTMxLCJleHAiOjE2NjE1NjE5MzF9.bC_MI4vQtBSEdT-hNTDybTyMco3at_s8nw57OR-PB2s';


    it('#取得验证码', async () => {
      await request(server)
        .post('/api/auth/captcha')
        //.send({ phone: '15505752823', password: '123456', nickName: '你说呢' })
        .expect('Content-Type', /json/)
        .expect(200)
        .expect(res => {
          let response = res.body;
          expect(response.code).to.be.equal(20000, '返回码异常');
          expect(response.data.captchaLength).to.be.equal((process.env.CAPTCHA_LEN || 4), '验证码长度异常');
          expect(response.data.text.length).to.be.equal(response.data.captchaLength, '验证码长度，异常');
        });
    });

    it('#admin登录', async () => {
      await request(server)
        .post('/api/auth/login')
        .send({ username: 'admin', password: (process.env.ADMIN_PASSWORD || 'admin')})
        .expect('Content-Type', /json/)
        .expect(200)
        .expect(res => {
          let response = res.body;
          //console.log('response:', response);
          expect(response.code).to.be.equal(20000, '返回码异常');
          expect(response.data.token).to.be.include('Bearer ', 'token验证码必需 [Bearer ] 打头');
        });
    });
    
    it('#admin错误登录', async () => {
      await request(server)
        .post('/api/auth/login')
        .send({ username: 'admin', password: 'wqerqwer3144'})
        .expect('Content-Type', /json/)
        .expect(200)
        .expect(res => {
          let response = res.body;
          //console.log('response:', response);
          expect(response.code).to.be.equal(40011, '密码错误登录返值，有问题');
        });
    });
    
    it('#admin取得登录信息', async () => {

          //expect(response.code).to.be.equal(20000, 'admin登录错误');
          // 2022-07-28 9:00 创建的token
          await request(server)
          .get('/api/auth/info')
          .set('Authorization', token)
          .expect(200)
          .expect(res => {
            let response = res.body;
            //console.log('get info :', response);
  
            expect(response.code).to.be.equal(20000, '取得登录信息，有问题');
            expect(response.data.user_name).to.be.equal('admin', '取得登录用户名必需是admin');
            expect(response.data.is_admin).to.be.equal(true, '取得登录用户admin必需是管理员');
            expect(response.data.is_deleted).to.be.equal(false, 'admin用户不能被删除');
          });
    });
    

    it('#admin登出', async () => {
      await request(server)
        .post('/api/auth/logout')
        .set('Authorization', token)
        .expect(200)
        .expect(res => {
          let response = res.body;
          //console.log('response:', response);
          expect(response.code).to.be.equal(20000, '返回码异常');
        });
    });
    

    it('#测试用户登录', async () => {
      await request(server)
        .post('/api/auth/login')
        .send({ username: test_user_name, password: test_user_name})
        .expect('Content-Type', /json/)
        .expect(200)
        .expect(res => {
          let response = res.body;
          //console.log('response:', response);
          expect(response.code).to.be.equal(20000, '返回码异常');
          expect(response.data.token).to.be.include('Bearer ', 'token验证码必需 [Bearer ] 打头');
        });
    });



    it('#测试用户取得信息', async () => {
        let ret = await request(server).post('/api/auth/login').send({ username: test_user_name, password: test_user_name});
        let token = ret._body.data.token;

        await request(server)
        .get('/api/auth/info')
        .set('Authorization', token)
        .expect(200)
        .expect(res => {
          let response = res.body;
          //console.log('get info :', response);

          expect(response.code).to.be.equal(20000, '取得登录信息，有问题');
          expect(response.data.user_name).to.be.equal(test_user_name, '取得登录用户名必需是' + test_user_name);
          expect(response.data.is_admin).to.be.equal(false, '取得登录用户测试用户不是管理员');
          expect(response.data.is_deleted).to.be.equal(false, 'admin用户不能被删除');
        });


    });



    
    /*
    it('#用户新增测试', async () => {
      await request(server)
        .post('/api/user')
        .send({ phone: '15505752823', password: '123456', nickName: '你说呢' })
        .expect('Content-Type', /json/)
        .expect(200)
        .expect(res => {
          let response = res.body;
          expect(response.code).to.be.equal(200, '用户新增异常');
        });
    });
    
    it('#用户查询测试', async () => {
      let userId = 2;
      await request(server)
        .get(`/api/user/${userId}`)
        .expect('Content-Type', /json/)
        .expect(200)
        .expect(res => {
          let response = res.body;
          expect(response.code).to.be.equal(200, '用户查询异常');
        });
    });
  
    it('#用户更新测试', async () => {
      let user = { id: 2, password: '123456789' };
      await request(server)
        .put(`/api/user`)
        .send(user)
        .expect('Content-Type', /json/)
        .expect(200)
        .expect(res => {
          let response = res.body;
          expect(response.code).to.be.equal(200, '用更新异常');
        });
    });
    */
    
  });
