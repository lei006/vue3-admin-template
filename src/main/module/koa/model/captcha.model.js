const {mongoose, Schema} = require('../core/mongo')



const CaptchaSchema = new Schema({
    captcha: { type: String,required:true, allowNull:false, comment: '验证码'},
    expiredAt: { type: Date, default: Date.now },
}, { 
    timestamps:true     //自动生成 创建时间，更新时间
});



module.exports = mongoose.model('Captcha', CaptchaSchema);
