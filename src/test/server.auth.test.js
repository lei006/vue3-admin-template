import assert from 'assert';
require('dotenv').config()

const jwt = require('jsonwebtoken')
const serviceAuth = require('../main/module/koa/service/auth.service')


describe('main/module/koa/service/auth.service', function () {


    it('认证码: 创建，获取，', function () {
        serviceAuth.createCaptcha().then(function(add_ret){

            assert(add_ret.captchaId);
            assert(add_ret.captchaLength > 0);
            assert(add_ret.picPath);
            assert(add_ret.text.length === add_ret.captchaLength);

            serviceAuth.getCaptcha(add_ret.captchaId).then(function(get_ret){
                assert(get_ret.captcha === add_ret.text);
                assert(get_ret.expiredAt.getTime() > Data.now());
            });
        });
    });

    it('用户信息: admin', function () {
        serviceAuth.getUserInfo('admin').then(function(info){
            assert(info);
            assert(info.user_name === 'admin');
            assert(info.friend_name === 'admin');
            assert(info.password === (process.env.ADMIN_PASSWORD || 'admin') );
            assert(info.is_admin === true);
            assert(info.is_deleted === false);
        });
    });

    it('用户信息: 不存在的用户', function () {
        serviceAuth.getUserInfo('32452346245245245%$^&').then(function(info){
            assert(1==2);
        });
    });

    it('token验证: 创䢖admin 的 token', async function () {
        let token = await serviceAuth.createJwtToken('admin');
        assert(token.length>6);
        const arr = token.split('.');
        assert(arr.length === 3);
        var verify = jwt.verify(token, process.env.JWT_SECRET);
        assert(verify.user_name === 'admin')
    });

    


});

