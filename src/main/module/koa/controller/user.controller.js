const Ret = require('../core/return')

const userService = require('../service/user.service')
const random = require('string-random');


class UserController{
    constructor(){
    }
    async add(ctx){

      let user = ctx.state.user;
      if(user.is_admin !== true){
          Ret.error(ctx, 40400, `请使用管理员帐户进行该操作`);
          return;
      }

      let user_name = random(5);
      let password = random(5);
      let friend_name = "新用户-"+password;

      let data = await userService.createUser(user_name,password, friend_name);
      Ret.success(ctx, data);
    }
    
    async delete(ctx){
      let user = ctx.state.user;
      if(user.is_admin !== true){
          Ret.error(ctx, 40400, `请使用管理员帐户进行该操作`);
          return;
      }

      //let data = await userService.delete(ctx.params.id);
      let data = await userService.delete_real_by_id(ctx.params.id);
      Ret.success(ctx, data);

    }


    async update(ctx){

      let user = ctx.state.user;
      if(user.is_admin !== true){


        let userinfo = await userService.getUserInfo({_id:ctx.params.id});
        if(!userinfo){
          Ret.error(ctx, 40100, `未找到该用户`);
          return;
        }

        if(user._id != userinfo._id) {
          Ret.error(ctx, 40200, `不允许修改其它用户的信息`);
          return;
        }

        console.log('一般用户，用这个');
        //一般用户，用这个 ： 一般用户，不可以修改用户名，管理员，及删除用户
        let data = await userService.update(ctx.params.id, ctx.request.body);
        Ret.success(ctx, data);
        return;
      }else{

        //////////////////////////////////////////////
        // 检查是否符合要求
        let test_user_name = ctx.request.body.user_name;
        let ret = await userService.check_username(test_user_name);
        if(!ret){
          Ret.error(ctx, 40101, `登录名[${test_user_name}]用户名不符合要求，请换一个`);
          return;
        }
        
        //////////////////////////////////////////////
        // 检查是否已经存在
        let test_user_id = ctx.request.body._id;
        ret = await userService.usernameIsExist(test_user_name);
        if(ret && ret._id != test_user_id){
          Ret.error(ctx, 40102, `登录名[${test_user_name}]已存在`);
          return;
        }

        //管理员，用这个
        let data = await userService.update_admin(ctx.params.id, ctx.request.body);
        Ret.success(ctx, data);
        return;

      }

    }
    async find(ctx){

      let userinfo = await userService.getUserInfo({_id:ctx.params.id});
      if(!userinfo){
        Ret.error(ctx, 40100, `未找到该用户`);
        return;
      }
      
      Ret.success(ctx, userinfo);
    }
    async list(ctx){
      let data = {items:[], count:0};

      data.items = await userService.getList();
      data.count = data.items.length;

      Ret.success(ctx, data);
    }


}


module.exports = new UserController;
