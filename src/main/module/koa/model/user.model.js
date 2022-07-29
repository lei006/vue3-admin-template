const {mongoose, Schema} = require('../core/mongo')

const UserSchema = new Schema({
    user_name: { type: String, required:true, unique:true, allowNull:false, comment: '登录名'},
    friend_name: { type: String ,required:true, allowNull:false, comment: '助记名'},    
    password: { type: String ,required:true, allowNull:false, comment: '密码'},
    position: { type: String ,required:false, allowNull:true, comment: '职位'},
    phone: { type: String ,required:false, allowNull:true, comment: '电话'},
    note: { type: String ,required:false, allowNull:true, comment: '备注'},
    is_admin: { type: Boolean,default:true, allowNull:false, comment: '管理员'},
    is_deleted: { type: Boolean,default:true, allowNull:false, comment: '已删除'},
}, { 
    timestamps: true, //默认为true，自动生成`createdAt`和`updatedAt`字段
    freezeTableName: false //默认为false,表名自动是使用复数形式，如user->users    
});



module.exports = mongoose.model('User', UserSchema);
