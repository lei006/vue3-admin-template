import assert from 'assert';
require('dotenv').config()

const jwt = require('jsonwebtoken')
const serviceUser = require('../main/module/koa/service/user.service')


describe('main/module/koa/service/user.service', function () {


    it('增加用户,删除用户', async function () {
        let user_name = 'test_user_test_test';
        let password = 'test_user';
        let friend_name = 'freid_name';
        
        let ret = await serviceUser.delete_real_by_username(user_name);
        ret = await serviceUser.createUser(user_name, password, friend_name)
        console.log(ret);
        assert(ret.user_name === user_name);
        assert(ret.password === password);
        assert(ret.friend_name === friend_name);

        ret = await serviceUser.delete_real_by_id(ret._id);

        assert(ret === 1);

    });


});

