const {mongoose, Schema} = require('../core/mongo')

///////////////////////////////////////////////////
// 报告-固定内容
///////////////////////////////////////////////////

const reportLayoutSchema = new Schema({
    friend_name: { type: String ,required:true, allowNull:false, comment: '容易记的名称'},
    user_name: { type: String,default:'system', allowNull:false, comment: '布局所属人'},
}, { 
    timestamps: true, //默认为true，自动生成`createdAt`和`updatedAt`字段
    freezeTableName: false //默认为false,表名自动是使用复数形式，如user->users    
});



/*

type PresetDoctorReportTemplate struct {
	Id       string `orm:"column(id);pk" bson:"id" json:"id"`
	Title    string `bson:"title" json:"title"`                     //域类型
	ParentId string `orm:"index" bson:"parent_id" json:"parent_id"` //字段名
	DataA    string `orm:"type(text)" bson:"data_a" json:"data_a"`  //a
	DataB    string `orm:"type(text)" bson:"data_b" json:"data_b"`  //b
	DataC    string `orm:"type(text)" bson:"data_c" json:"data_c"`  //c
	DataD    string `orm:"type(text)" bson:"data_d" json:"data_d"`  //d
	Sort     string `bson:"sort" json:"sort"`                       //排序使用

	SoftType int64 `orm:"index" bson:"soft_type" json:"soft_type"` //软件类型

	AllowDelete bool `orm:"index" bson:"allow_delete" json:"allow_delete"` //允许删除.

}


*/

module.exports = mongoose.model('ReportLayout', reportLayoutSchema);
