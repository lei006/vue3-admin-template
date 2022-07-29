
const modelUser = require('../model/user.model')

class UserService {
    async createUser(user_name, password, friend_name) {

        // 插入数据
        // await表达式: promise对象的值
        const res = await modelUser.create({ user_name, password, friend_name, is_admin:false,is_deleted:false});
        return res

    }
    async getUserInfo({ _id, user_name}) {

        let filter = {}
        _id && Object.assign(filter, { _id })
        user_name && Object.assign(filter, { user_name })  

        const res = await modelUser.findOne(filter, {password:0});

        return res;
    }
    async getUserAllInfo({ _id, user_name}) {

        let filter = {}
        _id && Object.assign(filter, { _id })
        user_name && Object.assign(filter, { user_name })  

        const res = await modelUser.findOne(filter);
        return res;
    }



    async usernameIsExist(username){
        const res = await modelUser.findOne({user_name:username});
        return res;
    }

    async check_username(username){
        var reg = /^[a-z0-9_-]{2,16}$/;
        return reg.test(username)
    }




    async update(id, data){

        data._id = undefined;
        data.createdAt = undefined;
        data.updatedAt = undefined;
        data.__v = undefined;

        data.user_name = undefined;
        data.is_admin = undefined;
        data.is_deleted = undefined;


        const res = await modelUser.updateOne({ _id: id }, data);
        return res.modifiedCount;
    }

    async update_admin(id, data){
        
        data._id = undefined;
        data.createdAt = undefined;
        data.updatedAt = undefined;
        data.__v = undefined;

        if(data.user_name){
            data.user_name = data.user_name.toLowerCase();
        }

        const res = await modelUser.updateOne({ _id: id }, data);
        
        return res.modifiedCount;
    }

    async getList() {
        return modelUser.find({is_deleted:false}, {password:0});
    }

    async delete(id){
        const res = await modelUser.updateOne({ _id: id }, {is_deleted:true});
        return res.modifiedCount;
    }

    //真实删除
    async delete_real_by_id(id){
        const res = await modelUser.deleteOne({ _id: id });
        return res.deletedCount;
    }
    async delete_real_by_username(username){
        const res = await modelUser.deleteOne({ user_name: username });
        return res.deletedCount;
    }
    
}



module.exports = new UserService();
