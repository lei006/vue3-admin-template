const {mongoose, Schema} = require('../core/mongo')

///////////////////////////////////////////////////
// 报告布局
///////////////////////////////////////////////////

const reportLayoutSchema = new Schema({
    friend_name: { type: String ,required:true, allowNull:false, comment: '容易记的名称'},
    user_name: { type: String,default:'system', allowNull:false, comment: '布局所属人'},
}, { 
    timestamps: true, //默认为true，自动生成`createdAt`和`updatedAt`字段
    freezeTableName: false //默认为false,表名自动是使用复数形式，如user->users    
});



module.exports = mongoose.model('ReportLayout', reportLayoutSchema);
