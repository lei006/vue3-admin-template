import assert from 'assert';
require('dotenv').config()

import app from '../main/module/koa/core/app';
import request from 'supertest';
import { expect } from 'chai';
const random = require('string-random');

const serviceUser = require('../main/module/koa/service/user.service')

const controllerUser = require('../main/module/koa/controller/user.controller')


let test_user_name = '_test_user';
let test_user_name_1 = '_test_user_1';
let test_user_id;
let test_user_1_id;
let test_user_token;
let admin_user_token;

describe('main/module/koa/controller/user.controller', () => {
    let server;
    before(async () => {
      server = app.listen(5000);

      //加入测试用户
      let ret = await serviceUser.delete_real_by_username(test_user_name);
      ret = await serviceUser.createUser(test_user_name,test_user_name,test_user_name);
      test_user_id = ret._id;

      ret = await serviceUser.delete_real_by_username(test_user_name_1);
      ret = await serviceUser.createUser(test_user_name_1,test_user_name_1,test_user_name_1);
      test_user_1_id = ret._id;

      
      ret = await request(server).post('/api/auth/login').send({ username: test_user_name, password: test_user_name});
      test_user_token = ret._body.data.token;

      ret = await request(server).post('/api/auth/login').send({ username: 'admin', password: (process.env.ADMIN_PASSWORD || 'admin')});
      admin_user_token = ret._body.data.token;

    });
    
    after(async () => {

      //删除测试用户
      await serviceUser.delete_real_by_id(test_user_id);
      await serviceUser.delete_real_by_username(test_user_name_1);
      if (server) {
        server.close();
      }
    });


    it('#管理员-增加用户', async () => {
      await request(server)
        .post('/api/user')
        .set('Authorization', admin_user_token)
        .expect('Content-Type', /json/)
        .expect(200)
        .expect(async res => {
          let response = res.body;
          //console.log('response',response);
          expect(response.code).to.be.equal(20000, '返回码异常');
          expect(response.data.is_admin).to.be.equal(false, '新加的用户,不能是管理员');
          expect(response.data.is_deleted).to.be.equal(false, '新加的用户,不能是已删除状态');

          //清理增加用户
          let tmp_user_id = response.data._id;
          await serviceUser.delete_real_by_id(tmp_user_id);

        });
    });

    it('#管理员-删除用户', async () => {

      // 1. 增加用户
      let ret = await request(server).post('/api/user').set('Authorization', admin_user_token);
      let tmp_user_id = ret._body.data._id;

      // 2. 删除用户
      await request(server)
        .delete(`/api/user/${tmp_user_id}`)
        .set('Authorization', admin_user_token)
        .expect('Content-Type', /json/)
        .expect(200)
        .expect(async res => {
          let response = res.body;
          //console.log('response',response);
          expect(response.code).to.be.equal(20000, '返回码异常');
          expect(response.data).to.be.equal(1, '删除用户数量必需为1');

          //清理增加用户
          await serviceUser.delete_real_by_id(tmp_user_id);


        });
    });

    it('#管理员-查询用户', async () => {

      // 2. 查询用户
      await request(server)
        .get(`/api/user/${test_user_id}`)
        .set('Authorization', admin_user_token)
        .expect('Content-Type', /json/)
        .expect(200)
        .expect(res => {
          let response = res.body;
          //console.log('response',response);
          expect(response.code).to.be.equal(20000, '返回码异常');
          //expect(response.data._id).to.be.equal(test_user_id, '查到的id数据,与传入的必需一致');
          expect(response.data.user_name).to.be.equal(test_user_name, '查到的user_name数据,与传入的必需一致');
          expect(response.data.password).to.be.equal(undefined, '查到的用户，不能有密码');
        });
    });


    it('#管理员-用户列表', async () => {
      await request(server)
        .get(`/api/user`)
        .set('Authorization', admin_user_token)
        .expect('Content-Type', /json/)
        .expect(200)
        .expect(res => {
          let response = res.body;
          //console.log('response',response.data);
          expect(response.code).to.be.equal(20000, '返回码异常');
          for(let i=0; i<response.data.count; i++){
            expect(response.data.items[i].password).to.be.equal(undefined, '列表中不能存在密码');
            expect(response.data.items[i].is_deleted).to.be.equal(false, '只能取出未删除的用户');
          }
          expect(response.data.count).to.be.equal(response.data.items.length, '列表数量必需一致');
        });
    });

    it('#管理员-更新用户', async () => {

      let userinfo = {};
      let update_ret;
      let get_ret;

      // 1. 构造用户信息
      get_ret = await request(server).get(`/api/user/${test_user_id}`).set('Authorization', test_user_token);
      //Object.assign(userinfo, get_ret._body.data);
      let old_user_info = get_ret._body.data;

      // 1. 构造用户信息
      userinfo.user_name = random(6, {letters: 'qweradflouiopuo6346247890ashgj', numbers: true});
      userinfo.friend_name = random(6, {letters: 'qweradflouiopuo6346247890ashgj', numbers: true});
      userinfo.position = random(6, {letters: 'qweradflouiopuo6346247890ashgj', numbers: true});
      userinfo.phone = random(6, {letters: 'qweradflouiopuo6346247890ashgj', numbers: true});
      userinfo.note = random(6, {letters: 'qweradflouiopuo6346247890ashgj', numbers: true});
      userinfo.phone = random(6, {letters: 'qweradflouiopuo6346247890ashgj', numbers: true});
      userinfo.is_admin = !old_user_info.is_admin;
      userinfo.is_deleted = !old_user_info.is_deleted;

      // 2. 更新用户
      update_ret = await request(server).put(`/api/user/${test_user_id}`).set('Authorization', admin_user_token).send(userinfo);
      //console.log('update_ret', update_ret._body);
      get_ret = await request(server).get(`/api/user/${test_user_id}`).set('Authorization', admin_user_token);
      //console.log('get_ret', get_ret._body);
      expect(get_ret._body.code).to.be.equal(20000, '返回码异常');
      expect(get_ret._body.data.user_name).to.be.equal(userinfo.user_name, '取到的用户名，与修改的必需一样');
      expect(get_ret._body.data.friend_name).to.be.equal(userinfo.friend_name, '取到的显示名，与修改的必需一样');
      expect(get_ret._body.data.position).to.be.equal(userinfo.position, '取到的部门，与修改的必需一样');
      expect(get_ret._body.data.phone).to.be.equal(userinfo.phone, '取到的phone，与修改的必需一样');
      expect(get_ret._body.data.note).to.be.equal(userinfo.note, '取到的phone，与修改的必需一样');
      expect(get_ret._body.data.is_admin).to.be.equal(userinfo.is_admin, '取到的is_admin，与修改的必需一样');
      expect(get_ret._body.data.is_deleted).to.be.equal(userinfo.is_deleted, '取到的is_deleted，与修改的必需一样');
      
    });


    it('#一般用户-更新自已成功', async () => {

      let userinfo = {};
      let update_ret;
      let get_ret;

      // 1. 构造用户信息
      get_ret = await request(server).get(`/api/user/${test_user_id}`).set('Authorization', test_user_token);
      //Object.assign(userinfo, get_ret._body.data);
      let old_user_info = get_ret._body.data;

      userinfo.user_name = random(6, {letters: 'qweradflouiopuo6346247890ashgj', numbers: true});
      userinfo.friend_name = random(6, {letters: 'qweradflouiopuo6346247890ashgj', numbers: true});
      userinfo.position = random(6, {letters: 'qweradflouiopuo6346247890ashgj', numbers: true});
      userinfo.phone = random(6, {letters: 'qweradflouiopuo6346247890ashgj', numbers: true});
      userinfo.note = random(6, {letters: 'qweradflouiopuo6346247890ashgj', numbers: true});
      userinfo.phone = random(6, {letters: 'qweradflouiopuo6346247890ashgj', numbers: true});
      userinfo.is_admin = !old_user_info.is_admin;
      userinfo.is_deleted = !old_user_info.is_deleted;

      // 2. 更新用户
      update_ret = await request(server).put(`/api/user/${test_user_id}`).set('Authorization', test_user_token).send(userinfo);
      //console.log('update_ret', update_ret._body);
      expect(update_ret._body.code).to.be.equal(20000, '返回码异常');
      expect(update_ret._body.data).to.be.equal(1, '更新数量必需为1');

      // 3. 更新现在的数据
      get_ret = await request(server).get(`/api/user/${test_user_id}`).set('Authorization', test_user_token);
      //console.log('get_ret', get_ret._body);
      expect(get_ret._body.code).to.be.equal(20000, '返回码异常');
      //console.log('取到的用户名，与修改的必需不一样', get_ret._body.data.user_name, old_user_info.user_name);

      // 注意，一般用户，不可以改，user_name，is_admin，is_deleted
      expect(get_ret._body.data.user_name).to.be.equal(old_user_info.user_name, '取到的用户名，与修改的必需一样(不能修改)');
      expect(get_ret._body.data.friend_name).to.be.equal(userinfo.friend_name, '取到的显示名，与修改的必需一样');
      expect(get_ret._body.data.position).to.be.equal(userinfo.position, '取到的部门，与修改的必需一样');
      expect(get_ret._body.data.phone).to.be.equal(userinfo.phone, '取到的phone，与修改的必需一样');
      expect(get_ret._body.data.note).to.be.equal(userinfo.note, '取到的phone，与修改的必需一样');
      expect(get_ret._body.data.is_admin).to.be.equal(old_user_info.is_admin, '取到的is_admin，与修改的必需不一样(不能修改)');
      expect(get_ret._body.data.is_deleted).to.be.equal(old_user_info.is_deleted, '取到的is_deleted，与修改的必需不一样(不能修改)');
      
    });


    it('#一般用户-更新别人出错', async () => {

      let userinfo = {};
      let update_ret;
      let get_ret;

      // 1. 构造用户信息
      get_ret = await request(server).get(`/api/user/${test_user_1_id}`).set('Authorization', test_user_token);
      //Object.assign(userinfo, get_ret._body.data);
      let old_user_info = get_ret._body.data;

      userinfo.user_name = random(6, {letters: 'qweradflouiopuo6346247890ashgj', numbers: true});
      userinfo.friend_name = random(6, {letters: 'qweradflouiopuo6346247890ashgj', numbers: true});
      userinfo.position = random(6, {letters: 'qweradflouiopuo6346247890ashgj', numbers: true});
      userinfo.phone = random(6, {letters: 'qweradflouiopuo6346247890ashgj', numbers: true});
      userinfo.note = random(6, {letters: 'qweradflouiopuo6346247890ashgj', numbers: true});
      userinfo.phone = random(6, {letters: 'qweradflouiopuo6346247890ashgj', numbers: true});
      userinfo.is_admin = !old_user_info.is_admin;
      userinfo.is_deleted = !old_user_info.is_deleted;

      // 2. 更新其它用户，信息
      update_ret = await request(server).put(`/api/user/${test_user_1_id}`).set('Authorization', test_user_token).send(userinfo);
      //console.log('update_ret', update_ret._body);
      expect(update_ret._body.code).to.be.equal(40200, '一般用户，只能改自己，不可以改别人');

    });



    it('#一般用户-删除用户出错', async () => {

      // 2. 删除用户
      await request(server)
        .delete(`/api/user/${test_user_1_id}`)
        .set('Authorization', test_user_token)
        .expect('Content-Type', /json/)
        .expect(200)
        .expect((res) => {
          let response = res.body;
          expect(response.code).to.not.be.equal(20000, '不可以删除别人');
        });
    });


    
  });
