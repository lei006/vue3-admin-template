const {mongoose, Schema} = require('../core/mongo')

///////////////////////////////////////////////////
// 预置字段
///////////////////////////////////////////////////

const presetSchema = new Schema({
    field_type: { type: String,required:true , allowNull:false, comment: '字段类型'},
    field_name: { type: String ,required:true, allowNull:false, comment: '字段名称'},
    sort_show: { type: Number, default:0, allowNull:false, comment: '显示排序'},
    def_val: { type: Mixed,required:true, allowNull:false, comment: '字段缺省值'},

    

    desc: { type: String , comment: '说明'},
}, { 
    timestamps: true, //默认为true，自动生成`createdAt`和`updatedAt`字段
    freezeTableName: false //默认为false,表名自动是使用复数形式，如user->users    
});


module.exports = mongoose.model('PresetField', presetSchema);

