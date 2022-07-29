const modelCaptcha = require('../model/captcha.model')
const serviceUser = require('./user.service')

const {AuthFailed, DBFailed, HttpException} = require('../core/exception')

const svgCaptcha = require('svg-captcha')
const random = require('string-random');

const jwt = require('jsonwebtoken')

class AuthService {
    async createCaptcha() {

        let captchaLength = process.env.CAPTCHA_LEN || 4;
        const captcha = svgCaptcha.create({
            size: captchaLength, // 验证码长度
            width:process.env.CAPTCHA_WIDTH || 139,
            height:process.env.CAPTCHA_HEIGHT || 38,
            fontSize: process.env.CAPTCHA_FONTSIZE || 50,
            ignoreChars: process.env.CAPTCHA_IGNORECHARS || '0oO1ilI', // 验证码字符中排除 0o1i
            noise: process.env.CAPTCHA_NOISE || 2, // 干扰线条的数量
            color: process.env.CAPTCHA_COLOR || false, // 验证码的字符是否有颜色，默认没有，如果设定了背景，则默认有
            //background: process.env.CAPTCHA_BACKGROUND || '#eee' // 验证码图片背景颜色
        })


        try{
            let captcha_timeout_s = process.env.CAPTCHA_TIMEOUT || 20; // 秒
            let ret = await modelCaptcha.create({ captcha: captcha.text.toLowerCase(), expiredAt: Date.now() + captcha_timeout_s * 1000 });
            let data = {
                captchaId: ret._id,
                captchaLength: captchaLength,
                picPath: captcha.data,
                text: ret.captcha
            }
            return data;
        }catch(error){
            console.error('createCaptcha error:', error);
        }
    }
    async getCaptcha(id){

        const whereOpt = {}
        id && Object.assign(whereOpt, { _id:id })
        //captcha && Object.assign(whereOpt, { captcha })
        
        let ret = await modelCaptcha.findOne(whereOpt)
        return ret;
    }
    
    async createJwtToken(username){
        try{
            let res = await this.getUserInfo(username);
            res.password = undefined;
            let token = jwt.sign(res, process.env.JWT_SECRET, {expiresIn: '30d'})
            return token;
        }catch(error){
            console.error(error);
            throw new HttpException("生成jwt token 出错", 50015, error);
        }
    }
    

    async getUserInfo(username){

        if(username === 'admin'){
            return {
                _id:'admin-id',
                user_name:'admin',
                friend_name:'admin',
                password:process.env.ADMIN_PASSWORD || 'admin',
                is_admin:true,
                is_deleted:false,
            }
        }
        
        let ret = await serviceUser.getUserAllInfo({user_name:username});
        let ret_data = {
            _id:ret._id,
            user_name: ret.user_name,
            friend_name: ret.user_name,
            password: ret.password,
            is_admin: ret.is_admin,
            is_deleted: ret.is_deleted,
        };
        return ret_data;
    }

    async verifyToken(token){

        return true;
    }
    
    async infoFromToken(token){

        let jwt = token.split('.')
        let info_str = Buffer.from(jwt[1], 'base64').toString();
        let info = JSON.parse(info_str);

        return info;
    }


}

module.exports = new AuthService();
