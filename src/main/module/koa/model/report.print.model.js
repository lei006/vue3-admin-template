const {mongoose, Schema} = require('../core/mongo')


///////////////////////////////////////////////////
// 报告打印时的图片
///////////////////////////////////////////////////

const reportSchema = new Schema({
    report_id: { type: String ,required:true, allowNull:false, comment: '报告所属ID'},
    is_deleted: { type: Boolean,default:false, allowNull:false, comment: '已删除'},
    thumbnail: { type: String ,required:true, allowNull:false, comment: '缩略图'},
}, { 
    timestamps: true, //默认为true，自动生成`createdAt`和`updatedAt`字段
    freezeTableName: false //默认为false,表名自动是使用复数形式，如user->users    
});


module.exports = mongoose.model('Report', reportSchema);
