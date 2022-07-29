const {mongoose, Schema} = require('../core/mongo')

///////////////////////////////////////////////////
// 报告
// 说明: 为了使用全局索引，所以把图片放在最外面
///////////////////////////////////////////////////

const reportSchema = new Schema({
    id: { type: String,required:true , allowNull:false, comment: '报告ID-唯一'},
    friend_name: { type: String ,required:true, allowNull:false, comment: '助记名'},
    is_deleted: { type: Boolean,default:false, allowNull:false, comment: '已删除'},
}, { 
    timestamps: true, //默认为true，自动生成`createdAt`和`updatedAt`字段
    freezeTableName: false //默认为false,表名自动是使用复数形式，如user->users    
});



module.exports = mongoose.model('Report', reportSchema);
