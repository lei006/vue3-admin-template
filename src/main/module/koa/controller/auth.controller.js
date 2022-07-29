
const Ret = require('../core/return')

const userService = require('../service/user.service')
const authService = require('../service/auth.service')


class AuthController {
    constructor(){
        this.x = 111;
        this.y = 222;
      }
    async login(ctx) {

        const { username, password, captcha, captchaId } = ctx.request.body;
        /*
        let ret = await authService.getCaptcha(captchaId);
        if(!ret){
            Ret.error(ctx, 404, "未找到验证码");
            return;
        }
        if(ret.captcha !== captcha.toLowerCase()){
            console.log("验证码出错", captchaId, ret.captcha , captcha.toLowerCase());
            Ret.error(ctx, 404, "验证码出错");
            return;
        }
        let expiredAt = ret.expiredAt.getTime();
        if(expiredAt < Date.now()){
            console.log("验证码过期", expiredAt , Date.now());
            Ret.error(ctx, 404, "验证码过期");
            return;
        }
        */
        
        let user_info = await authService.getUserInfo(username);
        if(!user_info){
            Ret.error(ctx, 40010, "用户不存在");
            return;
        }

        if(password !== user_info.password) {
            Ret.error(ctx, 40011, "密码出错");
            return;
        }
        
        let token = await authService.createJwtToken(username);
        let data = {
            token:'Bearer ' + token,
        }

        Ret.success(ctx, data);
    }

    async captcha(ctx) {
        let data = await authService.createCaptcha();
        Ret.ok(ctx, data);
    }
    async info(ctx){

        let user = ctx.state.user;
        let user_info = await authService.getUserInfo(user.user_name);
        user_info.password = undefined;
        if(!user_info){
            Ret.error(ctx, 40011, "用户不存在");
            return;
        }

        Ret.success(ctx, user_info);
    }
    async logout(ctx) {
        Ret.success(ctx);
    }
}

module.exports = new AuthController;
